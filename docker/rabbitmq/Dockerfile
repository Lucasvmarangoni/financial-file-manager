FROM rabbitmq:3-management@sha256:75eb36982d09f096ca59360379cbc193cf488d0a5580dcb12c49df0450cd7199 AS  rabbitmq

USER nonroot:nonroot


EXPOSE 5672
EXPOSE 15672

CMD ["rabbitmq-start"]
