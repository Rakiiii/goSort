package gosort

func QuicksortInt(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	left, right := 0, len(arr)-1

	pivot := len(arr) >> 1
	//pivot := rand.Int() % len(a)

	arr[pivot], arr[right] = arr[right], arr[pivot]

	for i, _ := range arr {
		if arr[i] < arr[right] {
			arr[left], arr[i] = arr[i], arr[left]
			left++
		}
	}

	arr[left], arr[right] = arr[right], arr[left]

	QuicksortInt(arr[:left])
	QuicksortInt(arr[left+1:])

	return arr
}
