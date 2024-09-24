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
	fmt.Println(Ft_profit([]int{7, 1, 5, 3, 6, 4})) // resultat : 5
	// si on achète au jour 1, nous payons 1,
	// et si nous le vendons au 4eme jour, nous gagnons 6, le bénéfice est 6-1
	fmt.Println(Ft_profit([]int{7, 6, 4, 3, 1})) // resultat : 0
}
