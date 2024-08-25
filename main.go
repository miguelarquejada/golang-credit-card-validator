package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type CreditCardRequest struct {
	Number string `json:"number"`
}

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

func ValidateCreditCardHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	var req CreditCardRequest
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&req); err != nil {
		http.Error(w, "Erro na leitura do JSON", http.StatusBadRequest)
		return
	}

	if len(req.Number) < 12 || len(req.Number) > 19 {
		http.Error(w, "Número deve ter entre 12 e 19 dígitos", http.StatusBadRequest)
		return
	}

	if _, err := strconv.Atoi(req.Number); err != nil {
		http.Error(w, "Deve possuir apenas números", http.StatusBadRequest)
		return
	}

	if RunLuhnAlgorithm(req.Number) {
		w.WriteHeader(http.StatusOK)
		return
	}

	http.Error(w, "Número INVÁLIDO!", http.StatusBadRequest)
}

func main() {
	http.HandleFunc("/validate", ValidateCreditCardHandler)

	fmt.Println("Servidor iniciado na porta 8080")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("Erro ao iniciar servidor: %v\n", err)
	}
}
