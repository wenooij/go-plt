//go:build stackedarea

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
	const n = 150000

	records := make([]plt.Record, 0, n)

	keys := []plt.Key{
		"a",
		"b",
		"c",
		"d",
		"e",
		"f",
		"g",
		"h",
		"i",
	}

	begin := plt.RelTime(0)
	end := plt.RelTime(2000)

	vs := make([]float64, len(keys))

	for i := 0; i < n; i++ {
		for j := 0; j < len(keys); j++ {
			records = append(records, plt.Record{
				RelTime: plt.RelTime(float64(i)/float64(n)*float64(end-begin)) + begin,
				Entry: plt.Entry{
					Key:   keys[j],
					Value: vs[j],
				},
			})
			vs[j] += rand.ExpFloat64() - rand.ExpFloat64()
			if vs[j] < 0 {
				vs[j] = 0
			}
		}
	}

	// preconditions.
	sort.Sort(plt.ByIntegratedKey{Keys: keys, Ints: plt.IntegrateKeys(keys, records)})
	sort.Sort(plt.ByRecord(records))

	dst := image.NewRGBA(image.Rect(0, 0, 1000, 100))
	plt.StackedArea{}.Render(dst, records, keys, begin, end)

	buf := bytes.NewBuffer(nil)
	png.Encode(buf, dst)
	ioutil.WriteFile("a.png", buf.Bytes(), 0755)
}
