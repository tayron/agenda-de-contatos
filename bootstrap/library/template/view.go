package template

import (
	"net/http"
	"os"
	"text/template"

	"github.com/tayron/agenda-contatos/bootstrap/library/session"
	"github.com/tayron/agenda-contatos/bootstrap/library/util"
)

// ObterInformacaoSistema -
func ObterInformacaoSistema(w http.ResponseWriter, r *http.Request) System {
	nome := session.GetDadoSessao("login", w, r)

	return System{
		Name:    os.Getenv("NOME_SISTEMA"),
		Version: os.Getenv("VERSAO_SISTEMA"),
		Usuario: nome,
	}
}

// LoadView -
func LoadView(w http.ResponseWriter, adicionalPath string, viewName string, parametros interface{}) {

	caminhoAplicacao := util.ObterCaminhoDiretorioAplicacao()

	var templates = template.Must(template.ParseGlob(caminhoAplicacao + "/template/*.html"))
	template.Must(templates.ParseGlob(caminhoAplicacao + "/template/layout/*.html"))

	if adicionalPath != "" {
		template.Must(templates.ParseGlob(caminhoAplicacao + "/" + adicionalPath))
	}

	templates.ExecuteTemplate(w, viewName, parametros)
}
