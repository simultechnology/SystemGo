package mysort

func QuickSort(bottom int, top int, data []int) {

	var (
		lower int
		upper int
		div   int
		tmp   int
	)

	if bottom >= top {
		return
	}

	div = data[bottom]

	lower = bottom
	upper = top
	for lower < upper {
		for lower <= upper && data[lower] <= div {
			lower++
		}

		for lower <= upper && data[upper] > div {
			upper--
		}

		if lower < upper {
			tmp = data[lower]
			data[lower] = data[upper]
			data[upper] = tmp
		}
	}

	tmp = data[bottom]
	data[bottom] = data[upper]
	data[upper] = tmp

	QuickSort(bottom, upper-1, data)
	QuickSort(upper+1, top, data)

}
