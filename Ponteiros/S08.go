//Crie uma função que receba um ponteiro para uma variável como argumento e modifique o valor da variável
//dentro da função. Certifique-se de que o ponteiro aponte para uma área de memória válida e que a memória
//seja liberada após o uso.

package main

import "fmt"

func ModificarValor(ptr *int) {
	*ptr = 42
}

func main() {
	var valor int
	ponteiro := &valor

	fmt.Println("Valor antes:", valor)

	ModificarValor(ponteiro)

	fmt.Println("Valor depois:", valor)
}
