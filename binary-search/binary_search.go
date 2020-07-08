package binarysearch

func SearchInts(input []int, key int) int {
	x := 0
	y := len(input) - 1

	for x <= y {
		m := (y + x) / 2
		if input[m] > key {
			y = m - 1
		} else if input[m] < key {
			x = m + 1
		} else {
			return m
		}
	}

	return -1
}
