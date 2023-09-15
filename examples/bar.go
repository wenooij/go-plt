//go:build bar

package main

import (
	"math/rand"

	"github.com/wenooij/go-plt"
)

func main() {
	for i := 0; i < 10; i++ {
		plt.Barf(plt.Series("//named"), rand.Float64())
		plt.Bar(2 * rand.ExpFloat64())
		plt.Bar(4 * rand.ExpFloat64())
		plt.Bar(2 + rand.ExpFloat64())
		plt.Bar(4 + rand.ExpFloat64())
	}
	plt.Flush()
}
