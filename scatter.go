package plt

import (
	"sort"

	"golang.org/x/exp/maps"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
)

type scatterPlot struct {
	series map[seriesKey]plotter.XYs
}

func (s *scatterPlot) render(p *plot.Plot) error {
	keys := maps.Keys(s.series)
	sort.Strings(keys)
	var scatters []any
	for _, k := range keys {
		data := s.series[k]
		if len(data) == 0 {
			continue
		}
		scatters = append(scatters, seriesName(k), data)
	}
	if len(scatters) > 0 {
		plotutil.AddScatters(p, scatters...)
	}
	return nil
}

func (p *scatterPlot) Reset() {
	p.series = make(map[seriesKey]plotter.XYs)
}

func Scatter(ps ...plotter.XY) {
	s := runtimeSeriesKey(2)
	Scatterf(Series(s), ps...)
}

func Scatterf(f *Format, ps ...plotter.XY) {
	Pscatterf(&globalPlot, f, ps...)
}

func Pscatter(p *Plot, ps ...plotter.XY) {
	s := runtimeSeriesKey(2)
	Pscatterf(p, Series(s), ps...)
}

func Pscatterf(p *Plot, f *Format, ps ...plotter.XY) {
	p.once.Do(p.initPlot)
	p.scatterPlot.series[f.seriesKey] = append(p.scatterPlot.series[f.seriesKey], ps...)
}
