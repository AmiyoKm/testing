package slices

func Sum(arr []int) int {
	res := 0

	for _, num := range arr {
		res += num
	}
	return res
}

func SumAll(arrays ...[]int) []int {
	sums := []int{}

	for _, arr := range arrays {
		sums = append(sums, Sum(arr))
	}

	return sums
}

func SumAllTails(arrays ...[]int) []int {
	sums := []int{}
	for _, arr := range arrays {
		if len(arr) == 0 {
			sums = append(sums, 0)
		} else {
			sums = append(sums, Sum(arr[1:]))
		}
	}
	return sums
}
