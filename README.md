# Agenda de contatos
Sistema simples com crud para armazenamento e consulta de contatos

## Configuração
Altere o arquivo .env na raiz do projeto

## Administrador
**Usuário:** administrador
**Senha:** yivLTC12

## Execução embiente desenvolvimento
```env AMBIENTE=desenvolvimento go run *.go```
## Deploy
```go build -o "agenda" -ldflags "-s -w" main.go && upx agenda```



## Executando binário da aplicação
Entre no diretório: dist/v1/hospeda-app/agenda e execute o comando 
```sudo ./install```

Para desinstalar basta executar ```sudo ./uninstall```