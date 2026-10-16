package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql" //usar o pacote de forma implicita, sem precisar chamar funções dele
)

func main() {
	stringConexao := "golang:golang@tcp(127.0.0.1:3306)/devbook?charset=utf8&parseTime=True&loc=Local"

	db, erro := sql.Open("mysql", stringConexao)
	if erro != nil {
		log.Fatal(erro)
	}

	defer db.Close()

	//reaproveitando a var erro
	if erro = db.Ping(); erro != nil {
		fmt.Println("Dentro do ping")
		log.Fatal(erro)
	}

	fmt.Println("Conexao esta aberta")

	linhas, erro := db.Query("select * from usuarios")

	if erro != nil {
		log.Fatal(erro)
	}
	defer linhas.Close()
	fmt.Println(linhas)

}
