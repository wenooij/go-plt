//go:build approx

package main

import (
	"math/rand"
	"time"

	"github.com/wenooij/go-plt"
)

func main() {
	const N = 1_000_00

	var a plt.Approx
	a.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

	for j := 0; j < 100; j++ {
		a.Reset()

		cnt := 0
		a.LogCall(N, func() { cnt++ })
		plt.Hist(float64(cnt))
	}
	plt.Flush()
}
