package slices

func Sum(arr []int) int {
	add := func(acc, x int) int {
		return acc + x
	}
	return Reduce(arr, add, 0)
}

func SumAll(arrays ...[]int) []int {
	sum := func(acc, x []int) []int {
		return append(acc, Sum(x))
	}

	return Reduce(arrays, sum, []int{})
}

func SumAllTails(arrays ...[]int) []int {
	sumTails := func(acc []int, x []int) []int {
		if len(x) == 0 {
			return append(acc, 0)
		} else {
			tail := x[1:]
			return append(acc, Sum(tail))
		}
	}
	return Reduce(arrays, sumTails, []int{})
}

func Reduce[A, B any](collection []B, fn func(A, B) A, initialValue A) A {
	var result = initialValue

	for _, x := range collection {
		result = fn(result, x)
	}
	return result
}

func Find[A any](collection []A, fn func(A) bool) (A, bool) {
	var val A
	for _, item := range collection {
		if fn(item) {
			return item, true
		}
	}
	return val, false
}
func Map[A, B any](collection []A, fn func(A) B) []B {
	res := make([]B, len(collection))
	for i, item := range collection {
		res[i] = fn(item)
	}
	return res
}
