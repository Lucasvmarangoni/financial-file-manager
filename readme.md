<div align="center">
<a href="#projeto" target="_blank">
    <img align="center" src="https://img.shields.io/badge/-Projeto-05122A?style=flat&logo=" alt=""/>
  </a> 
 <a href="#tecnologias">
     <img align="center" src="https://img.shields.io/badge/-Tecnologias-05122A?style=flat&logo=Tecnologias" alt=""/>  
</a>       
</div>

<br>

<div align="center">

# Financial file manager
## Sistema de Armazenamento e Controle de Documentos Financeiros

## <u>EM DESENVOLVIMENTO</u>
Acompanhe: <a href="https://github.com/users/Lucasvmarangoni/projects/11"> Project Board</a>

<br>

**Projeto de portfólio tem como objetivos aprendizado, prática e demonstração das minhas habilidades como desenvolvedor. Portanto, as escolhas no projeto são com esses objetivos.**

</div>

<br>



## PROJETO

Este projeto tem como objetivo desenvolver um sistema robusto para o armazenamento e controle de documentos financeiros. Ele oferece aos clientes e administradores de instituições financeiras a capacidade de fazer upload, gerenciar e recuperar documentos financeiros de forma eficiente.

### Principais recursos

- **Documentos**: Os documentos são inicialmente armazenados localmente no servidor e, posteriormente transferidos para um serviço de armazenamento em nuvem. Sendo requisito manter versões do documento para visualizar seu histórico.

- **Consultas**: Cada documento é acompanhado de metadados, incluindo nome, data, tipo de documento e informações relacionadas a transações financeiras.

- **Segurança**: O sistema mantém nível de segurança e controle de acesso rigoroso, além de observabilidade e registro de atividades para fins de auditoria.

- **Resiliência**: O sistema deve ser projetado para minimizar o impacto de falhas e garantir que as operações possam ser retomadas o mais rápido possível após uma falha.

*<a href="./docs/doc-funcional.md"> ⇝ <u>Documentação detalhada.</u> </a>*

<br>

## TECNOLOGIAS

**Linguagem**: Go (Golang) <br>
**Arquitetura**: Microservices e Domain-Driven Design (DDD) <br>

### Persistência de dados

- **Banco de dados**: CockroachDB 
- **Driver de banco de dados**: Pgx 
- **Armazenamento**: Google Cloud Storage
- **Cache**: Memchached
- **Transporte**: 
  - **Filas**: RabbitMQ 
  - **Comunicação de Serviço**: gRPC

### HTTP

- **API**: Rest, GraphQL e gRPC 
- **Roteador**: go-chi
- **Framework GraphQL**: 99designs/gqlgen

### Observabilidade

- **Logs**: Zerolog 
- **Métricas**: Prometheus 

### Segurança

- **Autenticação e Autorização**: JSON Web Token (JWT)
- **Criptografia (password)**: Bcrypt
- **Criptografia (Dados sensíveis)**: AES

### Infraestrutura

- **Contêineres**: Docker
- **CI/CD**: Github Actions



