package main

import "fmt"

// Ft_profit retourne le plus grand bénéfice possible en achetant et en vendant un bien.
func Ft_profit(prices []int) int {
	// Initialiser le bénéfice maximum à 0
	maxProfit := 0

	// Initialiser le prix d'achat minimum avec le premier prix
	minPrice := prices[0]

	// Parcourir le tableau des prix
	for _, price := range prices {
		// Mettre à jour le prix d'achat minimum si le prix actuel est inférieur
		if price < minPrice {
			minPrice = price
		}

		// Calculer le bénéfice potentiel en vendant à ce prix
		profit := price - minPrice

		// Mettre à jour le bénéfice maximum si le bénéfice actuel est plus élevé
		if profit > maxProfit {
			maxProfit = profit
		}
	}

	return maxProfit
}

func main() {
	// Tests
	result1 := Ft_profit([]int{7, 1, 5, 3, 6, 4}) // Résultat attendu : 5
	fmt.Println(result1)                          // Affiche 5

	result2 := Ft_profit([]int{1, 2, 3, 4, 5}) // Résultat attendu : 4 (acheter à 1, vendre à 5)
	fmt.Println(result2)                       // Affiche 4

	result3 := Ft_profit([]int{7, 6, 4, 3, 1}) // Résultat attendu : 0 (pas de bénéfice possible)
	fmt.Println(result3)                       // Affiche 0
}
