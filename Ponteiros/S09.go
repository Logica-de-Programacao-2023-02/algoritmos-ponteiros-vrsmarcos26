//Implemente uma função que receba um ponteiro para uma struct "Livro" com campos "Título", "Autor" e "Preço",
//e que adicione um desconto de 10% no preço do livro.

package main

import "fmt"

type Livro struct {
	Titulo string
	Autor  string
	Preco  float64
}

func AplicarDesconto(livro *Livro) {
	livro.Preco *= 0.9 // Reduz o preço em 10%
}

func main() {
	livro := Livro{
		Titulo: "Livro1",
		Autor:  "Autor1",
		Preco:  50.0,
	}

	fmt.Println("Livro antes do desconto:", livro)

	AplicarDesconto(&livro)

	fmt.Println("Livro depois do desconto:", livro)
}
