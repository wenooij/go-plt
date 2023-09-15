//go:build candle

package main

import (
	"math/rand"
	"time"

	"github.com/wenooij/go-plt"
)

func main() {
	var v float64

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := 0; i < 30*1000; i++ {
		plt.Candle(v)
		v += r.NormFloat64()
	}

	plt.Flush()
}
