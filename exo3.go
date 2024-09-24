package main

func Ft_non_overlap(intervals [][]int) int {
	if len(intervals) == 0 {
		return 0
	}
	intervals_a_supprimé := 0

	end := intervals[0][1]

	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] < end {
			intervals_a_supprimé++
		} else {
			end = intervals[i][1]
		}
	}

	return intervals_a_supprimé
}

/*
func main() {
	fmt.Println(Ft_non_overlap([][]int{{1, 2}, {2, 3}, {3, 4}, {1, 3}})) // Résultat : 1
	fmt.Println(Ft_non_overlap([][]int{{1, 2}, {2, 3}}))                 // Résultat : 0
	fmt.Println(Ft_non_overlap([][]int{{1, 2}, {1, 2}, {1, 2}}))         // Résultat : 2
}
*/
