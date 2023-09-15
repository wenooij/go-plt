package plt

import (
	"math"
	"sort"

	"golang.org/x/exp/maps"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

type candlePlot struct {
	series map[seriesKey]plotter.Values
}

func (c *candlePlot) render(p *plot.Plot) error {
	const candleDefaultSize = 30.0
	const candleMaxBoxes = 60.0

	keys := maps.Keys(c.series)
	sort.Strings(keys)

	candles := make([][]plotter.Valuer, len(keys))
	for i, k := range keys {
		buf := c.series[k]

		boxSize := candleDefaultSize
		boxes := math.Ceil(float64(len(buf)) / boxSize)
		if boxes > candleMaxBoxes {
			boxSize = math.Ceil(float64(len(buf)) / candleMaxBoxes)
			boxes = candleMaxBoxes
		}

		var candle []plotter.Valuer
		for i := 0; i < int(boxes); i++ {
			v := buf[int(boxSize*float64(i)):]
			if len(v) > int(boxSize) {
				v = v[:int(boxSize)]
			}
			candle = append(candle, v)
		}
		candles[i] = append(candles[i], candle...)
	}

	for i := range keys {
		cs := candles[i]

		n := len(cs)
		if n == 0 {
			continue
		}
		width := 3.0 / float64(n)
		if width > 1 {
			width = 1
		}

		var qs []plot.Plotter
		for i := range cs {
			q, err := plotter.NewBoxPlot(vg.Length(width)*vg.Inch, float64(i), cs[i])
			if err != nil {
				return err
			}
			m := q.Median
			if i > 0 {
				m0 := qs[i-1].(*plotter.BoxPlot).Median
				if m < m0 {
					q.FillColor = plotutil.Color(0)
				} else if m0 < m {
					q.FillColor = plotutil.Color(1)
				}
			}
			qs = append(qs, q)
		}

		p.Add(qs...)
	}

	return nil
}

func (p *candlePlot) Reset() {
	p.series = make(map[string]plotter.Values)
}

func Candle(vs ...float64) {
	s := runtimeSeriesKey(2)
	Candlef(Series(s), vs...)
}

func Candlef(f *Format, vs ...float64) {
	Pcandlef(&globalPlot, f, vs...)
}

func Pcandlef(p *Plot, f *Format, vs ...float64) {
	p.once.Do(p.initPlot)
	p.candlePlot.series[f.seriesKey] = append(p.candlePlot.series[f.seriesKey], vs...)
}
