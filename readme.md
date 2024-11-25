<div align="center">
<a href="#projeto" target="_blank">
    <img align="center" src="https://img.shields.io/badge/-Projeto-05122A?style=flat&logo=" alt=""/>
  </a> 
 <a href="#tecnologias">
     <img align="center" src="https://img.shields.io/badge/-Tecnologias-05122A?style=flat&logo=Tecnologias" alt=""/>  
</a>   
 <a href="#execução">
     <img align="center" src="https://img.shields.io/badge/-Execução-05122A?style=flat&logo=Execução" alt=""/>  
</a>     
</div>

<br>

<div align="center">

# Financial file manager
## Sistema de Armazenamento e Controle de Documentos Financeiros

## <u>PROJETO PARADO</u>
   

**Acompanhe: <a href="https://github.com/users/Lucasvmarangoni/projects/11"> Project Board</a>**

</div>

**Nota:** O projeto teve uma pausa no desenvolvimento, que durou de 29 Mar 2024 até 02 Jul 2024.  
**Nota:** O desenvolvimento do projeto foi pausado, pois estou me decicando a outro projeto, desenvolvido com código fechado. Nesse momento, não tenho intenção de modificar o código desse projeto, contudo posso fazer atualizações nos *<a href="./docs/registros-de-engenharia-de-software.md"> <u>Registros de engenharia de software</u>.</a>*, fazendo correções a erros que tive na engenharia do projeto.   


## PROJETO

**Projeto de portfólio, tem como objetivos aprendizado, prática e demonstração das minhas habilidades como desenvolvedor. Portanto, as escolhas no projeto são, principalmente, com esses objetivos.**

Este projeto tem como objetivo desenvolver um sistema robusto para o armazenamento e controle de documentos financeiros. Ele oferece aos clientes e administradores de instituições financeiras a capacidade de fazer upload, gerenciar e recuperar documentos financeiros de forma eficiente.

### Principais recursos

- **Documentos**: Os documentos são inicialmente armazenados localmente no servidor e, posteriormente transferidos para um serviço de armazenamento em nuvem. Sendo requisito manter versões do documento para visualizar seu histórico.

- **Consultas**: Cada documento é acompanhado de metadados, incluindo nome, data, tipo de documento e informações relacionadas a transações financeiras.

- **Segurança**: O sistema mantém nível de segurança e controle de acesso rigoroso, além de observabilidade e registro de atividades para fins de auditoria.

- **Resiliência**: O sistema deve ser projetado para minimizar o impacto de falhas e garantir que as operações possam ser retomadas o mais rápido possível após uma falha.

⇝ *<a href="./docs/doc-de-requisitos.md"> <u>Documentação de requisitos</u>.</a>*

⇝ *<a href="./docs/registros-de-engenharia-de-software.md"> <u>Registros de engenharia de software</u>.</a>* 

⇝ *<a href="./docs/bug-log.md"> <u>Bug Log</u>.</a>* Registro de bugs/problemas e suas soluções.



### Artigos e conteúdo

Tenho o hábito de criar meu próprio material durante o estudo. A partir deste projeto, decidi publicá-los.

⇝ *<a href="https://medium.com/@lucasvm.ti/desenvolvimento-voltado-para-a-auditoria-em-software-0d15c56bf99c"> <u>Desenvolvimento voltado para a auditoria em Software</u>.</a>*

⇝ *<a href="https://medium.com/@lucasvm.ti/erros-e-logs-4e1fcd79a937"> <u>Como configurar erros e logs corretamente</u>.</a>*

### Bibliotecas 

Bibliotecas e pacotes que criei em razão desse projeto.

⇝ *<a href="https://github.com/Lucasvmarangoni/logella"> <u>Logella</u>.</a>*


<br>

## EXECUÇÃO 

Para iniciar o projeto nasta executar o comando:

    bash bash scripts/start.sh

Antes disso é necessário realizar as configurações do ambiente.

### Configurações

#### Permissões

De permissão para o container escrever na pasta logs. Isso é necessário para que o sistema possa salvar os arquivos de logs na pasta logs, localizada na raiz do projeto.

Inicialmente pode ser necessário dar permissão de read and writer (5) para others. Isso apenas para poder verificar quem é o usuário do container no host. Antes de executar o programa execute:

    chmod -R 775 logs

Em seguida pode iniciar o programa. Ele vai criar os arquivos de log na pasta logs.

Para identificar o usuário do container no host basta verificar qual usuário escreveu os arquivos de log:

    ls -lt logs

A saida será algo assim:

    $ ls -lt logs
    total 0
    -rw-r--r-- 1 65532 65532 0 Jul  9 21:28 modsec_audit.log
    -rw-r--r-- 1 65532 65532 0 Jul  9 21:28 modsec_debug.log

Identificado o usuário e grupo, basta dar a permissão. Eu escolhi dar permissão de owner e mantive o grupo do meu usuário.

    sudo chown 65532:<myusergroup> logs

Pronto. As permissões necessárias estão prontas.

#### Variáveis e Secrets

Configure o arquivo .env. As variáveis necessárias estão no arquivo example.env, aqui nesse repositório.

<a href="https://github.com/Lucasvmarangoni/financial-file-manager/blob/main/example.env">example.env</a>

Configure as secrets, elas estão no diretório "secrets", aqui nesse repositório.

<a href="https://github.com/Lucasvmarangoni/financial-file-manager/tree/main/secrets">secrets</a>


## TECNOLOGIAS

**Linguagem**: Go (Golang) <br>
**Arquitetura**: Microservices e Domain-Driven Design (DDD) <br>

### Persistência de dados

- **Banco de dados**: CockroachDB 
- **Driver de banco de dados**: Pgx 
- **Armazenamento**: Google Cloud Storage
- **Cache**: Memcached
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
- **Busca e Análise**: Elasticsearch

### Segurança

- **Autenticação e Autorização**: JSON Web Token (JWT)
- **Criptografia (password)**: Bcrypt
- **Criptografia (Dados sensíveis)**: AES e SHA-256

### Infraestrutura

- **Contêineres**: Docker
- **CI/CD**: Github Actions



