# syntax=docker/dockerfile:1
FROM golang:1.26.11 AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .

#Compila
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o /bin/api ./cmd/api
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o /bin/worker ./cmd/worker

#Imagem da api
FROM gcr.io/distroless/static-debian12 AS api
COPY --from=builder /bin/api /api
EXPOSE 3000
USER nonroot:nonroot
ENTRYPOINT [ "/api" ]

#Worker
FROM gcr.io/distroless/static-debian12 AS worker
COPY --from=builder /bin/worker /worker
USER nonroot:nonroot
ENTRYPOINT ["/worker"]
