package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// Entidade representa uma linha do arquivo CSV
type Entidade struct {
	Nome      string
	Idade     int
	Pontuacao int
}

// PorNome implementa a interface sort.Interface para ordenar por nome
type PorNome []Entidade

func (p PorNome) Len() int      { return len(p) }
func (p PorNome) Swap(i, j int) { p[i], p[j] = p[j], p[i] }
func (p PorNome) Less(i, j int) bool {
	// Critério 1: nome iniciado com letra maiúscula
	if strings.ToLower(p[i].Nome) != strings.ToLower(p[j].Nome) {
		return strings.ToLower(p[i].Nome) < strings.ToLower(p[j].Nome)
	}

	// Critério 2: idade menor
	return p[i].Idade < p[j].Idade
}

// LerArquivo lê o arquivo CSV e retorna um slice de Entidade
func LerArquivo(nomeArquivo string) ([]Entidade, error) {
	arquivo, err := os.Open(nomeArquivo)
	if err != nil {
		return nil, err
	}
	defer arquivo.Close()

	leitorCSV := csv.NewReader(arquivo)
	linhas, err := leitorCSV.ReadAll()
	if err != nil {
		return nil, err
	}

	// Verificar se a primeira linha é "Nome,Idade,Pontuação"
	if len(linhas) == 0 || linhas[0][0] != "Nome" || linhas[0][1] != "Idade" || (linhas[0][2] != "Pontuacao" && linhas[0][2] != "Pontuação") {
		return nil, fmt.Errorf("o arquivo não possui a estrutura esperada")
	}

	var entidades []Entidade
	for i, linha := range linhas {
		// Ignorar a primeira linha
		if i == 0 {
			continue
		}

		idade, _ := strconv.Atoi(linha[1])
		pontuacao, _ := strconv.Atoi(linha[2])
		entidade := Entidade{Nome: linha[0], Idade: idade, Pontuacao: pontuacao}
		entidades = append(entidades, entidade)
	}

	return entidades, nil
}

// EscreverArquivo escreve o slice de Entidade ordenado em um novo arquivo CSV
func EscreverArquivo(nomeArquivo string, entidades []Entidade) error {
	arquivo, err := os.Create(nomeArquivo)
	if err != nil {
		return err
	}
	defer arquivo.Close()

	escritorCSV := csv.NewWriter(arquivo)
	defer escritorCSV.Flush()

	// Escrever a primeira linha
	primeiraLinha := []string{"Nome", "Idade", "Pontuação"}
	if err := escritorCSV.Write(primeiraLinha); err != nil {
		return err
	}

	// Escrever as demais linhas
	for _, entidade := range entidades {
		linha := []string{entidade.Nome, strconv.Itoa(entidade.Idade), strconv.Itoa(entidade.Pontuacao)}
		if err := escritorCSV.Write(linha); err != nil {
			return err
		}
	}

	return nil
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Uso: go run main.go <arquivo-origem.csv> <arquivo-destino.csv>")
		os.Exit(1)
	}

	arquivoOrigem := os.Args[1]
	arquivoDestino := os.Args[2]

	entidades, err := LerArquivo(arquivoOrigem)
	if err != nil {
		fmt.Printf("Erro ao ler o arquivo de origem: %v\n", err)
		os.Exit(1)
	}

	// Ordenar por nome (considerando letras minúsculas) e, em seguida, por idade
	sort.Sort(PorNome(entidades))
	if err := EscreverArquivo(arquivoDestino, entidades); err != nil {
		fmt.Printf("Erro ao escrever o arquivo ordenado: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Programa executado com sucesso.")
}
