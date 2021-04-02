# Agenda de contatos
Sistema simples com crud para armazenamento e consulta de contatos

## Configuração
Altere o arquivo .env na raiz do projeto

## Administrador
**Usuário:** administrador
**Senha:** yivLTC12

## Execução embiente desenvolvimento
```env AMBIENTE=desenvolvimento go run *.go```
## Gerando arquivo binário da aplicação
```go build -o "agenda" -ldflags "-s -w" main.go && upx agenda```

## Execução em ambiente de produção
Deve-se configurar dados de conexão ao banco de dados no arquivo .env

Executar o comando: ./deploy.sh para subir a apilcação usando Docker
```
> ./deploy
```
Saída do comando acima será semelhante:
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

Conforme a saída acima a aplicação poderá se acessada através do endreço: http://172.23.0.21
