package plt

import (
	"image"
	"image/color"
	"image/draw"
	"math/rand"
	"time"
)

type Key string

type Entry struct {
	Key   Key
	Value float64
}

type RelTime time.Duration

type Record struct {
	RelTime
	Entry
}

type StackedArea struct {
}

// precondition: records is sorted
// precondition: keys is sorted
// precondition: start < end
func (a StackedArea) Render(dst draw.Image, records []Record, keys []Key, start, end RelTime) {
	bounds := dst.Bounds()
	min, max := bounds.Min, bounds.Max

	// X geom math
	dx := float64(end - start)
	dstDX := float64(max.X - min.X)
	step := dstDX / dx
	if step < 1 {
		step = 1
	}
	getX := func(t RelTime) int {
		return int(float64(t-start)/dx*dstDX) + min.X
	}

	// Y geom math
	dy := 1.0
	dstDY := float64(max.Y - min.Y)
	getY := func(cv float64) int {
		return int(float64(cv-0)/dy*dstDY) + min.Y
	}

	// Value norming.
	cumnorm := func(vs []float64) {
		sum := 0.0
		for i, v := range vs {
			sum += v
			vs[i] = sum
		}
		if sum != 0 {
			for i := range vs {
				vs[i] /= sum
			}
		}
	}

	// Key-index mapping
	keyIndex := make(map[Key]int, len(keys))
	for i, k := range keys {
		keyIndex[k] = i
	}

	// Key-color mapping.
	// TODO: use deterministic color-hasher.
	colors := make([]color.Color, len(keys))
	for i := range keys {
		colors[i] = color.RGBA{
			R: uint8(rand.Intn(256)),
			G: uint8(rand.Intn(256)),
			B: uint8(rand.Intn(256)),
			A: 255,
		}
	}

	// Primitive draw method.
	drawRect := func(c color.Color, x0, y0, x1, y1 int) {
		draw.Draw(dst, image.Rect(x0, y0, x1, y1), image.NewUniform(c), image.Point{}, draw.Src)
	}

	// Draw a frame of values
	draw := func(vs []float64, t RelTime) {
		x0 := getX(t)
		x1 := x0 + int(step)
		y0 := 0
		var y1 int
		for i, v := range vs {
			y1 = getY(v)
			drawRect(colors[i], x0, y0, x1, y1)
			y0 = y1
		}
		// x0 := getX(t)
		// x1 := x0 + int(step)
		// y0 := 0
		// y1 := getY(vs[0])
		// println(vs[0], y1)
		// fmt.Println(vs)
		// drawRect(colors[0], x0, y0, x1, y1)
	}

	// Render the plot.
	var lastRel RelTime
	zeroes := make([]float64, len(keys))
	vs := make([]float64, len(keys))
	for i, r := range records {
		if i > 0 && lastRel != r.RelTime {
			cumnorm(vs)
			draw(vs, lastRel)
			copy(vs, zeroes)
		}
		vs[keyIndex[r.Key]] += r.Value
		lastRel = r.RelTime
	}
}
