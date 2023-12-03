//Escreva uma função em Go que receba um ponteiro para uma string e atualize o valor da string para seu reverso.

package main

func InverterString(ptr *string) {
	runes := []rune(*ptr)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	*ptr = string(runes)
}
