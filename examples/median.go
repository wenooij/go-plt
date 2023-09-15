//go:build median

package main

import "github.com/wenooij/go-plt/bucketize"

func main() {
	bucketize.Median([]float64{0, 0, 1, 1, 2, 3, 4, 4, 5, 5, 6})
}
