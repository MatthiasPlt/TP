package main

import "fmt"

func Ft_coin(coins []int, amount int) int {

	ndp := make([]int, amount+1)

	for i := range ndp {
		ndp[i] = amount + 1
	}

	ndp[0] = 0

	// calcule du nombre  minimum de pieces pour chaque valeurs
	for _, coin := range coins {
		for i := coin; i <= amount; i++ {
			if ndp[i-coin]+1 < ndp[i] {
				ndp[i] = ndp[i-coin] + 1
			}
		}
	}

	// renvoie -1 si impossible
	if ndp[amount] > amount {
		return -1
	}
	return ndp[amount]
}

func main() {
	fmt.Println(Ft_coin([]int{1, 2, 5}, 11)) // Résultat : 3 (5 + 5 + 1)
	fmt.Println(Ft_coin([]int{2}, 3))        // Résultat : -1 (impossible)
	fmt.Println(Ft_coin([]int{1}, 0))        // Résultat : 0 (pas de pièces nécessaires)
}
