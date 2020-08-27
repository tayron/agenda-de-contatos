package bootstrap

import (
	"path/filepath"

	"github.com/tayron/agenda-contatos/bootstrap/library/util"
	"github.com/tayron/agenda-contatos/models"
	"github.com/gorilla/sessions"
	"github.com/joho/godotenv"
)

var store = sessions.NewCookieStore([]byte("1n+3gra-s1s+3ma"))

// Inicializa as configurações básica do sistema
func init() {

	caminhoAplicacao := util.ObterCaminhoDiretorioAplicacao()

	environmentPath := filepath.Join(caminhoAplicacao, ".env")
	err := godotenv.Load(environmentPath)

	if err != nil {
		panic(err)
	}

	models.CriarTabelaContato()
	models.CriarTabelaUsuario()
	models.CriarUsuarioAdministrador()
}

// StartApplication - Carrega as rotas e inializa a aplicação
func StartApplication() {
	CarregarRotas()
	StartarServidor()
}
