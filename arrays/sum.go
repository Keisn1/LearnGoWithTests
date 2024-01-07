package arrays

func Sum(arr []int) (res int) {
	for _, nbr := range arr {
		res += nbr
	}
	return
}

func SumAll(nbrSlices [][]int) (res []int) {
	for _, nbrSlice := range nbrSlices {
		res = append(res, Sum(nbrSlice))
	}
	return
}

func SumAllTails(nbrSlices [][]int) (res []int) {
	for _, nbrSlice := range nbrSlices {
		if len(nbrSlice) == 0 {
			res = append(res, 0)
			continue
		}
		res = append(res, Sum(nbrSlice[1:]))
	}
	return
}
