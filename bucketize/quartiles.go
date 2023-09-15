package bucketize

func MedianSortedQuartiles(records []float64) (min, q0, mean, q1, max float64) {
	n := len(records)
	return records[0], records[n/4], records[n/4], records[3*n/4], records[len(records)-1]
}
