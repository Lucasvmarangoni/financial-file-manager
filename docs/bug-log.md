# BUG LOG

## DB Persistence

### Solution:

It was necessary to use:

```go
err = crdbpgx.ExecuteTx(ctx, conn, pgx.TxOptions{}, func(tx pgx.Tx) error {
    // Your transaction logic here
})
```

when executing repository operations.

Additionally,

Instead of using the `tx` from the anonymous function `func(tx pgx.Tx) error`, do not use:

```go
// tx, err := conn.Begin(ctx)
// if err != nil {
//     conn.Close(ctx)
//     return nil, nil, errors.ErrCtx(err, "conn.Begin")
// }
```

This was what I was doing before, using this instance throughout the project.

## The user has already been created, it does not display the log and returns response

I used sync.WaitGroup to solve this.


![Image](https://github.com/Lucasvmarangoni/financial-file-manager/assets/101952043/2d858f64-bad8-4399-865f-669c923c41c1)



## The 200 status doesn't usually return when creating a user

The 200 status doesn't usually return when creating a user, only if the server is shut down or if attempting to immediately register the same user, in which case it returns both 200 and 400

### Solution

```go
func (m *UserManagement) CreateManagement(messageChannel chan amqp.Delivery) error {

	m.RabbitMQ.Consume(messageChannel, config.GetEnv("rabbitMQ_routingkey_userCreate").(string))

	for message := range messageChannel {
		var user entities.User
		err := json.Unmarshal(message.Body, &user)
		if err != nil {
			return errors.ErrCtx(err, "json.Unmarshal")
		}

		err = m.Repository.Insert(&user, context.Background())
		if err != nil {			
			return errors.ErrCtx(err, "Repository.Insert")
		}
		log.Info().Str("context", "UserHandler").Msgf("User created successfully (%s)", user.ID)
		return nil  // The error consists in not returning nil inside the loop.
	}	
	return nil
    }
```

Isso causou outro problema grave. Ao criar o usuário, o loop era encerrado, tornando impossível realizar qualquer nova operação de criação de usuário

### Solução:

Modifiquei o método *CreateManagement* para enviar diretamente para o returnChannel. Além disso, utilizei uma condicional else para garantir que nil fosse inserido no returnChannel apenas quando não houvesse erros.

```go
func (m *UserManagement) CreateManagement(messageChannel chan amqp.Delivery, returnChannel chan error) {

	m.RabbitMQ.Consume(messageChannel, config.GetEnv("rabbitMQ_routingkey_userCreate").(string))

	for message := range messageChannel {
		var user entities.User
		err := json.Unmarshal(message.Body, &user)
		if err != nil {
			returnChannel <- errors.ErrCtx(err, "json.Unmarshal")
		}

		err = m.Repository.Insert(&user, context.Background())
		if err != nil {
			returnChannel <- errors.ErrCtx(err, "Repository.Insert")
		} else {
			returnChannel <- nil
			log.Info().Str("context", "UserHandler").Msgf("User created successfully (%s)", user.ID)
		}
	}
}
```

Dessa forma, foi necessário alterar a chamada do método *CreateManagement* no *init*() do *router*, o que resultou em um código muito mais simples e limpo, principalmente ao melhorar a repartição de responsabilidades

```go
go func() {
		userManagement.CreateManagement(u.MessageChannel, returnChannel)		
	}()
```

antes:

```go
	var err error
	go func() {
		err = userManagement.CreateManagement(u.MessageChannel)
		if err != nil {
			returnChannel <- errors.ErrCtx(err, "u.CreateManagement")
		}
		returnChannel <- err
	}()
```