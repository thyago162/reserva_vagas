package main

import (
	"encoding/json"
	"log"
	"net/smtp"
	"os"

	"github.com/rabbitmq/amqp091-go"
)

type EmailReservaPayload struct {
	Email     string `json:"email"`
	Nome      string `json:"nome"`
	Senha     string `json:"senha"`
	Protocolo string `json:"protocolo"`
}

func main() {
	amqpURL := os.Getenv("RABBITMQ_URL")
	if amqpURL == "" {
		amqpURL = "amqp://admin:admin123.@localhost:5672/"
	}

	conn, err := amqp091.Dial(amqpURL)
	if err != nil {
		log.Fatalf("Worker falhou ao conectar no RabbitMQ: %v", err)
	}

	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Worker falhou ao abrir canal: %v", err)
	}

	defer ch.Close()

	q, err := ch.QueueDeclare(
		"fila_email_reserva",
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		log.Fatalf("Falha ao declarar fila no worker: %v", err)
	}

	msgs, err := ch.Consume(
		q.Name,
		"",
		false,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		log.Fatalf("Falha ao registrar consumidor: %v", err)
	}

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			var payload EmailReservaPayload
			err := json.Unmarshal(d.Body, &payload)
			if err != nil {
				log.Printf("[ERROR] Falha ao desserializar payload: %v", err)
				d.Nack(false, true)
				continue
			}

			err = enviarEmailPeloMailHog(payload.Email, payload.Nome, payload.Senha, payload.Protocolo)
			if err != nil {
				log.Printf("[ERROR] Falha ao enviar email pelo Mailhog: %v", err)
				d.Nack(false, true)
				continue
			}

			d.Ack(false)
			log.Printf("[SUCESSO] E-mail enviado e mensagem confirmada para: %s", payload.Email)
		}
	}()

	<-forever
}

func enviarEmailPeloMailHog(destinatario, nome, senha, protocolo string) error {
	auth := smtp.PlainAuth("", "", "", "localhost")

	from := "noreply@sistemadevagas.com"
	to := []string{destinatario}

	assunto := "Sua Reserva de Vaga foi Realizada! - " + protocolo
	corpo := "Olá, " + nome + "!\n\n" +
		"Sua reserva foi concluída com sucesso.\n" +
		"Foi criada uma conta de acesso para você no sistema.\n" +
		"Sua senha temporária é: " + senha + "\n\n" +
		"Guarde bem esta senha. Atenciosamente, Equipe de Vagas."

	msg := []byte("To: " + destinatario + "\r\n" +
		"Subject: " + assunto + "\r\n" +
		"Content-Type: text/plain; charset=UTF-8\r\n" +
		"\r\n" + corpo + "\r\n")

	// Endereço do MailHog local no Docker (porta 1025)
	err := smtp.SendMail("localhost:1025", auth, from, to, msg)
	return err
}
