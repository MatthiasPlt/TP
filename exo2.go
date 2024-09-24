package main

import "fmt"

func Ft_missing(nums []int) int {
	n := len(nums)
	total := n * (n + 1) / 2
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return total - sum
}

func main() {
	fmt.Println(Ft_missing([]int{3, 1, 2}))                   // Résultat : 0
	fmt.Println(Ft_missing([]int{0, 1}))                      // Résultat : 2
	fmt.Println(Ft_missing([]int{9, 6, 4, 2, 3, 5, 7, 0, 1})) // Résultat : 8
}
