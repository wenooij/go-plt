//go:build octiles

package main

import (
	"math/rand"
	"sort"

	"github.com/wenooij/go-plt"
	"github.com/wenooij/go-plt/bucketize"
)

func main() {
	const n = 10000

	vs := make([]float64, 0, n)
	for i := 0; i < n; i++ {
		vs = append(vs, rand.NormFloat64())
	}

	sort.Float64s(vs)

	// min, q0, q1, q2, q3, q4, q5, q6, max := bucketize.MeanOctiles(vs)
	// plt.Boxf(plt.Series("MeanOctiles"), min, q0, q1, q2, q3, q4, q5, q6, max)

	// min, q0, q1, q2, q3, q4, q5, q6, max = bucketize.MedianSortedOctiles(vs)
	// plt.Boxf(plt.Series("MedianSortedOctiles"), min, q0, q1, q2, q3, q4, q5, q6, max)

	min, q0, q1, q2, q3, q4, q5, q6, max := bucketize.MedianSortedOctiles(vs)
	plt.Barf(plt.Series("m-"), min)
	plt.Barf(plt.Series("m0"), q0)
	plt.Barf(plt.Series("m1"), q1)
	plt.Barf(plt.Series("m2"), q2)
	plt.Barf(plt.Series("m3"), q3)
	plt.Barf(plt.Series("m4"), q4)
	plt.Barf(plt.Series("m5"), q5)
	plt.Barf(plt.Series("m6"), q6)
	plt.Barf(plt.Series("m+"), max)
	plt.Flush()
}
