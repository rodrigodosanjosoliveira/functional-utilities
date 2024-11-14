# Functional Utilities in Go

Este repositório contém uma implementação de funções utilitárias genéricas inspiradas em operações funcionais comuns, como `Map`, `Filter` e `Reduce`. Essas funções são escritas em Go e fazem uso de **Generics**, introduzidos na linguagem a partir da versão 1.18.

## Funcionalidades

### 1. `Map`
Aplica uma função a cada elemento de uma fatia, retornando uma nova fatia com os resultados.

```go
result := Map(s, strings.ToUpper)
// Exemplo: []string{"a", "b", "c"} -> []string{"A", "B", "C"}
```

### 2. `Filter`
Retorna uma nova fatia contendo apenas os elementos que satisfazem uma condição.

```go
evenNumbers := Filter(n, isEven[int])
// Exemplo: []int{1, 2, 3, 4} -> []int{2, 4}
```

### 3. `Reduce`
Aplica uma função a cada elemento de uma fatia, retornando um único valor.

```go
sum := Reduce(n, 0, add[int])
// Exemplo: []int{1, 2, 3, 4} -> 10
```

## Exemplo de Uso
```go
package main

import (
"fmt"
"strings"
)

func main() {
s := []string{"a", "b", "c"}
n := []int{1, 2, 3, 4}

    fmt.Println(Map(s, strings.ToUpper))          // ["A", "B", "C"]
    fmt.Println(Filter(n, isEven[int]))           // [2, 4]
    fmt.Println(Filter(n, isOdd[int]))            // [1, 3]
    fmt.Println(Reduce(n, 0, func(cur, next int) int { return cur + next })) // 10
    fmt.Println(Reduce(n, 1, func(cur, next int) int { return cur * next })) // 24
    fmt.Println(Reduce(s, "", func(cur, next string) string { return cur + next })) // "abc"
}
```

## Estrutura de Código

- **Map:** Aplica uma transformação a cada elemento.
- **Filter:** Filtra os elementos com base em uma condição.
- **Reduce:** Reduz a fatia a um valor único com base em uma função acumuladora.
- Funções auxiliares:
    - `isEven` e `isOdd` para verificar se números são pares ou ímpares.

## Pré-requisitos

- Go 1.18 ou superior(_devido ao Generics_)

## Como Executar

1. Clone o repositório:
```bash
git clone https://github.com/seu-usuario/functional-utilities-go.git
cd functional-utilities-go
```

2. Compile e execute o programa:
```bash
go run main.go
```

## Testes

Este repositório não contém testes unitários neste momento. Adicione-os utilizando a biblioteca `testing` do Go para validar as funções.


## Contribuindo

Contribuições são bem-vindas! Sinta-se à vontade para abrir issues ou enviar um **pull request**.