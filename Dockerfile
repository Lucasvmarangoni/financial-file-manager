FROM golang:1.22.0-alpine AS build

RUN apk add --no-cache git

WORKDIR /app

COPY api/ ./api
COPY cmd/ ./cmd
COPY config/ ./config
COPY internal/ ./internal
COPY pkg/ ./pkg
COPY test/ ./test
COPY go.mod .
COPY go.sum .

RUN go mod tidy 

WORKDIR /app/cmd

RUN go build -o /server 

FROM gcr.io/distroless/base-debian12:latest

WORKDIR /app

COPY --from=build /server /server

EXPOSE  8000

USER nonroot:nonroot

ENTRYPOINT ["/server"]