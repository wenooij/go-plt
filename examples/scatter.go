//go:build scatter

package main

import (
	"math/rand"

	"github.com/wenooij/go-plt"
	"gonum.org/v1/plot/plotter"
)

func main() {
	const n = 900

	for i := 0; i < n; i++ {
		x, y := rand.NormFloat64(), rand.NormFloat64()
		plt.Scatterf(plt.Series("norm"), plotter.XY{X: x, Y: y})
	}

	for i := 0; i < n; i++ {
		x, y := rand.ExpFloat64(), rand.ExpFloat64()
		plt.Scatterf(plt.Series("exp"), plotter.XY{X: x, Y: y})
	}

	for i := 0; i < n; i++ {
		x, y := rand.Float64(), rand.Float64()
		plt.Scatterf(plt.Series("uniform"), plotter.XY{X: x, Y: y})
	}

	plt.Flush()
}
