package bootstrap

import (
	"net/http"

	"github.com/tayron/agenda-contatos/bootstrap/library/util"

	"github.com/gorilla/mux"
	"github.com/tayron/agenda-contatos/controllers"
)

// CarregarRotas - Método que define e carrega as rodas da aplicação
func CarregarRotas() {
	r := mux.NewRouter()

	diretorioAplicacao := util.ObterCaminhoDiretorioAplicacao()
	pastaPublic := diretorioAplicacao + "/public/"

	s := http.StripPrefix("/public/", http.FileServer(http.Dir(pastaPublic)))
	r.PathPrefix("/public/").Handler(s)

	r.HandleFunc("/contatos/listar", controllers.ListarContatosPublico).Methods("GET")
	r.HandleFunc("/", controllers.ListarContato).Methods("GET")

	r.HandleFunc("/contatos", controllers.ListarContato).Methods("GET")
	r.HandleFunc("/contato/cadastrar", controllers.CadastrarContato)
	r.HandleFunc("/contato/editar/{id:[0-9]+}", controllers.EditarContato)
	r.HandleFunc("/contato/excluir", controllers.ExcluirContato).Methods("POST")

	r.HandleFunc("/login", controllers.Login).Methods("GET")
	r.HandleFunc("/login", controllers.Login).Methods("POST")
	r.HandleFunc("/logout", controllers.Logout).Methods("GET")

	http.Handle("/", r)
}
