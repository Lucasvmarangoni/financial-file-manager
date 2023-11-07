# Sistema de Armazenamento e Controle de Documentos Financeiros


## ⇝ EM PLANEJAMENTO


## Projeto

Este projeto tem como objetivo desenvolver um sistema robusto para o armazenamento e controle de documentos financeiros. Ele oferece aos clientes e administradores de instituições financeiras a capacidade de fazer upload, gerenciar e recuperar documentos financeiros de forma eficiente.

### Principais recursos

- **Armazenamento Eficiente:** Os documentos são inicialmente armazenados localmente no servidor e, posteriormente transferidos para um serviço de armazenamento em nuvem.

- **Metadados Inteligentes:** Cada documento é acompanhado de metadados, incluindo nome, data, tipo de documento e informações relacionadas a transações financeiras.

- **Recuperação Personalizada:** Os usuários podem buscar documentos com base em critérios como tipo de documento, período de datas e outras informações relevantes.

- **Segurança e Controle:** O sistema mantém nível de segurança e controle de acesso rigoroso, garantindo que as informações sejam protegidas e apenas os documentos autorizados sejam acessados.

-  **Conformidade Regulatória**: A aplicação atende as normas estabelecidas na LGPD.

*<a href="./docs/doc-funcional.md"> ⇝ <u>Documentação detalhada.</u> </a>*


## Tecnologias

**Linguagem**: Go (Golang); <br>
**Banco de dados**: CockroachDB; <br>
**Driver de banco de dados**: Pgx<br>
**Armazenamento**: Google Cloud Storage;<br>
**Design de API**: GraphQL; <br>
**Arquitetura**: Clean Architecture;<br>


**Integridade**

- RabbitMQ

**Observabilidade**

- Zerolog
- Prometheus

**Segurança**

- JWT
- Bcrypt
- Crypto/aes
- Crypto/cipher

**Infraestrutura**

- Docker
- Github Actions



