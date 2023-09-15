package plt

import (
	"sort"

	"golang.org/x/exp/maps"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

type barPlot struct {
	series map[seriesKey]float64
}

type barFormat struct{}

func (b *barPlot) render(p *plot.Plot) error {
	n := len(b.series)
	if n == 0 {
		return nil
	}
	width := 3.0 / float64(n)
	if width > 1 {
		width = 1
	}

	keys := maps.Keys(b.series)
	sort.Strings(keys)

	names := make([]string, n)
	bars := make(plotter.Values, n)
	for i, k := range keys {
		names[i] = seriesName(k)
		bars[i] = b.series[k]
	}
	p.NominalX(names...)

	// TODO: Add labels.
	e, err := plotter.NewBarChart(bars, vg.Length(width)*vg.Inch)
	if err != nil {
		return err
	}
	e.Color = plotutil.Color(0)
	p.Add(e)

	return nil
}

func (p *barPlot) Reset() {
	p.series = make(map[seriesKey]float64)
}

func Bar(v float64) {
	s := runtimeSeriesKey(2)
	Barf(Series(s), v)
}

func Barf(f *Format, v float64) {
	Pbarf(&globalPlot, f, v)
}

func Pbarf(p *Plot, f *Format, v float64) {
	p.once.Do(p.initPlot)
	p.barPlot.series[f.seriesKey] += v
}
