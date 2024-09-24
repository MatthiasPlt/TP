package main

func Ft_max_substring(s string) int {
	charMap := make(map[rune]int)
	left := 0
	maxLength := 0

	for right, char := range s {
		if idx, found := charMap[char]; found && idx >= left {
			left = idx + 1
		}

		charMap[char] = right

		currentLength := right - left + 1
		if currentLength > maxLength {
			maxLength = currentLength
		}
	}

	return maxLength
}

/*
func main() {
	result1 := Ft_max_substring("abcabcbb") // Résultat attendu : 3
	fmt.Println(result1)                    // Affiche 3

	result2 := Ft_max_substring("bbbbb") // Résultat attendu : 1
	fmt.Println(result2)                 // Affiche 1

}
*/
