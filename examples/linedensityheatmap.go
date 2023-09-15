//go:build linedensityheatmap

package main

import (
	"bytes"
	"image"
	"image/png"
	"io/ioutil"
	"math/rand"
	"sort"

	"github.com/wenooij/go-plt"
)

func main() {
	const n = 10000

	records := make([]plt.LineRecord, 0, n)

	begin := plt.RelTime(0)
	end := plt.RelTime(100)

	var min float64
	var max float64

	for i := 0; i < n; i++ {
		v := rand.NormFloat64()
		if v < min {
			min = v
		}
		if max < v {
			max = v
		}
		records = append(records, plt.LineRecord{
			RelTime: plt.RelTime(float64(i)/float64(n)*float64(end-begin)) + begin,
			Value:   v,
		})
	}

	// _, lo, _, hi, _ := plt.Quartiles(records)
	lo, hi := min, max

	// preconditions.
	sort.Sort(plt.ByLineRecord(records))

	dst := image.NewRGBA(image.Rect(0, 0, 1000, 100))
	plt.LineDensityHeatMap{}.Render(dst, records, 20, lo, hi, begin, end)

	buf := bytes.NewBuffer(nil)
	png.Encode(buf, dst)
	ioutil.WriteFile("a.png", buf.Bytes(), 0755)
}
