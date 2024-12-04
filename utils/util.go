package utils

func CountOccurnace(l []int) map[int]int {
	out := map[int]int{}

	for _, i := range l {
		out[i] += 1
	}

	return out
}
