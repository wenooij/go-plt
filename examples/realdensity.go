//go:build realdensity

package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"time"

	"github.com/wenooij/go-plt"
)

func main() {
	const n = 10000

	for i := 0; i < n; i++ {
		start := time.Now()
		for i := 0; i < 100; i++ {
			exec()
		}
		d := time.Since(start)
		plt.LineDensity(plt.RelTime(i/10), float64(d))
	}

	plt.FlushLineDensity()
}

var e1 error

func exec() {
	i := rand.Intn(16)

	var b [1024]byte
	rand.Read(b[:])

	e1 = ioutil.WriteFile(fmt.Sprintf("./examples/realdensity/%d", i), b[:], 0755)
}
