package controllers

import (
	"html/template"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
	aplicacaoTemplate "github.com/tayron/agenda-contatos/bootstrap/library/template"
	"github.com/tayron/agenda-contatos/models"
	"github.com/tayron/gopaginacao"
)

// ListarContatosPublico -
func ListarContatosPublico(w http.ResponseWriter, r *http.Request) {
	filtroNome := strings.TrimSpace(r.FormValue("nome"))

	var listaContatos []models.Contato
	ContatoModel := models.Contato{}

	numeroTotalRegistro := models.ObterNumeroContatosPorNome(filtroNome)
	htmlPaginacao, offset, err := gopaginacao.CriarPaginacao(numeroTotalRegistro, r)

	if err == nil {
		listaContatos = ContatoModel.BuscarTodosFiltrandoPorNome(filtroNome, offset)
	}

	var Contatos = struct {
		ListaContatos []models.Contato
		FiltroNome    string
		Paginacao     template.HTML
	}{
		ListaContatos: listaContatos,
		FiltroNome:    filtroNome,
		Paginacao:     template.HTML(htmlPaginacao),
	}

	parametros := aplicacaoTemplate.Parametro{
		System:    aplicacaoTemplate.ObterInformacaoSistema(w, r),
		Parametro: Contatos,
	}

	aplicacaoTemplate.LoadView(w, "template/contato/*.html", "listarContatosPublicoPage", parametros)
}

// ListarContato -
func ListarContato(w http.ResponseWriter, r *http.Request) {
	ValidarSessao(w, r)

	filtroNome := strings.TrimSpace(r.FormValue("nome"))

	var listaContatos []models.Contato
	ContatoModel := models.Contato{}

	numeroTotalRegistro := models.ObterNumeroContatosPorNome(filtroNome)
	htmlPaginacao, offset, err := gopaginacao.CriarPaginacao(numeroTotalRegistro, r)

	if err == nil {
		listaContatos = ContatoModel.BuscarTodosFiltrandoPorNome(filtroNome, offset)
	}

	var Contatos = struct {
		ListaContatos []models.Contato
		FiltroNome    string
		Paginacao     template.HTML
	}{
		ListaContatos: listaContatos,
		FiltroNome:    filtroNome,
		Paginacao:     template.HTML(htmlPaginacao),
	}

	parametros := aplicacaoTemplate.Parametro{
		System:    aplicacaoTemplate.ObterInformacaoSistema(w, r),
		Parametro: Contatos,
	}

	aplicacaoTemplate.LoadView(w, "template/contato/*.html", "listarContatosPage", parametros)
}

// CadastrarContato -
func CadastrarContato(w http.ResponseWriter, r *http.Request) {
	ValidarSessao(w, r)

	flashMessage := aplicacaoTemplate.FlashMessage{}

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
			flashMessage.Type, flashMessage.Message = aplicacaoTemplate.ObterTipoMensagemGravacaoSucesso()
		} else {
			flashMessage.Type, flashMessage.Message = aplicacaoTemplate.ObterTipoMensagemGravacaoErro()
		}
	}

	parametros := aplicacaoTemplate.Parametro{
		System:       aplicacaoTemplate.ObterInformacaoSistema(w, r),
		FlashMessage: flashMessage,
	}

	aplicacaoTemplate.LoadView(w, "template/contato/*.html", "cadastrarContatoPage", parametros)
}

// EditarContato -
func EditarContato(w http.ResponseWriter, r *http.Request) {
	ValidarSessao(w, r)

	parametrosURL := mux.Vars(r)
	id, _ := strconv.Atoi(parametrosURL["id"])

	flashMessage := aplicacaoTemplate.FlashMessage{}

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
			flashMessage.Type, flashMessage.Message = aplicacaoTemplate.ObterTipoMensagemGravacaoSucesso()
		} else {
			flashMessage.Type, flashMessage.Message = aplicacaoTemplate.ObterTipoMensagemGravacaoErro()
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

	parametros := aplicacaoTemplate.Parametro{
		System:       aplicacaoTemplate.ObterInformacaoSistema(w, r),
		FlashMessage: flashMessage,
		Parametro:    Contato,
	}

	aplicacaoTemplate.LoadView(w, "template/contato/*.html", "editarContatoPage", parametros)
}

// ExcluirContato -
func ExcluirContato(w http.ResponseWriter, r *http.Request) {
	ValidarSessao(w, r)

	idContato, _ := strconv.Atoi(r.FormValue("id"))
	flashMessage := aplicacaoTemplate.FlashMessage{}

	ContatoModel := models.Contato{
		ID: idContato,
	}

	retornoExclusao := ContatoModel.Excluir()

	if retornoExclusao == true {
		flashMessage.Type, flashMessage.Message = aplicacaoTemplate.ObterTipoMensagemExclusaoSucesso()
	} else {
		flashMessage.Type, flashMessage.Message = aplicacaoTemplate.ObterTipoMensagemExclusaoErro()
	}

	var Contatos = struct {
		ListaContatos []models.Contato
		FiltroNome    string
	}{
		ListaContatos: ContatoModel.BuscarTodos(),
		FiltroNome:    "",
	}

	parametros := aplicacaoTemplate.Parametro{
		System:       aplicacaoTemplate.ObterInformacaoSistema(w, r),
		FlashMessage: flashMessage,
		Parametro:    Contatos,
	}

	aplicacaoTemplate.LoadView(w, "template/contato/*.html", "listarContatosPage", parametros)
}
