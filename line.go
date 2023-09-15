package plt

import (
	"sort"

	"golang.org/x/exp/maps"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
)

type linePlot struct {
	series map[seriesKey]plotter.XYs
}

func (l *linePlot) render(p *plot.Plot) error {
	keys := maps.Keys(l.series)
	sort.Strings(keys)
	var plots []any
	for _, k := range keys {
		if len(l.series[k]) == 0 {
			continue
		}
		plots = append(plots, seriesName(k), l.series[k])
	}
	if len(plots) > 0 {
		plotutil.AddLinePoints(p, plots...)
	}
	return nil
}

func (l *linePlot) Reset() {
	l.series = make(map[string]plotter.XYs)
}

func Line(ys ...float64) {
	s := runtimeSeriesKey(2)
	Linef(Series(s), ys...)
}

func Linef(f *Format, ys ...float64) {
	Plinef(&globalPlot, f, ys...)
}

func Pline(p *Plot, ys ...float64) {
	s := runtimeSeriesKey(2)
	Plinef(p, Series(s), ys...)
}

func Plinef(p *Plot, f *Format, ys ...float64) {
	p.once.Do(p.initPlot)
	data := p.linePlot.series[f.seriesKey]
	for _, y := range ys {
		data = append(data, plotter.XY{X: float64(len(data)), Y: y})
	}
	p.linePlot.series[f.seriesKey] = data
}
