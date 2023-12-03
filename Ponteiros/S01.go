//Escreva uma função que receba um ponteiro para um inteiro e um valor inteiro n.
//A função deve atualizar o valor do inteiro para a soma dos n primeiros números naturais.

package main

import "fmt"

package Ponteiros

import "fmt"

func SomaDosNaturais(ptr *int, n int) {
	if n < 0 {
		fmt.Print("O valor de n deve ser um número não negativo.")
		return
	}

	sum := 0
	for i := 1; i <= n; i++ {
		sum += i
	}

	*ptr = sum
}

