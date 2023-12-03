//Escreva uma função em Go que receba um ponteiro para uma variável float64 e atualize o valor da variável
//para a média aritmética entre o seu valor atual e o valor da constante Pi.

package main

import (
	"fmt"
	"math"
)

func AtualizarMediaComPi(ptr *float64) {
	valorAtual := *ptr
	media := (valorAtual + math.Pi) / 2
	*ptr = media
}

func main() {
	valor := 3.14
	fmt.Println("Valor antes:", valor)

	AtualizarMediaComPi(&valor)

	fmt.Println("Valor depois:", valor)
}
