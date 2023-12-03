//Escreva uma função em Go que receba um ponteiro para um struct Livro com campos título e autor,
//e altere o título do livro para "Desconhecido" se o autor for "Anônimo".

package main

import "fmt"

type Livro struct {
	Titulo string
	Autor  string
}

func AlterarTitulo(livro *Livro) {
	if livro.Autor == "Anônimo" {
		livro.Titulo = "Desconhecido"
	}
}

func main() {
	livro := Livro{
		Titulo: "Livro1",
		Autor:  "Anônimo",
	}

	fmt.Println("Livro antes:", livro)

	AlterarTitulo(&livro)

	fmt.Println("Livro depois:", livro)
}
