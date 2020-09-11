package models

import (
	"fmt"
	"os"

	"github.com/tayron/agenda-contatos/bootstrap/library/database"
)

type Contato struct {
	ID           int
	Nome         string
	Departamento string
	Ramal        string
	Telefone     string
	Celular      string
	Email        string
}

// CriarTabelaContato - Cria caso não exista tabela contatos no banco
func CriarTabelaContato() {
	db := database.ObterConexao()
	defer db.Close()

	var sql string = `create table if not exists contatos (
		id integer auto_increment,
		nome varchar(255) NOT NULL,
		departamento varchar(255),
		ramal varchar(255),
		telefone varchar(255),
		celular varchar(255),
		email varchar(255),
		criacao datetime DEFAULT CURRENT_TIMESTAMP,	
		alteracao datetime ON UPDATE CURRENT_TIMESTAMP,
		PRIMARY KEY (id)		
	)`

	database.ExecutarQuery(db, sql)
}

// Gravar -
func (c Contato) Gravar() bool {
	db := database.ObterConexao()
	defer db.Close()

	var sql string = `insert into contatos 
		(nome, departamento, ramal, telefone, celular, email) values 
		(?, ?, ?, ?, ?, ?)`

	stmt, err := db.Prepare(sql)

	if err != nil {
		panic(err)
	}

	resultado, err := stmt.Exec(
		c.Nome,
		c.Departamento,
		c.Ramal,
		c.Telefone,
		c.Celular,
		c.Email)

	numeroRegistrosAlterados, err := resultado.RowsAffected()

	if err != nil {
		panic(err)
	}

	if numeroRegistrosAlterados > 0 {
		return true
	}

	return false
}

// Atualizar -
func (c Contato) Atualizar() bool {

	db := database.ObterConexao()
	defer db.Close()

	var sql string = `UPDATE contatos SET 
	nome = ?, departamento = ?, ramal = ?, telefone = ?, celular = ?, email = ? 
	WHERE id = ?`

	stmt, err := db.Prepare(sql)

	if err != nil {
		panic(err)
	}

	resultado, err := stmt.Exec(
		c.Nome,
		c.Departamento,
		c.Ramal,
		c.Telefone,
		c.Celular,
		c.Email,
		c.ID)

	if err != nil {
		panic(err)
	}

	_, err = resultado.RowsAffected()

	if err != nil {
		return false
	}

	return true
}

// Excluir -
func (c Contato) Excluir() bool {
	db := database.ObterConexao()
	defer db.Close()

	stmt, _ := db.Prepare("DELETE FROM contatos WHERE id = ?")
	var _, err = stmt.Exec(c.ID)

	if err != nil {
		return false
	}

	return true
}

// BuscarTodos - Retorna todos os contatos
func (c Contato) BuscarTodos() []Contato {

	db := database.ObterConexao()
	defer db.Close()

	var sql string = `SELECT id, nome, departamento, ramal, telefone, celular, email
		FROM contatos ORDER BY nome ASC`

	rows, _ := db.Query(sql)
	defer rows.Close()

	var listaContatos []Contato
	for rows.Next() {

		var contatoModel Contato

		rows.Scan(&contatoModel.ID,
			&contatoModel.Nome,
			&contatoModel.Departamento,
			&contatoModel.Ramal,
			&contatoModel.Telefone,
			&contatoModel.Celular,
			&contatoModel.Email)

		listaContatos = append(listaContatos, contatoModel)
	}

	return listaContatos
}

// BuscarTodosFiltrandoPorNome - Retorna todos os contatos filtrando por nome
func (c Contato) BuscarTodosFiltrandoPorNome(filtroNome string, offset int) []Contato {

	db := database.ObterConexao()
	defer db.Close()

	var sql string = `SELECT id, nome, departamento, ramal, telefone, celular, email
		FROM contatos WHERE nome LIKE ? ORDER BY id DESC LIMIT ? OFFSET ?`

	filtroNomeLike := fmt.Sprintf("%%%s%%", filtroNome)

	numeroRegistro := os.Getenv("NUMERO_REGISTRO_POR_PAGINA")
	rows, err := db.Query(sql, filtroNomeLike, numeroRegistro, offset)

	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var listaContatos []Contato
	for rows.Next() {

		var contatoModel Contato

		rows.Scan(&contatoModel.ID,
			&contatoModel.Nome,
			&contatoModel.Departamento,
			&contatoModel.Ramal,
			&contatoModel.Telefone,
			&contatoModel.Celular,
			&contatoModel.Email)

		listaContatos = append(listaContatos, contatoModel)
	}

	return listaContatos
}

// BuscarPorID - Busca contato por ID
func (c Contato) BuscarPorID() Contato {

	db := database.ObterConexao()
	defer db.Close()

	var sql string = `SELECT id, nome, departamento, ramal, telefone, celular, email
		FROM contatos WHERE id = ? ORDER BY id DESC`

	rows, _ := db.Query(sql, c.ID)
	defer rows.Close()

	var contatoModel Contato
	for rows.Next() {
		rows.Scan(&contatoModel.ID,
			&contatoModel.Nome,
			&contatoModel.Departamento,
			&contatoModel.Ramal,
			&contatoModel.Telefone,
			&contatoModel.Celular,
			&contatoModel.Email)

		return contatoModel
	}

	return contatoModel
}

// ObterNumeroContatosPorNome -
func ObterNumeroContatosPorNome(filtroNome string) int {
	db := database.ObterConexao()
	defer db.Close()

	var sql string = `SELECT count(0) FROM contatos WHERE nome LIKE ?`

	filtroNomeLike := fmt.Sprintf("%%%s%%", filtroNome)

	rows, err := db.Query(sql, filtroNomeLike)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var numero int = 0
	for rows.Next() {
		rows.Scan(&numero)
	}

	return numero
}
