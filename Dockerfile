FROM golang:1.24.3 AS builder

WORKDIR /app

COPY . .

RUN go mod tidy

RUN CGO_ENABLED=0 GOOS=linux go build -o app ./cmd/main.go 


FROM alpine

WORKDIR /app

COPY --from=builder app ./

ENTRYPOINT [ "/app/app" ]

