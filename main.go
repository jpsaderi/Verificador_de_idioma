package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func pesquisar(scanner *bufio.Scanner, input []string) int {
	var acertos int = 0

	for scanner.Scan() {
		line := scanner.Text()

		for i := 0; i < len(input); i++ {
			if line == input[i] {
				acertos++
			}
		}
	}

	return acertos
}

func main() {
	var acertos_espanhol int = 0
	var acertos_ingles int = 0
	var acertos_frances int = 0
	var acertos_portugues int = 0

	var resultado_espanhol, resultado_ingles, resultado_frances, resultado_portugues float64

	// Input do usuário.

	fmt.Println("Informe a frase que deseja saber o idioma: ")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan() // use `for scanner.Scan()` to keep reading
	input := scanner.Text()

	// Divide palavras do input em um array.

	input_slice := strings.Split(input, " ")

	// Abre arquivos de texto em cada língua e checa possível erro.

	file_espanhol, ferr_espanhol := os.Open("dados/espanhol.txt")
	if ferr_espanhol != nil {
		panic(ferr_espanhol)
	}

	file_ingles, ferr_ingles := os.Open("dados/ingles.txt")
	if ferr_ingles != nil {
		panic(ferr_ingles)
	}

	file_frances, ferr_frances := os.Open("dados/frances.txt")
	if ferr_frances != nil {
		panic(ferr_frances)
	}

	file_portugues, ferr_portugues := os.Open("dados/portugues.txt")
	if ferr_portugues != nil {
		panic(ferr_portugues)
	}

	// Cria buffer para leitura dos arquivos.

	scanner_espanhol := bufio.NewScanner(file_espanhol)
	scanner_ingles := bufio.NewScanner(file_ingles)
	scanner_frances := bufio.NewScanner(file_frances)
	scanner_portugues := bufio.NewScanner(file_portugues)

	// Pesquisa sequencial em cada arquivo.

	acertos_espanhol = pesquisar(scanner_espanhol, input_slice)
	acertos_ingles = pesquisar(scanner_ingles, input_slice)
	acertos_frances = pesquisar(scanner_frances, input_slice)
	acertos_portugues = pesquisar(scanner_portugues, input_slice)

	// Resultado.

	resultado_espanhol = float64(acertos_espanhol) / float64(len(input_slice))
	resultado_ingles = float64(acertos_ingles) / float64(len(input_slice))
	resultado_frances = float64(acertos_frances) / float64(len(input_slice))
	resultado_portugues = float64(acertos_portugues) / float64(len(input_slice))

	fmt.Println("Probabilidade de ser espanhol:", resultado_espanhol)
	fmt.Println("Probabilidade de ser ingles:", resultado_ingles)
	fmt.Println("Probabilidade de ser frances:", resultado_frances)
	fmt.Println("Probabilidade de ser portugues:", resultado_portugues)
}
