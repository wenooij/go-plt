package bucketize

import (
	"fmt"
	"math"

	"gonum.org/v1/gonum/stat/distuv"
)

type Normal struct {
	buckets []Bucket
}

func MakeNormal(n int, mu, stddev float64) Normal {
	norm := distuv.Normal{Mu: mu, Sigma: stddev}
	b := Normal{buckets: make([]Bucket, n)}
	step := 1 / float64(n)
	lo := norm.Quantile(step)
	b.buckets[0] = Bucket{
		Lo: math.Inf(-1),
		Hi: lo,
	}
	var next float64
	for i := 1; i < n; i++ {
		next = norm.Quantile(float64(i+1) * step)
		b.buckets[i] = Bucket{
			Lo: lo,
			Hi: next,
		}
		lo = next
	}
	return b
}

func (a Normal) Bucketize(v float64) BucketEntry {
	for i, b := range a.buckets {
		if b.Has(v) {
			return BucketEntry{Index: i, Bucket: b}
		}
	}
	panic(fmt.Errorf("Normal.Bucketize: failed to map a value: %v", v))
}
