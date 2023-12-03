//Escreva uma função em Go que receba um ponteiro para um struct Conta com campos saldo e titular,
//e adicione um valor ao saldo da conta.

package main

type Conta struct {
	Saldo   float64
	Titular string
}

func AdicionarSaldo(conta *Conta, valor float64) {
	conta.Saldo += valor
}
