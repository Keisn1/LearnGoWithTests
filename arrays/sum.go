package arrays

func Sum(arr []int) int {
	plus := func(a, b int) int { return a + b }
	return Reduce(arr, plus, 0)
}

func Reduce[A, B any](arr []A, accumulator func(B, A) B, initVal B) B {
	var res = initVal
	for _, x := range arr {
		res = accumulator(res, x)
	}
	return res
}

func SumAll(nbrSlices ...[]int) []int {
	sums := func(a, b []int) []int {
		a = append(a, Sum(b))
		return a
	}
	return Reduce(nbrSlices, sums, []int{})
}

func SumAllTails(nbrSlices ...[]int) []int {
	sumsTails := func(acc, b []int) []int {
		if len(b) == 0 {
			acc = append(acc, 0)
		} else {
			return append(acc, Sum(b[1:]))
		}
		return acc
	}

	return Reduce(nbrSlices, sumsTails, []int{})
}
