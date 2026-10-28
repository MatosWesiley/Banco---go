package main

import "fmt"

func msotrarmenu() {

	fmt.Println("===== BANCO GO =====")
	fmt.Println("1 - Consultar saldo")
	fmt.Println("2 - Depositar")
	fmt.Println("3 - Realizar saque")
	fmt.Println("0- Sair")

}

func main() {

	var saldo float64 = 0
	var opcao int
	var deposito float64
	var saque float64

	for {

		msotrarmenu()

		fmt.Println("Escolha uma opcao:")
		fmt.Scan(&opcao)

		switch opcao {

		case 1:
			fmt.Println("Seu saldo atual e de R$", saldo)

		case 2:
			fmt.Println("Qual sera o valor do deposito:")
			fmt.Scan(&deposito)

			saldo = saldo + deposito

		case 3:
			fmt.Println("Qual o valor do saque:")
			fmt.Scan(&saque)

			saldo = saldo - saque

		case 0:
			fmt.Println("Saindo do Banco")
			return

		default:
			fmt.Println("Opcao invalida!")
		}

	}
}
