package plt

import "image/color"

type HeatColorHasher struct {
	qstep, q0, q1, q2, q3 float64
	c0, c1, c2, c3        color.RGBA
}

func MakeHeatColorHasher(lo, hi float64) HeatColorHasher {
	r := hi - lo
	qstep := r / 4
	return HeatColorHasher{
		qstep: qstep,
		q0:    lo,
		q1:    lo + qstep,
		q2:    hi - qstep,
		q3:    hi,
		c0:    color.RGBA{R: 255, G: 255, B: 255, A: 255}, // white
		c1:    color.RGBA{R: 255, G: 255, B: 66, A: 255},  // yellow
		c2:    color.RGBA{R: 255, G: 66, B: 66, A: 255},   // red
		c3:    color.RGBA{A: 255},                         // black
	}
}

func (a HeatColorHasher) Hash(v float64) color.RGBA {
	colorBlend := func(t float64, c0, c1 color.RGBA) color.RGBA {
		return color.RGBA{
			R: uint8((1-t)*float64(c0.R) + t*float64(c1.R)),
			G: uint8((1-t)*float64(c0.G) + t*float64(c1.G)),
			B: uint8((1-t)*float64(c0.B) + t*float64(c1.B)),
			A: 255,
		}
	}

	switch {
	case v < a.q0:
		return a.c0
	case v < a.q1:
		return colorBlend((v-a.q0)/a.qstep, a.c0, a.c1)
	case v < a.q2:
		return colorBlend((v-a.q1)/a.qstep, a.c1, a.c2)
	case v < a.q3:
		return colorBlend((v-a.q2)/a.qstep, a.c2, a.c3)
	default: // q4 <= v
		return a.c3
	}
}
