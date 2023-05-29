package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	fmt.Println("Bem-vindo ao jogo de adivinhar o número!")
	fmt.Println("Tente adivinhar um número entre 1 e 100.")

	var historicoTentativas []int

	for {
		numeroAleatorio := rand.Intn(100) + 1
		tentativas := 0

		for {
			fmt.Print("Digite um número: ")
			var palpite int
			_, err := fmt.Scanf("%d", &palpite)

			if err != nil {
				fmt.Println("Entrada inválida. Por favor, digite um número válido.")
				continue
			}

			tentativas++
			historicoTentativas = append(historicoTentativas, palpite)

			if palpite < numeroAleatorio {
				fmt.Println("Palpite baixo. Tente novamente!")
			} else if palpite > numeroAleatorio {
				fmt.Println("Palpite alto. Tente novamente!")
			} else {
				fmt.Printf("Parabéns! Você acertou o número em %d tentativas!\n", tentativas)
				break
			}
		}

		fmt.Print("Deseja jogar novamente? (s/n): ")
		var resposta string
		_, err := fmt.Scanf("%s", &resposta)

		if err != nil {
			fmt.Println("Entrada inválida. Encerrando o jogo.")
			break
		}

		if resposta != "s" && resposta != "S" {
			fmt.Println("Histórico de tentativas:")
			for i, tentativa := range historicoTentativas {
				fmt.Printf("Tentativa %d: %d\n", i+1, tentativa)
			}
			break
		}
	}

	fmt.Println("Obrigado por jogar! Até a próxima!")
}
