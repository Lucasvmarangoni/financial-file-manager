<div align="center">
<a href="#contexto" target="_blank">
    <img align="center" src="https://img.shields.io/badge/-Contexto-05122A?style=flat&logo=" alt=""/>
  </a>
  </a>  
     <a href="#projeto">
     <img align="center" src="https://img.shields.io/badge/-Projeto-05122A?style=flat&logo=Tecnologias" alt=""/>
     </a>
  <a href="#histórias-de-usuário">
     <img align="center" src="https://img.shields.io/badge/-Histórias-05122A?style=flat&logo=Tecnologias" alt=""/>  
      </a>  
       <a href="#requisitos-não-funcionais">
     <img align="center" src="https://img.shields.io/badge/-RNF-05122A?style=flat&logo=Tecnologias" alt=""/>  
      </a>       
</div>

<br><br>

# Documentação de Especificação Funcional

## Contexto

Sou desenvolvedor de uma instituição financeira que lida com uma grande quantidade de documentos financeiros, incluindo contratos, extratos bancários, faturas, relatórios fiscais e outros. O objetivo é criar uma aplicação em Go que simplifique o gerenciamento desses documentos, permitindo a manipulação de arquivos e, ao mesmo tempo, atendendo às necessidades financeiras da empresa.

## Projeto

1. **Armazenamento e Manipulação de Arquivos:**

   Realizar operações como fazer upload, categorização e adição de metadados aos documentos, tudo isso envolvendo a manipulação de arquivos no sistema de arquivos do servidor.

   Quando um usuário faz o upload de um documento financeiro, como uma fatura, contrato ou extrato, o arquivo é inicialmente armazenado localmente no servidor. Isso é feito para que o upload seja rápido e eficiente.

   Simultaneamente, informações sobre esse arquivo, (metadados), são coletadas e armazenadas em um banco de dados. 
   
   Esses metadados incluem informações como:   
   - Nome do arquivo.
   - Data de upload.
   - Tipo de documento (por exemplo, contrato, fatura, extrato).
   - Cliente ou fornecedor associado ao documento.
   - Número de transação ou referência financeira, se aplicável.

- O armazenar metadados no banco de dados permite buscas eficientes e organização dos documentos de acordo com as necessidades do banco ou dos clientes.

    **Integração com Cloud Storage:**

    Após o upload, o arquivo local no servidor será **transferido** para o serviço de armazenamento em nuvem. Isso é feito para garantir a persistência e a disponibilidade dos arquivos, mesmo em casos de falhas no servidor.

    - Armazenar os arquivos em um serviço em nuvem oferece maior possibilidade e facilidade de escalabilidade, segurança e redundância (disponibilidade).

2. **API:**
   - Será desenvolvida uma API para que os usuários possam acessar e interagir com o sistema. A API permitiria que os clientes ou administradores acessem os recursos do sistema via solicitações HTTP, como upload, download e pesquisa de documentos.

3. **Usuários**

   No contexto da ideia apresentada, "clientes" serão os clientes da instituição financeira. Podendo realizar operações somente pertinentes ao usuário (cliente). 
   
   Os "administradores" serão os funcionários da instituição financeira, com amplo acesso ao sistema, tendo poder para realizar operações que impactem os clientes ou outros administradores, de acordo com sua regras de autorização.

   **Administradores**

   Os administradores serão divididos grupos com diferentes niveis de autorização, seguindo o princípio de "privilégio mínimo" que significa que os administradores devem ter apenas as permissões necessárias para executar suas tarefas, de acordo com suas funções.
   
    **Recuperação e Visualização:**

    1. Quando um usuário deseja acessar ou baixar um documento, a aplicação consulta o banco de dados para encontrar os metadados associados a esse documento. Isso permite a recuperação eficiente das informações relevantes.

    2. Após a recuperação dos metadados, a aplicação pode direcionar o usuário para baixar o(s) arquivo(s) correspondente(s) do serviço de armazenamento em nuvem.
    

## Histórias de Usuário:

1. **Armazenamento de Documentos:**
   - Como um usuário, quero ser capaz de fazer upload de documentos financeiros em vários formatos (PDF, DOC, XLS, etc.) para o sistema.
   - Como um usuário, quero que os documentos sejam categorizados por tipo, data e cliente/fornecedor.
   - Como um usuário, quero que sejam adicionados metadados relevantes aos documentos para facilitar a pesquisa.

2. **Controle de Versões:**
   - Como um usuário, desejo que sejam mantidas várias versões de um documento e possível visualizar seu histórico.
   - Como um usuário, quero poder comparar versões anteriores de documentos para identificar alterações.

3. **Pesquisa e Recuperação:**
   - Como um usuário, desejo uma pesquisa avançada que me permita localizar documentos com base em palavras-chave, datas e metadados associados.
   - Como um usuário, quero poder baixar documentos recuperados.

4. **Integração Financeira:**
   - Como um usuário, quero a capacidade de vincular documentos a transações financeiras específicas, como faturas ou contratos, para um melhor rastreamento.
      
      - Ex: Vincular um comprovante de pagamento à um contrato.       

   - Como um usuário, desejo a capacidade de gerar relatórios financeiros a partir dos documentos armazenados.

5. **Segurança e Auditoria:**
   - Como um usuário, desejo que o sistema seja seguro, com permissões de acesso controladas por funções e registro de atividades para fins de auditoria.


## Requisitos não funcionais

1. **Segurança**: A aplicação deve ser altamente segura, garantindo a proteção dos dados financeiros dos clientes. Isso inclui criptografia de dados, autenticação de usuários e autorização rigorosa.

2. **Tempo de Resposta**: Buscar ter o melhor tempo de resposta posível, para garantir que as ações dos usuários sejam concluídas com eficiência.
   - Cache, indexação de bancos de dados e otimização de consultas.

3. **Escalabilidade**: O sistema deve ser escalável para acomodar um grande volume de documentos e usuários.

4. **Monitoramento e Registro**: Implemente monitoramento em tempo real e registros de atividades para rastrear o desempenho, identificar problemas e garantir a integridade dos dados.

5. **Documentação da API**: A documentação deve ser precisa, clara e completa para permitir integrações fáceis com outros sistemas.

6. **Conformidade Regulatória**: A aplicação deve atender as normas estabelecidas na LGPD.

7. **Auditoria**: A aplicação deve ser capaz de rastrear e registrar atividades de usuários, permitindo auditoria para fins de conformidade e segurança.

8. **Exclusão de Dados**: Usuários com acesso "clientes" não podem realizar nem uma operação de deletar dados, apenas usuários administradores com autorização devem poder deletar dados.

