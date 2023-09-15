package bucketize

func MedianSortedOctiles(vs []float64) (min, q0, q1, q2, q3, q4, q5, q6, max float64) {
	n := len(vs)
	min, q1, q3, q5, max = MedianSortedQuartiles(vs)
	q0, q2, q4, q6 = vs[n/8], vs[3*n/8], vs[5*n/8], vs[7*n/8]
	return
}
