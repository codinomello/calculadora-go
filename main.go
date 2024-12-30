package main

import (
	"fmt"
	"os"
)

// define uma interface para número que pode ser int64 ou float64
type número interface {
	int64 | float64
}

func main() {
	// imprime as opções do menu
	fmt.Println("calculadora - go")
	fmt.Println()
	fmt.Println("1 - soma")
	fmt.Println("2 - subtração")
	fmt.Println("3 - multiplicação")
	fmt.Println("4 - divisão")
	fmt.Println("5 - sair")

	var opção int
	// solicita ao usuário que escolha uma opção
	fmt.Print("\nescolha uma opção: ")
	if _, err := fmt.Scanln(&opção); err != nil || opção < 1 || opção > 5 {
		fmt.Println("\nopção inválida")
		os.Exit(1)
	}

	// sai do programa se o usuário escolher a opção 5
	if opção == 5 {
		fmt.Println("\nsaindo do programa...")
		os.Exit(0)
	}

	var primeiro, segundo float64
	// solicita ao usuário que digite o primeiro número
	fmt.Print("\ndigite o 1° número: ")
	if _, err := fmt.Scanln(&primeiro); err != nil {
		fmt.Println("erro: entrada inválida")
		os.Exit(1)
	}

	// solicita ao usuário que digite o segundo número
	fmt.Print("digite o 2° número: ")
	if _, err := fmt.Scanln(&segundo); err != nil {
		fmt.Println("erro: entrada inválida")
		os.Exit(1)
	}

	// realiza o cálculo com base na opção escolhida
	resultado, err := calcular(opção, primeiro, segundo)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("\nresultado: %v", resultado)
	}
}

// função genérica para realizar o cálculo com base na opção escolhida
func calcular[T número](opção int, primeiro, segundo T) (resultado T, err error) {
	switch opção {
	case 1:
		resultado = primeiro + segundo // adição
	case 2:
		resultado = primeiro - segundo // subtração
	case 3:
		resultado = primeiro * segundo // multiplicação
	case 4:
		if segundo == 0 {
			return 0, fmt.Errorf("erro: divisão por zero") // erro de divisão por zero
		}
		resultado = primeiro / segundo // divisão
	default:
		return 0, fmt.Errorf("erro: opção inválida") // erro de opção inválida
	}
	return resultado, nil
}
