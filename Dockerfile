FROM golang:1.22.0-alpine as financial-file-manager

WORKDIR /app

COPY . .

RUN go mod tidy

FROM alpine:latest

WORKDIR /app

RUN go build -o financial-file-manager

COPY --from=builder /app/financial-file-manager /app/financial-file-manager

EXPOSE  8000

CMD ["./financial-file-manager"]