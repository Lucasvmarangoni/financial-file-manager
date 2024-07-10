FROM golang@sha256:8e96e6cff6a388c2f70f5f662b64120941fcd7d4b89d62fec87520323a316bd9 AS build

RUN apk update --no-cache && apk upgrade --no-cache \
    && apk add --no-cache mailcap git tzdata ca-certificates

WORKDIR /app

COPY api/ ./api
COPY cmd/ ./cmd
COPY config/ ./config
COPY internal/ ./internal
COPY pkg/ ./pkg
COPY test/ ./test
COPY logs/ ./logs
COPY go.mod .
COPY go.sum .

RUN go mod tidy 

WORKDIR /app/cmd

RUN go build -o /server 

FROM gcr.io/distroless/base-debian12:latest

WORKDIR /app

COPY --from=build /server /server

VOLUME /app/logs

EXPOSE  8000

USER nonroot:nonroot

ENTRYPOINT ["/server"]