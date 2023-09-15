//go:build cov

package main

import (
	"math/rand"

	"github.com/wenooij/go-plt"
)

func main() {
	plt.CovInit(4)

	for i := 0; i < 1000; i++ {
		n0, n1, n2 := rand.NormFloat64()+4, rand.NormFloat64()+6, rand.NormFloat64()+12
		n3 := n0 + n1 - n2

		plt.Cov(n0)
		plt.Cov(n1)
		plt.Cov(n2)
		plt.Cov(n3)

		plt.CovFlush()
	}

	plt.Flush()
}
