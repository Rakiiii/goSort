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

func QuicksortFloat64(arr []float64) []float64 {
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

	QuicksortFloat64(arr[:left])
	QuicksortFloat64(arr[left+1:])

	return arr
}