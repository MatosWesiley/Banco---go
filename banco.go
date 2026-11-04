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

	for {

		msotrarmenu()

		fmt.Println("Escolha uma opção:")
		fmt.Scan(&opcao)

		switch opcao {

		case 1:
			consultarSaldo(saldo)

		case 2:

			saldo = realizarDeposito(saldo)

		case 3:

			saldo = realizarSaque(saldo)

		case 0:
			fmt.Println("Saindo do Banco")
			return

		default:
			fmt.Println("Opcao invalida!")
		}

	}
}

func consultarSaldo(saldo float64) {

	fmt.Printf("Seu saldo atual é de R$ %.2f\n\n", saldo)
}

func realizarDeposito(saldo float64) float64 {

	var deposito float64

	fmt.Println("Qual sera o valor do deposito:")
	fmt.Scan(&deposito)

	if deposito > 0 {

		saldo = saldo + deposito

		fmt.Println("Deposito realizado com sucesso!")
		fmt.Printf("Seu novo saldo e de R$ %.2f\n\n", saldo)

	} else {

		fmt.Println("Valor invalido, insira outro valor!")
	}

	return saldo
}

func realizarSaque(saldo float64) float64 {

	var saque float64

	fmt.Println("Qual o valor do saque:")
	fmt.Scan(&saque)

	if saque > 0 {

		if saldo >= saque {

			saldo = saldo - saque

			fmt.Printf("Saque no valor de R$ %.2f ", saque)
			fmt.Print("realizado com sucesso!\n\n")

		} else {

			fmt.Println("Saldo Insuficiente!")

		}

	} else {

		fmt.Println("Saque Invalido!")

	}

	return saldo

}
