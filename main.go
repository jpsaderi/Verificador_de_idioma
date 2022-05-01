package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func pesquisar(dicionario map[string]*list.List, entrada []string) int {
	var acertos int = 0
	for i := range entrada { // Percorre as palavras da entrada
		primeira_letra := string(entrada[i][0])
		if _, ok := dicionario[primeira_letra]; ok { // Se existe essa chave, atribui a lista
			lista := dicionario[primeira_letra]
			for elemento := lista.Front(); elemento != nil; elemento = elemento.Next() {
				if entrada[i] == elemento.Value { // Se a palavra da entrada for igual a do arquivo
					acertos += 1
					break
				}
			}
		}
	}

	return acertos
}

func main() {
	// Espanhol
	file_espanhol, err := os.Open("Entrada/espanhol.txt")
	if err != nil {
		panic(err)
	}

	f_espanhol := bufio.NewScanner(file_espanhol)
	f_espanhol.Split(bufio.ScanLines)

	dic_espanhol := make(map[string]*list.List)

	for f_espanhol.Scan() {
		palavra := f_espanhol.Text()

		if _, ok := dic_espanhol[string(palavra[0])]; !ok { //Se não existe essa chave, cria ela
			dic_espanhol[string(palavra[0])] = list.New()
		}
		dic_espanhol[string(palavra[0])].PushBack(palavra) //Adiciona essa palavra ao map na chave da primeira letra dessa palavra
	}
	file_espanhol.Close()

	//Frances
	file_frances, err := os.Open("Entrada/frances.txt")
	if err != nil {
		panic(err)
	}

	f_frances := bufio.NewScanner(file_frances)
	f_frances.Split(bufio.ScanLines)

	dic_frances := make(map[string]*list.List)

	for f_frances.Scan() {
		palavra := f_frances.Text()

		if _, ok := dic_frances[string(palavra[0])]; !ok { //Se não existe essa chave, cria ela
			dic_frances[string(palavra[0])] = list.New()
		}
		dic_frances[string(palavra[0])].PushBack(palavra) //Adiciona essa palavra ao map na chave da primeira letra dessa palavra
	}
	file_frances.Close()

	//Ingles
	file_ingles, err := os.Open("Entrada/ingles.txt")
	if err != nil {
		panic(err)
	}

	f_ingles := bufio.NewScanner(file_ingles)
	f_ingles.Split(bufio.ScanLines)

	dic_ingles := make(map[string]*list.List)

	for f_ingles.Scan() {
		palavra := f_ingles.Text()

		if _, ok := dic_ingles[string(palavra[0])]; !ok { //Se não existe essa chave, cria ela
			dic_ingles[string(palavra[0])] = list.New()
		}
		dic_ingles[string(palavra[0])].PushBack(palavra) //Adiciona essa palavra ao map na chave da primeira letra dessa palavra
	}
	file_ingles.Close()

	//Portugues
	file_portugues, err := os.Open("Entrada/portugues.txt")
	if err != nil {
		panic(err)
	}

	f_portugues := bufio.NewScanner(file_portugues)
	f_portugues.Split(bufio.ScanLines)

	dic_portugues := make(map[string]*list.List)

	for f_portugues.Scan() {
		palavra := f_portugues.Text()

		if _, ok := dic_portugues[string(palavra[0])]; !ok { //Se não existe essa chave, cria ela
			dic_portugues[string(palavra[0])] = list.New()
		}
		dic_portugues[string(palavra[0])].PushBack(palavra) //Adiciona essa palavra ao map na chave da primeira letra dessa palavra
	}
	file_portugues.Close()

	// Aqui recebe o input e separa em uma lista de strings
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

	re := regexp.MustCompile("[^a-zA-ZáéíóúàèìòùâêîôûãõñçÁÉÍÓÚÀÈÌÒÙÂÊÎÔÛÃÕÑÇ\\s]") // Mantem apenas as letras e caracteres especiais em letras
	input = re.ReplaceAllString(input, "")                                         // Substitui as letras nao correspondentes por vazio no texto
	input = strings.ReplaceAll(input, "  ", " ")                                   // Remove espaços duplos

	// Divide palavras do input em um array.

	input_slice := strings.Split(input, " ")

	acertos_espanhol = pesquisar(dic_espanhol, input_slice)
	acertos_frances = pesquisar(dic_frances, input_slice)
	acertos_ingles = pesquisar(dic_ingles, input_slice)
	acertos_portugues = pesquisar(dic_portugues, input_slice)

	resultado_espanhol = float64(acertos_espanhol) / float64(len(input_slice))
	resultado_ingles = float64(acertos_ingles) / float64(len(input_slice))
	resultado_frances = float64(acertos_frances) / float64(len(input_slice))
	resultado_portugues = float64(acertos_portugues) / float64(len(input_slice))

	fmt.Println("Probabilidade de ser espanhol:", resultado_espanhol)
	fmt.Println("Probabilidade de ser ingles:", resultado_ingles)
	fmt.Println("Probabilidade de ser frances:", resultado_frances)
	fmt.Println("Probabilidade de ser portugues:", resultado_portugues)

}
