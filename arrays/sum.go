package arrays

func Sum(arr []int) (res int) {
	for _, nbr := range arr {
		res += nbr
	}
	return
}
