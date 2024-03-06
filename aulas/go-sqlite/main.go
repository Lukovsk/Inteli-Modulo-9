package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, _ := sql.Open("sqlite3", "./exemplo.db")
	defer db.Close() // Defer Closing the database

	// Criando a tabla
	sqlStmt := `
  CREATE TABLE IF NOT EXISTS usuarios
  (id INTEGER PRIMARY KEY, nome TEXT, idade INTEGER)
  `
	// Preparando o sql statement de forma segura
	command, err := db.Prepare(sqlStmt)
	if err != nil {
		log.Fatal(err.Error())
	}

	// Executando o comando sql
	command.Exec()

	// Criando uma função para inserir usuários
	insertUser := func(db *sql.DB, nome string, idade int) {
		stmt := `INSERT INTO usuarios(nome, idade) VALUES (?, ?)`
		statement, err := db.Prepare(stmt)
		if err != nil {
			log.Fatalln(err.Error())
		}
		_, err = statement.Exec(nome, idade)
		if err != nil {
			log.Fatalln(err.Error())
		}
	}

	insertUser(db, "Alice", 30)
	insertUser(db, "Bob", 25)

	displayUsers(db)
}

func displayUsers(db *sql.DB) {
	row, err := db.Query("SELECT * FROM usuarios ORDER BY nome")
	if err != nil {
		log.Fatal(err)
	}
	defer row.Close()
	for row.Next() {
		var id int
		var name string
		var age string
		row.Scan(&id, &name, &age)
		log.Println("Usuário: ", id, " ", name, " ", age)
	}
}
