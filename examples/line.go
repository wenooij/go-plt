//go:build line

package main

import (
	"math/rand"

	"github.com/wenooij/go-plt"
)

func main() {
	var v float64

	for i := 0; i < 300; i++ {
		plt.Line(v)
		v += rand.Float64() * rand.NormFloat64()
	}

	for i := 0; i < 300; i++ {
		plt.Line(v)
		v += rand.NormFloat64() * rand.NormFloat64()
	}

	for i := 0; i < 300; i++ {
		plt.Line(v)
		v += rand.ExpFloat64() * rand.NormFloat64()
	}

	plt.Flush()
}
