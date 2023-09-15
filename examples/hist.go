//go:build hist

package main

import (
	"math/rand"

	"github.com/wenooij/go-plt"
)

func main() {
	const n = 1_000

	for i := 0; i < n; i++ {
		plt.Hist(rand.NormFloat64())
	}

	plt.Flush()
}
