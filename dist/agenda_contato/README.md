# Agenda de Contatos

Sistema de agenda de contatos para uso empresarial, permite cadastro de:

* Nome da pessoa
* Departamento
* Email
* Telefone
* Ramal
* Celular

O sistema roda sobre container Docker e tem uma área administrativa onde se pode acessar e gerenciar os dados de contatos.
O sistema tem uma área pública que não precisa de login e senha para acessar, nesta área pode-se visualizar e pesquisar todos os contatos cadastrados na agenda.

## Configuração:

**Observação:** Antes de mais nada, deve-se configurar arquivo .env com dados de conexão, caso contrário ao subir o container, ele ficará inativo por conta da conexão com banco. Portando deve-se ajustar as configurações e restar o containner para que ele venha a funcionar.

1. Para utilizar o sistema você deve-se possuir um servidor de banco de dados MySQL ou MariaDB e criar banco de dados chamado sistema_interno_contato.

2. Deve-se informar dados de conexão do banco de dados no arquivo .env

3. A aplicação roda na porta 80 dentro do container docker, caso deseje alterar, basta mudar as configuarções no arquivo .env.

## Executando aplicação:

A aplicação roda sobre container Docker bastando executar o comando: docker-compose up --build -d

A saída do comendo acima será algo semelhante com:

```
Building agenda.contato
Step 1/7 : FROM golang
 ---> bfb42f2526a7
Step 2/7 : MAINTAINER Hospeda App <ti@hospeda.app>
 ---> Using cache
 ---> 1497cdb89403
Step 3/7 : RUN apt-get update
 ---> Using cache
 ---> 0fcc5ac0ba86
Step 4/7 : RUN apt-get install ffmpeg -y
 ---> Using cache
 ---> 092621c625cc
Step 5/7 : WORKDIR /aplicacao
 ---> Using cache
 ---> 00b58c5c8b8b
Step 6/7 : ENTRYPOINT ./agenda
 ---> Using cache
 ---> ff8afb590a91
Step 7/7 : EXPOSE 80
 ---> Using cache
 ---> 3157ef7e361c
Successfully built 3157ef7e361c
Successfully tagged agenda_contato_agenda.contato:latest
Starting agenda.contato ... done
```

Se a aplicação for executada em sistema Linux baseado em Debian, pode executar o script deploy.sh com comando no terminal ./deploy.sh, a saída será algo semelhante: 

```
---------------------------------------------------------
Subindo aplicação:
---------------------------------------------------------
Building agenda.contato
Step 1/7 : FROM golang
 ---> bfb42f2526a7
Step 2/7 : MAINTAINER Hospeda App <ti@hospeda.app>
 ---> Using cache
 ---> 1497cdb89403
Step 3/7 : RUN apt-get update
 ---> Using cache
 ---> 0fcc5ac0ba86
Step 4/7 : RUN apt-get install ffmpeg -y
 ---> Using cache
 ---> 092621c625cc
Step 5/7 : WORKDIR /aplicacao
 ---> Using cache
 ---> 00b58c5c8b8b
Step 6/7 : ENTRYPOINT ./agenda
 ---> Using cache
 ---> ff8afb590a91
Step 7/7 : EXPOSE 80
 ---> Using cache
 ---> 3157ef7e361c
Successfully built 3157ef7e361c
Successfully tagged agenda_contato_agenda.contato:latest
Starting agenda.contato ... done


---------------------------------------------------------
Obtendo ip da aplicação:
---------------------------------------------------------

172.23.0.21
```

Conforme exibido acima a aplicação estará disponível através do endereço: http://172.23.0.21

## Dados de acesso
**Usuário:** administrador
**Senha:** yivLTC12