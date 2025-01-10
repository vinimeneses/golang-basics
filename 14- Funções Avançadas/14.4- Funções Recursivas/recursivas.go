package main

import (
	"fmt"
	"math/big"
)

// FIBO COM RECURSIVIDADE
func fibonacci(posicao uint64) uint64 {
	if posicao <= 1 {
		return posicao
	}
	return fibonacci(posicao-2) + fibonacci(posicao-1)
	
	
}

// FIBO O(N)

func fibonacci2(n int) big.Int {
    if n <= 1 {
        return *big.NewInt(int64(n))
    }

    prev := big.NewInt(0)
    curr := big.NewInt(1)
    
    for i := 2; i <= n; i++ {
        next := new(big.Int).Add(prev, curr)
        prev = curr
        curr = next
    }
	return *curr
}
	

func main() {
	fmt.Println("----------------FUNÇÃO LENTA - COMPLEXIDADE O(2^N)----------------")

	posicao := uint(47)

	for i := uint(0); i < posicao; i++ {
		fmt.Printf("%d, ", fibonacci(uint64(i)))
	}

	fmt.Println()
	fmt.Println("------------------------------------------------------------------------------------")
	fmt.Println("----------------RÁPIDO - COMPLEXIDADE O(N)----------------")
	
	posicao2 := 120

	for i := 0; i < posicao2; i++ {
		result := fibonacci2(i)
		fmt.Printf("%s, ", result.String())
	}
}
