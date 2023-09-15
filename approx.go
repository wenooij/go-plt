package plt

import "math"

type Rand interface {
	Uint64() uint64
}

// Approx implements approximate counting.
//
// It supports counts up to 2^{MaxUint64} while minimizing calls to the RNG.
type Approx struct {
	exp uint64
	Rand
}

// Reset the Approx instance.
func (a *Approx) Reset() {
	a.exp = 0
}

// EstimateCount returns an estimate of the count for the Approx instance.
func (a *Approx) EstimateCount() float64 {
	return math.Pow(2, float64(a.exp))
}

// LogCall is a logarithmic closure runner.
//
// It calls fn O(lg2[N+n]/N) times, where N is the total count since Reset.
func (a *Approx) LogCall(n int, fn func()) {
	exp := a.exp
	mask := uint64(1)<<(exp%64) - 1

	for ; n > 0; n-- {
		e := exp

		// Large-EXP decomposition.
		for 64 <= e {
			if a.Uint64() != 0 {
				continue
			}
			e -= 64
		}
		rv := a.Uint64()

		// EXP flips.
		if rv&mask != 0 {
			continue
		}

		// Increment and call fn.
		exp++
		mask = uint64(1)<<e - 1
		fn()
	}

	a.exp = exp
}
