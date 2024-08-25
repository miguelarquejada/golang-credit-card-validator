package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func RunLuhnAlgorithm(number string) bool {
	var numberSlice []int
	for i := 0; i < len(number); i++ {
		char := number[i]
		num, err := strconv.Atoi(string(char))
		if err != nil {
			log.Fatalf("Erro na conversão: %v", err)
		}

		numberSlice = append(numberSlice, num)
	}

	var sum int = 0
	multiplyNext := false
	for i := len(numberSlice) - 1; i >= 0; i-- {
		num := numberSlice[i]
		if multiplyNext {
			num *= 2
			if num > 9 {
				num = num - 9
			}
		}
		sum += num
		multiplyNext = !multiplyNext
	}

	return sum != 0 && sum%10 == 0
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("\nDigite um número de cartão de crédito, apenas números (ou 'sair' para terminar): ")
		scanner.Scan()
		input := scanner.Text()

		// Verifica se o usuário deseja sair
		if strings.ToLower(input) == "sair" {
			fmt.Println("Saindo...")
			break
		}

		// Possuir somente números
		if _, err := strconv.Atoi(input); err != nil {
			fmt.Println("Deve possuir apenas números.")
			continue
		}

		// Ter entre 12 e 19 dígitos
		if len(input) > 19 || len(input) < 12 {
			fmt.Println("Número deve ter entre 12 e 19 dígitos.")
			continue
		}

		// todo: Os 6 primeiros dígitos devem indicar um IIN (Issuer identification number) válido

		// Passar na validação do Algoritmo de Luhn
		if RunLuhnAlgorithm(input) {
			fmt.Println("Número VÁLIDO!")
		} else {
			fmt.Println("Número INVÁLIDO!")
		}
	}
}
