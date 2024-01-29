package arrays

func Sum(arr []int) int {
	var sum int
	for _, nbr := range arr {
		sum += nbr
	}
	return sum
}

func SumAll(nbrSlices [][]int) []int {
	var sums []int
	for _, nbrSlice := range nbrSlices {
		sums = append(sums, Sum(nbrSlice))
	}
	return sums
}

func SumAllTails(nbrSlices [][]int) []int {
	var sumsTail []int
	for _, nbrSlice := range nbrSlices {
		if len(nbrSlice) == 0 {
			sumsTail = append(sumsTail, 0)
			continue
		}
		sumsTail = append(sumsTail, Sum(nbrSlice[1:]))
	}
	return sumsTail
}
