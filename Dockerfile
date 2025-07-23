FROM golang:1.24.3 as builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod tidy

COPY *.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -o app


FROM alpine

WORKDIR /app

COPY --from=builder app ./

ENTRYPOINT [ "/app/app" ]

