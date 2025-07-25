FROM golang:1.24.3 AS builder

WORKDIR /app

COPY . .

RUN go mod tidy

RUN CGO_ENABLED=0 GOOS=linux go build -o server ./cmd/main.go 


FROM alpine

WORKDIR /root

COPY --from=builder /app/server ./

ENTRYPOINT [ "./server" ]

