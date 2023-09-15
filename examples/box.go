//go:build box

package main

import (
	"math/rand"

	"github.com/wenooij/go-plt"
)

func main() {
	const n = 1000

	for i := 0; i < n; i++ {
		plt.Boxf(plt.Series("uniform"), rand.Float64())
		plt.Boxf(plt.Series("norm"), rand.NormFloat64())
		plt.Boxf(plt.Series("exp"), rand.ExpFloat64())
	}

	plt.Flush()
}
