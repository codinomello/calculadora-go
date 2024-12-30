package main

import (
	"fmt"
	"os"
)

type número interface {
	int64 | float64
}

func main() {
	fmt.Println("calculadora - go")
	fmt.Println()
	fmt.Println("1 - soma")
	fmt.Println("2 - subtração")
	fmt.Println("3 - multiplicação")
	fmt.Println("4 - divisão")
	fmt.Println("5 - sair")

	var opção int
	fmt.Print("\nescolha uma opção: ")
	if _, err := fmt.Scanln(&opção); err != nil || opção < 1 || opção > 5 {
		fmt.Println("\nopção inválida")
		os.Exit(1)
	}

	if opção == 5 {
		fmt.Println("\nsaindo do programa...")
		os.Exit(0)
	}

	var primeiro, segundo float64
	fmt.Print("\ndigite o 1° número: ")
	if _, err := fmt.Scanln(&primeiro); err != nil {
		fmt.Println("erro: entrada inválida")
		os.Exit(1)
	}

	fmt.Print("digite o 2° número: ")
	if _, err := fmt.Scanln(&segundo); err != nil {
		fmt.Println("erro: entrada inválida")
		os.Exit(1)
	}

	resultado, err := calcular(opção, primeiro, segundo)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("\nresultado: %v", resultado)
	}
}

func calcular[T número](opção int, primeiro, segundo T) (resultado T, err error) {
	switch opção {
	case 1:
		resultado = primeiro + segundo
	case 2:
		resultado = primeiro - segundo
	case 3:
		resultado = primeiro * segundo
	case 4:
		if segundo == 0 {
			return 0, fmt.Errorf("erro: divisão por zero")
		}
		resultado = primeiro / segundo
	default:
		return 0, fmt.Errorf("erro: opção inválida")
	}
	return resultado, nil
}
