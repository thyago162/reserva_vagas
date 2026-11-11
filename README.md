# reserva_vagas

Sistema inteligente de reserva de vagas escolares. O cidadão conversa com um agente de IA, que localiza as escolas municipais mais próximas pela distância e ajuda a concluir a matrícula. Projeto pessoal.

[![Arquitetura do sistema](docs/arquitetura.png)](docs/arquitetura.png)

## Componentes

| Repositório | Função | Tecnologias |
|---|---|---|
| [reserva_vagas](https://github.com/thyago162/reserva_vagas) (este) | API de vagas e reservas e worker de e-mails | Go, GORM, RabbitMQ |
| [reserva_vagas_front](https://github.com/thyago162/reserva_vagas_front) | Interface de conversa | Vue.js 3 |
| [reserva_vagas_mcp_client](https://github.com/thyago162/reserva_vagas_mcp_client) | Orquestrador entre o front e o LLM | Python, FastAPI |
| [reserva_vagas_mcp_server](https://github.com/thyago162/reserva_vagas_mcp_server) | Ferramentas MCP e sessões | Python, Redis |

## Destaques técnicos

- **Envio assíncrono de e-mails:** ao confirmar a reserva, a API usa goroutines e channels para publicar as mensagens numa fila do RabbitMQ, e um worker dedicado (mesmo repositório, `main` próprio) faz o envio. Assim a resposta é rápida mesmo com milhares de responsáveis a notificar.
- **Agente de IA com MCP:** o LLM decide quais ferramentas usar (localização, listagem de escolas, verificação e registro de reserva).
- **Busca por proximidade:** as escolas são ordenadas pela distância ao endereço do cidadão.

## Estrutura

```
<cmd/api/main.go>      API
<cmd/worker/main.go>   worker de e-mails
<internal/...>         <descreva as pastas principais>
```

## Como rodar

```bash
git clone https://github.com/thyago162/reserva_vagas
cd reserva_vagas
cp .env.example .env   # preencha as variáveis
go run <cmd/api>       # API
go run <cmd/worker>    # worker de e-mails (em outro terminal)
```

## Variáveis de ambiente

| Variável | Descrição |
|---|---|
| `<DB_...>` | Conexão com o banco de dados |
| `<RABBITMQ_URL>` | Conexão com o RabbitMQ |
| `<SMTP_...>` | Envio de e-mails |

## Demonstração

[Vídeo no YouTube](https://www.youtube.com/watch?v=uNiyldIHgVc)

## Status

Projeto em desenvolvimento.