FROM golang:1.21.1-alpine as financial-file-manager

WORKDIR /app

COPY . .

EXPOSE  8080
