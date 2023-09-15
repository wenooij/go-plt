package plt

import (
	"bytes"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"io/ioutil"
	"math"
	"sort"

	"github.com/wenooij/go-plt/bucketize"
)

type LineRecord struct {
	RelTime RelTime
	Value   float64
}

type LineDensityHeatMap struct{}

// precondition: records is sorted
// precondition: keys is sorted
// precondition: start < end
func (a LineDensityHeatMap) Render(dst draw.Image, records []LineRecord, yBuckets int, lo, hi float64, start, end RelTime) {
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
	dstDY := float64(max.Y - min.Y)
	yStep := dstDY / float64(yBuckets)
	getBucketY := func(i int) int {
		return max.Y - int(math.Ceil(float64(i+1)*yStep))
	}

	// bucketer
	bucketizer := bucketize.MakeLinear(yBuckets, lo, hi)
	// bucketizer := bucketize.MakeNormal(yBuckets, 0, 2)
	getBucket := func(y float64) int {
		return bucketizer.Bucketize(y).Index
	}
	bucketize := func(buckets []float64, vs []float64) {
		for _, v := range vs {
			b := getBucket(v)
			buckets[b]++
		}
	}

	// colorizer
	makeColors := func(hiBucket float64) HeatColorHasher {
		return MakeHeatColorHasher(0, hiBucket)
	}

	// Primitive draw method.
	drawRect := func(c color.Color, x0, y0, x1, y1 int) {
		draw.Draw(dst, image.Rect(x0, y0, x1, y1), image.NewUniform(c), image.Point{}, draw.Src)
	}

	// Draw a frame of values
	draw := func(buckets []float64, t RelTime) {
		var hiBucket float64
		for _, v := range buckets {
			if hiBucket < v {
				hiBucket = v
			}
		}
		colors := makeColors(hiBucket)

		for i, v := range buckets {
			x0, y0 := getX(t), getBucketY(i)
			x1, y1 := x0+int(step), y0+int(yStep)
			drawRect(colors.Hash(v), x0, y0, x1, y1)
		}
	}

	// Render the plot.
	var lastRel RelTime
	zeroes := make([]float64, yBuckets)
	vs := make([]float64, 0, 64)
	buckets := make([]float64, yBuckets)
	for i, r := range records {
		if i > 0 && lastRel != r.RelTime {
			bucketize(buckets, vs)
			draw(buckets, lastRel)
			copy(buckets, zeroes)
			vs = vs[:0]
		}
		vs = append(vs, r.Value)
		lastRel = r.RelTime
	}
}

type lineDensityHeatMapPlotter struct {
	records []LineRecord
}

var globalLineDensityPlotter = lineDensityHeatMapPlotter{}

func LineDensity(t RelTime, vs ...float64) {
	s := runtimeSeriesKey(2)
	LineDensityf(Series(s), t, vs...)
}

func LineDensityf(_ *Format, t RelTime, vs ...float64) {
	for _, v := range vs {
		globalLineDensityPlotter.records = append(globalLineDensityPlotter.records, LineRecord{
			RelTime: t,
			Value:   v,
		})
	}
}

func FlushLineDensity() {
	sort.Sort(ByLineRecord(globalLineDensityPlotter.records))
	var begin, end RelTime
	var lo, hi float64
	for _, r := range globalLineDensityPlotter.records {
		if r.Value < lo {
			lo = r.Value
		}
		if hi < r.Value {
			hi = r.Value
		}
		if r.RelTime < begin {
			begin = r.RelTime
		}
		if end < r.RelTime {
			end = r.RelTime
		}
	}

	dst := image.NewRGBA(image.Rect(0, 0, 1000, 100))
	LineDensityHeatMap{}.Render(dst, globalLineDensityPlotter.records, 10, lo, hi, begin, end)

	buf := bytes.NewBuffer(nil)
	png.Encode(buf, dst)
	ioutil.WriteFile("a.png", buf.Bytes(), 0755)

	ResetLineDensity()
}

func ResetLineDensity() {
	globalLineDensityPlotter = lineDensityHeatMapPlotter{}
}
