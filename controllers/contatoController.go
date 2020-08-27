package controllers

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/tayron/agenda-contatos/bootstrap/library/template"
	"github.com/tayron/agenda-contatos/models"
	"github.com/gorilla/mux"
)

func ListarContatosPublico(w http.ResponseWriter, r *http.Request) {
	filtroNome := strings.TrimSpace(r.FormValue("nome"))

	var listaContatos []models.Contato
	ContatoModel := models.Contato{}

	if filtroNome != "" {
		listaContatos = ContatoModel.BuscarTodosFiltrandoPorNome(filtroNome)
	} else {
		listaContatos = ContatoModel.BuscarTodos()
	}

	var Contatos = struct {
		ListaContatos []models.Contato
		FiltroNome    string
	}{
		ListaContatos: listaContatos,
		FiltroNome:    filtroNome,
	}

	parametros := template.Parametro{
		System:    template.ObterInformacaoSistema(w, r),
		Parametro: Contatos,
	}

	template.LoadView(w, "template/contato/*.html", "listarContatosPublicoPage", parametros)
}

// ListarContato -
func ListarContato(w http.ResponseWriter, r *http.Request) {
	ValidarSessao(w, r)

	filtroNome := strings.TrimSpace(r.FormValue("nome"))

	var listaContatos []models.Contato
	ContatoModel := models.Contato{}

	if filtroNome != "" {
		listaContatos = ContatoModel.BuscarTodosFiltrandoPorNome(filtroNome)
	} else {
		listaContatos = ContatoModel.BuscarTodos()
	}

	var Contatos = struct {
		ListaContatos []models.Contato
		FiltroNome    string
	}{
		ListaContatos: listaContatos,
		FiltroNome:    filtroNome,
	}

	parametros := template.Parametro{
		System:    template.ObterInformacaoSistema(w, r),
		Parametro: Contatos,
	}

	template.LoadView(w, "template/contato/*.html", "listarContatosPage", parametros)
}

// CadastrarContato -
func CadastrarContato(w http.ResponseWriter, r *http.Request) {
	ValidarSessao(w, r)

	flashMessage := template.FlashMessage{}

	if r.Method == "POST" {
		ContatoEntidade := models.Contato{
			Nome:         r.FormValue("nome"),
			Departamento: r.FormValue("departamento"),
			Ramal:        r.FormValue("ramal"),
			Telefone:     r.FormValue("telefone"),
			Celular:      r.FormValue("celular"),
			Email:        r.FormValue("email"),
		}

		retornoGravacao := ContatoEntidade.Gravar()

		if retornoGravacao == true {
			flashMessage.Type, flashMessage.Message = template.ObterTipoMensagemGravacaoSucesso()
		} else {
			flashMessage.Type, flashMessage.Message = template.ObterTipoMensagemGravacaoErro()
		}
	}

	parametros := template.Parametro{
		System:       template.ObterInformacaoSistema(w, r),
		FlashMessage: flashMessage,
	}

	template.LoadView(w, "template/contato/*.html", "cadastrarContatoPage", parametros)
}

// EditarContato -
func EditarContato(w http.ResponseWriter, r *http.Request) {
	ValidarSessao(w, r)

	parametrosURL := mux.Vars(r)
	id, _ := strconv.Atoi(parametrosURL["id"])

	flashMessage := template.FlashMessage{}

	if r.Method == "POST" {
		ContatoModel := models.Contato{
			ID:           id,
			Nome:         r.FormValue("nome"),
			Departamento: r.FormValue("departamento"),
			Ramal:        r.FormValue("ramal"),
			Telefone:     r.FormValue("telefone"),
			Celular:      r.FormValue("celular"),
			Email:        r.FormValue("email"),
		}

		retornoGravacao := ContatoModel.Atualizar()

		if retornoGravacao == true {
			flashMessage.Type, flashMessage.Message = template.ObterTipoMensagemGravacaoSucesso()
		} else {
			flashMessage.Type, flashMessage.Message = template.ObterTipoMensagemGravacaoErro()
		}
	}

	ContatoModel := models.Contato{
		ID: id,
	}

	contato := ContatoModel.BuscarPorID()

	if contato.ID == 0 {
		http.Redirect(w, r, "/", 302)
	}

	var Contato = struct {
		Contato models.Contato
	}{
		Contato: contato,
	}

	parametros := template.Parametro{
		System:       template.ObterInformacaoSistema(w, r),
		FlashMessage: flashMessage,
		Parametro:    Contato,
	}

	template.LoadView(w, "template/contato/*.html", "editarContatoPage", parametros)
}

// ExcluirContato -
func ExcluirContato(w http.ResponseWriter, r *http.Request) {
	ValidarSessao(w, r)

	idContato, _ := strconv.Atoi(r.FormValue("id"))
	flashMessage := template.FlashMessage{}

	ContatoModel := models.Contato{
		ID: idContato,
	}

	retornoExclusao := ContatoModel.Excluir()

	if retornoExclusao == true {
		flashMessage.Type, flashMessage.Message = template.ObterTipoMensagemExclusaoSucesso()
	} else {
		flashMessage.Type, flashMessage.Message = template.ObterTipoMensagemExclusaoErro()
	}

	var Contatos = struct {
		ListaContatos []models.Contato
		FiltroNome    string
	}{
		ListaContatos: ContatoModel.BuscarTodos(),
		FiltroNome:    "",
	}

	parametros := template.Parametro{
		System:       template.ObterInformacaoSistema(w, r),
		FlashMessage: flashMessage,
		Parametro:    Contatos,
	}

	template.LoadView(w, "template/contato/*.html", "listarContatosPage", parametros)
}
