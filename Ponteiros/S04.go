//Escreva uma função em Go que receba um ponteiro para uma variável inteira e atualize o valor da variável
//para a soma dos valores dos seus dois últimos dígitos (por exemplo, se o valor inicial da variável for 1234,
//o novo valor será 3+4=7).

package main

func SomaUltimosDigitos(ptr *int) {
	numero := *ptr

	digito1 := numero % 10
	numero /= 10
	digito2 := numero % 10

	*ptr = digito1 + digito2
}
