package messasing

import (
	"context"
	"encoding/json"
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitPublisher struct {
	channel *amqp.Channel
}

func NewRabbitPublisher(ch *amqp.Channel) *RabbitPublisher {
	return &RabbitPublisher{channel: ch}
}

func (p *RabbitPublisher) PublicarEmailReservaAsync(email, nome, senha string) {
	go func() {
		defer func() {
			if r := recover(); r != nil {
				log.Printf("[ERRO CRITICO] Panic recuperado ao enviar mensagem: %v", r)
			}
		}()

		payload := map[string]string{
			"email":     email,
			"nome":      nome,
			"senha":     senha,
			"prococolo": "RES-" + time.Now().Format("20060102150405"),
		}

		body, err := json.Marshal(payload)
		if err != nil {
			log.Printf("[ERRO] Falha ao serializar payload: %v", err)
			return
		}

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		err = p.channel.PublishWithContext(ctx,
			"",
			"fila_email_reserva",
			false,
			false,
			amqp.Publishing{
				ContentType: "application/json",
				Body:        body,
			},
		)
		if err != nil {
			log.Printf("[ERRO] Falha ao publicar no RabbitMQ: %v", err)
		}

		log.Printf("[SUCESSO] Mensagem enfileirada para: %s", email)
	}()
}
