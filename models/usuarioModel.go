package models

import (
	"github.com/tayron/agenda-contatos/bootstrap/library/database"
)

type Usuario struct {
	ID                int
	Nome              string
	Login             string
	Senha             string
	Ativo             bool
	PermiteExclusao   bool
	permiteEdicao     bool
	PermiteSerListado bool
}

// CriarTabelaUsuario - Cria caso não exista tabela usuaários no banco
func CriarTabelaUsuario() {
	db := database.ObterConexao()
	defer db.Close()

	var sql string = `create table if not exists usuarios (
		id integer auto_increment,
		nome varchar(255) NOT NULL,
		login varchar(255) NOT NULL UNIQUE,		
		senha varchar(255),		
		ativo bool DEFAULT 0,
		permite_exclusao bool DEFAULT 1,
		permite_edicao bool DEFAULT 1,
		permite_ser_listado bool DEFAULT 1,
		criacao datetime DEFAULT CURRENT_TIMESTAMP,	
		alteracao datetime ON UPDATE CURRENT_TIMESTAMP,
		PRIMARY KEY (id)
	)`

	database.ExecutarQuery(db, sql)
}

// CriarUsuarioAdministrador - Cria usuário default do sistema
func CriarUsuarioAdministrador() {
	var usuarioModel Usuario

	listaUsuarios := usuarioModel.BuscarTodos()

	if len(listaUsuarios) == 0 {
		db := database.ObterConexao()
		defer db.Close()

		var sql string = `insert into usuarios 
			(nome, login, senha, ativo, permite_exclusao, permite_edicao, permite_ser_listado) 
			values (?, ?, ?, ?, ?, ?, ?)`

		stmt, _ := db.Prepare(sql)

		usuarioModel := Usuario{
			Nome:              "Tayron",
			Login:             "tayron",
			Senha:             "$2a$14$ZN3eWRZs30egm9pwDOucVeBBu28LMoou4JCTf0EsU2pzLCLyshYnu",
			Ativo:             true,
			PermiteExclusao:   false,
			permiteEdicao:     false,
			PermiteSerListado: false,
		}

		_, err := stmt.Exec(
			usuarioModel.Nome,
			usuarioModel.Login,
			usuarioModel.Senha,
			usuarioModel.Ativo,
			usuarioModel.PermiteExclusao,
			usuarioModel.permiteEdicao,
			usuarioModel.PermiteSerListado,
		)

		if err != nil {
			panic(err)
		}
	}
}

// BuscarTodos -Retorna todos os usuários independente de status
func (u Usuario) BuscarTodos() []Usuario {

	db := database.ObterConexao()
	defer db.Close()

	var sql string = `SELECT id, nome, login, ativo
		FROM usuarios ORDER BY id DESC`

	rows, _ := db.Query(sql)
	defer rows.Close()

	var listaUsuarios []Usuario
	for rows.Next() {

		var usuarioModel Usuario

		rows.Scan(&usuarioModel.ID,
			&usuarioModel.Nome,
			&usuarioModel.Login,
			&usuarioModel.Ativo)

		listaUsuarios = append(listaUsuarios, usuarioModel)
	}

	return listaUsuarios
}

// BuscarPorLoginStatus - Busca usuario por login e status
func (u Usuario) BuscarPorLoginStatus() Usuario {
	db := database.ObterConexao()
	defer db.Close()

	var sql string = `SELECT id, nome, login, senha, ativo
		FROM usuarios WHERE login = ? AND ativo = ?`

	rows, _ := db.Query(sql, u.Login, u.Ativo)
	defer rows.Close()

	var usuarioModel Usuario
	for rows.Next() {
		rows.Scan(&usuarioModel.ID,
			&usuarioModel.Nome,
			&usuarioModel.Login,
			&usuarioModel.Senha,
			&usuarioModel.Ativo)
		return usuarioModel
	}

	return usuarioModel
}
