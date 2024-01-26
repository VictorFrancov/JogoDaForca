package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Digite a palavra a ser adivinhada.")
		os.Exit(1)
	}

	palavra := os.Args[1]

	letrasCorretas := make([]string, 0)

	maxTentativas := 6

	tentativas := 0

	palavraOculta := make([]string, len(palavra))
	for i := range palavraOculta {
		palavraOculta[i] = "_"
	}

	for {
		fmt.Println(strings.Join(palavraOculta, " "))

		fmt.Print("Digite uma letra: ")
		letra, err := lerEntrada()
		if err != nil {
			fmt.Println("Erro ao ler a entrada.")
			os.Exit(1)
		}

		if letraJaTentada(letra, letrasCorretas) {
			fmt.Println("Você já tentou essa letra. Tente outra.")
			continue
		}

		if strings.Contains(palavra, letra) {
			fmt.Println("Letra correta!")
			letrasCorretas = append(letrasCorretas, letra)

			for i, char := range palavra {
				if string(char) == letra {
					palavraOculta[i] = letra
				}
			}
		} else {
			if tentativas == maxTentativas {
				fmt.Println("Você excedeu o número máximo de tentativas. A palavra era:", palavra)
				break
			}
			fmt.Println("Letra incorreta. Tente novamente.")
			tentativas++
		}

		if strings.Join(palavraOculta, "") == palavra {
			fmt.Println("Parabéns! Você acertou a palavra:", palavra)
			break
		}

	}
}

func lerEntrada() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	entrada, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(strings.ToLower(entrada)), nil
}

func letraJaTentada(letra string, letrasTentadas []string) bool {
	for _, l := range letrasTentadas {
		if l == letra {
			return true
		}
	}
	return false
}
