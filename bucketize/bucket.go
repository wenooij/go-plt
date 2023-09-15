package bucketize

type Bucket struct {
	Lo float64
	Hi float64
}

func (b Bucket) Cmp(v float64) int {
	switch {
	case v < b.Lo:
		return 1
	case b.Hi <= v:
		return -1
	default:
		return 0
	}
}

func (b Bucket) Has(v float64) bool {
	return b.Lo <= v && v < b.Hi
}

type BucketEntry struct {
	Index int
	Bucket
}
