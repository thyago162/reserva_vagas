# Reserva de Vagas - Arquitetura & Ecossistema

    O Reserva de Vagas é um sistema de alta performance, modular e distribuído, desenvolvido sob medida para cenários de alta concorrência.
    A solução adota o Model Context Protocol (MCP) para permitir a integração nativa com assistentes inteligentes e sistemas de IA, mantendo o isolamento total da regra de negócio central.

 [Arquitetura do sistema](docs/arquitetura.png)

# reserva_vagas (Backend Core)

    Linguagem / Framework: Go (Golang)

    Tecnologias Integradas: Mysql, Redis, RabbitMQ

    Descrição: É a peça fundamental do ecossistema. Contém todas as regras de negócio, governança de dados, validações de elegibilidade e segurança.

# Principais Responsabilidades:

    Concorrência Segura: Gerenciamento de disputas por vagas em tempo real utilizando locks distribuídos no Redis (evitando overselling).
    Persistência de Dados: Modelagem e transações com integridade no PostgreSQL.
    Processamento Assíncrono: Emissão de eventos de auditoria e envio de comprovantes/notificações via RabbitMQ.

# Pré-requisitos & Execução Local

## Dependências Exigidas: 
    Go: 1.22
    RabbitMQ
    Mysql8


