![header-igor-projetos](https://github.com/igor-rl/assets/blob/main/img/github-projetcs-header.jpg)

# FullCycle Imersão 16 - Desafio 1 - Ordenação de Dados em Go

## Informações do desafio
Objetivo:
Criar um programa em Go que lê informações de um arquivo. Realizar ordenação dos dados lidos por diferentes critérios (por exemplo, ordenação por nome e por idade). A ordenação por nome deve ser feita utilizando a tabela ASCII. Salvar o resultado ordenado em outro arquivo.


### Requisitos Técnicos:
Use a biblioteca padrão do Go para manipulação de arquivos. Considere trabalhar com dados estruturados, como linhas de um arquivo CSV.


### Estrutura do Programa:
Crie um arquivo de entrada contendo dados (por exemplo, CSV, cada linha representando uma entidade com algumas informações).
Implemente funções que leem o arquivo, realizam a ordenação dos dados por nome e por idade e retornam os resultados.
Salve os resultados ordenados em dois novos arquivos (um para cada critério de ordenação).


### Exemplo Arquivo de Entrada (arquivo-origem.csv)
```
Nome,Idade,Pontuação
Carlos,30,80
Carlos,22,75
Maria,30,95
Joao,25,80
carlos,15,90
```

### Exemplo de Arquivo de Saída(arquivo-destino.csv)
```
Nome,Idade,Pontuação
Carlos,22,75
Carlos,30,80
carlos,15,90
Joao,25,80
Maria,30,95
```

### Observações:
- Utilize a lógica de manipulação de arquivos em Go.
- Implemente a ordenação por nome e por idade.
- Se desejar, crie funções separadas para leitura, processamento e escrita de arquivos.
- Lide com erros de maneira apropriada.


Precisamos rodar o comando go run main.go ./arquivo-origem.csv ./arquivo-destino.csv. Será usado um arquivo CSV aleatório nosso para testar sua aplicação.

<div align="center">

[![GitHub](https://img.shields.io/badge/GitHub-Igor_Lage-blue?style=social&logo=github)](https://github.com/igor-rl) 

![Static Badge](https://img.shields.io/badge/07--12--2023-black)


</div>