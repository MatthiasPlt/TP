package main

import (
	"fmt"
)

func Ft_min_window(s string, t string) string {
	if len(s) == 0 || len(t) == 0 {
		return ""
	}
	charCount := make(map[rune]int)
	for _, char := range t {
		charCount[char]++
	}

	left, right := 0, 0
	minLength := len(s) + 1
	minWindow := ""
	required := len(charCount)
	formed := 0
	windowCounts := make(map[rune]int)

	for right < len(s) {
		char := rune(s[right])
		windowCounts[char]++

		if _, exists := charCount[char]; exists && windowCounts[char] == charCount[char] {
			formed++
		}

		for left <= right && formed == required {
			char = rune(s[left])

			if right-left+1 < minLength {
				minLength = right - left + 1
				minWindow = s[left : right+1]
			}

			windowCounts[char]--
			if _, exists := charCount[char]; exists && windowCounts[char] < charCount[char] {
				formed--
			}
			left++
		}

		right++
	}

	if minLength == len(s)+1 {
		return ""
	}

	return minWindow
}

func main() {
	result1 := Ft_min_window("ADOBECODEBANC", "ABC") // Résultat attendu : "BANC"
	fmt.Println(result1)

	result2 := Ft_min_window("a", "aa") // Résultat attendu : ""
	fmt.Println(result2)
}
