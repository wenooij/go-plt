package bucketize

func Median(vs []float64) float64 {
	switch len(vs) {
	case 0:
		return 0
	case 1:
		return vs[0]
	default:
		idx := selectMedian(vs, 0, len(vs)-1, len(vs))
		return vs[idx]
	}
}

func selectMedian(vs []float64, left, right, n int) int {
	for left < right {
		pivotIndex := pivot(vs, left, right)
		pivotIndex = partition(vs, left, right, pivotIndex, n)
		if n == pivotIndex {
			return n
		} else if n < pivotIndex {
			right = pivotIndex - 1
		} else {
			left = pivotIndex + 1
		}
	}
	return left
}

func pivot(vs []float64, left, right int) int {
	// for 5 or less elements just get median
	if right-left < 5 {
		return partition5(vs[left : right+1])
	}
	// otherwise move the medians of five-element subgroups to the first n/5 positions
	for i := left; i <= right; i += 5 {
		// get the median position of the i'th five-element subgroup
		subRight := i + 4
		if subRight > right {
			subRight = right
		}
		median5 := partition5(vs[i : subRight+1])

		t := left + (i-left)/5
		vs[median5], vs[t] = vs[t], vs[median5]
	}

	// compute the median of the n/5 medians-of-five
	mid := (right-left)/10 + left + 1
	return selectMedian(vs, left, left+(right-left)/5, mid)
}

func partition(vs []float64, left, right, pivotIndex, n int) int {
	pivotValue := vs[pivotIndex]
	vs[pivotIndex], vs[right] = vs[right], vs[pivotIndex] // Move pivot to end
	storeIndex := left
	// Move all elements smaller than the pivot to the left of the pivot
	for i := left; i < right; i++ {
		if vs[i] < pivotValue {
			vs[storeIndex], vs[i] = vs[i], vs[storeIndex]
			storeIndex++
		}
	}
	// Move all elements equal to the pivot right after
	// the smaller elements
	storeIndexEq := storeIndex
	for i := storeIndex; i < right; i++ {
		if vs[i] == pivotValue {
			vs[storeIndexEq], vs[i] = vs[i], vs[storeIndexEq]
			storeIndexEq++
		}
	}
	vs[right], vs[storeIndexEq] = vs[storeIndexEq], vs[right] // Move pivot to its final place
	// Return location of pivot considering the desired location n
	if n < storeIndex {
		return storeIndex // n is in the group of smaller elements
	}
	if n <= storeIndexEq {
		return n // n is in the group equal to pivot
	}
	return storeIndexEq // n is in the group of larger elements
}

// precondition: len(vs) <= 5
func partition5(vs []float64) int {
	for i := 1; i < len(vs); {
		j := i
		for vs[j] < vs[j-1] {
			vs[j-1], vs[j] = vs[j], vs[j-1]
			j--
		}
		i += 2
	}
	return len(vs) / 2
}
