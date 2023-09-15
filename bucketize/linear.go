package bucketize

import (
	"fmt"
	"math"
	"sort"
)

type Linear struct {
	buckets []Bucket
}

func MakeLinear(n int, lo, hi float64) Linear {
	if n < 2 {
		panic("MakeLinear: less than 2 buckets is not supported")
	}
	b := Linear{}
	b.buckets = append(b.buckets, Bucket{
		Lo: math.Inf(-1),
		Hi: lo,
	})
	r := hi - lo
	step := r / float64(n)
	var next float64
	for i := 0; i < n-2; i++ {
		next = lo + step
		b.buckets = append(b.buckets, Bucket{
			Lo: lo,
			Hi: next,
		})
		lo = next
	}
	b.buckets = append(b.buckets, Bucket{
		Lo: next,
		Hi: math.Inf(+1),
	})
	return b
}

func (a Linear) Bucketize(v float64) BucketEntry {
	// TODO: use binary search at large threshold.
	i, found := sort.Find(len(a.buckets), func(i int) int {
		return -a.buckets[i].Cmp(v)
	})
	if !found {
		panic(fmt.Errorf("Linear.Bucketize: failed to map a value: %v", v))
	}
	return BucketEntry{Index: i, Bucket: a.buckets[i]}
}
