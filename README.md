# Agenda de contatos
Sistema de agenda de contatos, versão 1.0

## Tecnologias
* Go versão 1.14.4 linux/amd64
* https://github.com/ColorlibHQ/AdminLTE/releases/tag/v3.0.5
* Bootstrap v4.5.0 (https://getbootstrap.com/)
* https://icons.getbootstrap.com/

## Dependências
* github.com/joho/godotenv
* github.com/gorilla/mux
* golang.org/x/crypto/bcrypt
* github.com/go-sql-driver/mysql

Para instalar dependência basta executar o comando go get -u link-repositório, exemplo: : ```go get -u github.com/go-sql-driver/mysql```

## Configurando a aplicação 
Altere o arquivo .env na raiz do projeto

## Administrador
**Usuário:** tayron
**Senha:** NftK2O7y

## Execução embiente desenvolvimento
```env AMBIENTE=desenvolvimento go run *.go```

## Executando binário da aplicação
Entre no diretório: dist/v1/hospeda-app/agenda e execute o comando 
```sudo ./install```

Para desinstalar basta executar ```sudo ./uninstall```