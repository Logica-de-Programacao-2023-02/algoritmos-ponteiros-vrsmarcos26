//Implemente uma função que receba um ponteiro para uma slice e seu tamanho e preencha-o com os n primeiros números primos.

package main

import (
	"fmt"
	"math"
)

func PreencherComPrimos(slice *[]int, tamanho int) {
	if tamanho <= 0 {
		fmt.Println("O tamanho deve ser um número positivo.")
		return
	}

	*slice = []int{}

	for i := 2; len(*slice) < tamanho; i++ {
		if primo(i) {
			*slice = append(*slice, i)
		}
	}
}

func primo(num int) bool {
	if num < 2 {
		return false
	}

	for i := 2; i <= int(math.Sqrt(float64(num))); i++ {
		if num%i == 0 {
			return false
		}
	}

	return true
}

func main() {
	var primos []int
	tamanho := 5

	PreencherComPrimos(&primos, tamanho)

	fmt.Printf("Os %d primeiros números primos são: %v\n", tamanho, primos)
}
