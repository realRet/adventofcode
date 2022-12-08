package utils

func FindMaxInArray(v []int) int {
	max := v[0]
	for i := 1; i < len(v); i++ {
		if max < v[i] {
			max = v[i]
		}
	}

	return max
}

func GetVerticalArray(a [][]int, index int) []int {
	verticalArray := make([]int, len(a))
	for n := 0; n < len(a); n++ {
		verticalArray[n] = a[n][index]
	}
	return verticalArray
}
