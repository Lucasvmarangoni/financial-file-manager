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

## Problema no return da response da criação de usuário

Problema: Em algumas requests, o channel não esta recebendo os payloads. Isso faz com que a response não retorne. 

algumas vezes (Na maioria funciona normal) "err = <- u.ReturnChannel" não recebe o valor e isso impede o codigo de proseguir e retornar a response.

O pior é que a próxima request recebe a response da request anterior e isso vai se tornando uma bola de neve.

___
**Exemplo**:

Request 1: cria um usuário valido. (não retorna response)

na aplicação vejo nos logs que o usuário foi criado com sucesso. No banco de dados confirmo que foi criado corretamente.

Request 2: envio com o mesmo payload da Request 1 e a response retorna 201 Created, Quando deveria retornar :  400 bad request.
___

Foi o primeiro bug que não consegui resolver, decidi mudar a abordagem..

Eu realmente queria descobrir, odeio desistir dessa forma, mas tentei muito tempo e não estava encontrando respostas ou indicidios do fator causador do problema. Algumas vezes funcionava 100% perfeito, enquanto em outras não. tentei fazer todo tipo de alteração, algumas até malucas, mas nada resolveu definitivamente, como podem ver, postei aqui que havia solicionado duas vezes, eu acreditei que havia solucionado e só estava fazendo besteira adicionando goroutines desnecessárias e prejudiciais ao programa (eu não possuia muito conhecimento sobre go nessa época). 

_____
### Fluxo de Execução 

UserHandler.Create -> UserService.Create -> UserManagement.Management -> UserReposiroty.Insert

#### UserService.Create

Processa os dados (validação,  encryptação), coloca na fila rabbitMQ e aguarda o retorno pelo channel:

--> <a href="https://github.com/Lucasvmarangoni/financial-file-manager/blob/85122f90b940e90809c29a306c006557c4faee69/internal/modules/user/domain/services/create.service.go">link para essa versão</a>

```go
{...}
err = u.RabbitMQ.Publish(string(userJSON), "application/json", config.GetEnvString("rabbitMQ", "exchange"), config.GetEnvString("rabbitMQ", "queue_user"), config.GetEnvString("rabbitMQ", "routingkey_userCreate"))
	if err != nil {
		return errors.ErrCtx(err, "RabbitMQ.Publish")
	}

err = <- u.ReturnChannel	
	if err != nil {
		return errors.ErrCtx(err, "CreateManagement")
	}
	return nil
```

#### UserManagement.Management

Esta em constante execução dentro de uma goroutine. Recebe os dados pela fila para persistir no banco e retornar por meio do channel:

É executado junto com a inicialização do sistema na func init do router.go:

<a href="https://github.com/Lucasvmarangoni/financial-file-manager/blob/d5266451f0beb7795528b8acaa5764be9ea104df/internal/modules/user/http/routers/routers.go">Link para essa versão</a>

```go
{...}
u.RabbitMQ.Consume(u.MessageChannel, config.GetEnvString("rabbitMQ", "routingkey_userCreate"))
	for i := 0; i < config.GetEnvInt("concurrency", "create_management"); i++ {
		go userManagement.CreateManagement(u.MessageChannel, userService.ReturnChannel)	
	}
{...}
```

Dentro do metodo:

<a href="https://github.com/Lucasvmarangoni/financial-file-manager/blob/d5266451f0beb7795528b8acaa5764be9ea104df/internal/modules/user/domain/management/management.go">Link para essa versão</a>

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
_____

### Solução 

Antes de tudo, removi as goroutines desnecessárias e sem sentido que eu havia implementado no handler e no service.

Além disso, movi o consumer do  UserManagement.Management para o init do router. Isso me possibilitou utilizar multiplas goroutines executando o  UserManagement.Management.

<a href="https://github.com/Lucasvmarangoni/financial-file-manager/blob/ef41d08066deb94cda8b01fe2e7da5d35dd937db/internal/modules/user/http/routers/routers.go">Link para essa versão</a>

```go
func (u *UserRouter) init() *handlers.UserHandler {
	userRepository := repositories.NewUserRepository(u.Conn)
	userService := services.NewUserService(userRepository, u.RabbitMQ, u.MessageChannel, u.Memcached, *u.Memcached_1)
	userManagement := management.NewManagement(userRepository)
	userHandler := handlers.NewUserHandler(userService)

	u.RabbitMQ.Consume(u.MessageChannel, config.GetEnvString("rabbitMQ", "routingkey_userCreate"))
	for i := 0; i < config.GetEnvInt("concurrency", "create_management"); i++ {
		go userManagement.CreateManagement(u.MessageChannel)
	}
	return userHandler
}
```

A solução para o problema eu resolvi utilizando o cache, assim mantive a persistência dos dados deos usuários de forma concorrente em goroutines. 

O cache foi utilizado para verificar se já existe usuário e retornar a response sem aguardar o processamento.

Consequentemente foi necessário criar um novo sistema de cache para armazenar apenas email e cpf, que são os campos unique e que poderiam causar problemas na persistência do banco. Outras validações ocorrem na própria entity, portanto dificilmente os dados chegariam ao repositório estando incorretos. 
 
Essa abordagem foi necessária pois o sistema de cache já criado utiliza como key o ID, então eu não conseguiria buscar durante a criação de um novo usuário. 

Agora quando um novo usuário tenta se registrar é feita a consulta ao cache para verificar se existe um usuário com aquele email ou CPF. Caso exista é retornado erro.

```go
errD.New("duplicate key value violates unique constraint")
```

Se não existir, o processamento segue até ser publicado na fila, retornando 201 Created para o client enquanto a persistência dos dados ocorre em uma goroutine própria. 

```go
func (u *UserService) Create(name, lastName, cpf, email, password string) error {

	newUser, err := entities.NewUser(name, lastName, cpf, email, password)
	if err != nil {
		return errors.ErrCtx(err, "entities.NewUser")
	}

	err = u.CheckIfUserAlreadyExists(newUser.HashEmail, newUser.HashCPF, nil)
	if err != nil {
		return errors.ErrCtx(err, "u.CheckIfUserAlreadyExists")
	}

	err = u.encrypt(newUser)
	if err != nil {
		return errors.ErrCtx(err, "u.encrypt")
	}

	userJSON, err := json.Marshal(newUser)
	if err != nil {
		return errors.ErrCtx(err, "json.Marshal")
	}

	err = u.RabbitMQ.Publish(string(userJSON), "application/json", config.GetEnvString("rabbitMQ", "exchange"), config.GetEnvString("rabbitMQ", "queue_user"), config.GetEnvString("rabbitMQ", "routingkey_userCreate"))
	if err != nil {
		return errors.ErrCtx(err, "RabbitMQ.Publish")
	}

	u.setToMemcacheIfNotNil(newUser)
	u.Memcached_1.SetUnique(newUser.HashCPF)
	u.Memcached_1.SetUnique(newUser.HashEmail)

	return nil
}
```

## Outras alterações

Antes o projeto utilizava apenas SHA-256 para criptografar as Hashs de consulta na entidade. Decidi implementar o HMAC em conjunto ao SHA-256, para assim tornar a segurança mais robusta.

