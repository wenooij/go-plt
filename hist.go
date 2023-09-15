package plt

import (
	"image/color"
	"math"
	"sort"

	"golang.org/x/exp/maps"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
)

type histPlot struct {
	series map[seriesKey]plotter.Values
}

func makeHist(vs plotter.Valuer, colorIndex int) (*plotter.Histogram, error) {
	// Sturges' Rule.
	n := 1 + 3.322*math.Log(float64(vs.Len()))
	h, err := plotter.NewHist(vs, int(math.Round(n)))
	if err != nil {
		return nil, err
	}
	c := plotutil.Color(colorIndex).(color.RGBA)
	h.FillColor = c
	h.Normalize(1)
	return h, nil
}

func (h *histPlot) render(p *plot.Plot) error {
	keys := maps.Keys(h.series)
	sort.Strings(keys)
	// TODO: add legends.
	var hists []plot.Plotter
	for i, k := range keys {
		hist := h.series[k]
		h, err := makeHist(hist, i)
		if err != nil {
			return err
		}
		hists = append(hists, h)
	}

	p.Add(hists...)
	return nil
}

func (h *histPlot) Reset() {
	h.series = make(map[string]plotter.Values)
}

func Hist(vs ...float64) {
	s := runtimeSeriesKey(2)
	Histf(Series(s), vs...)
}

func Histf(f *Format, vs ...float64) {
	Phistf(&globalPlot, f, vs...)
}

func Phist(p *Plot, f *Format, vs ...float64) {
	s := runtimeSeriesKey(2)
	Phistf(p, Series(s), vs...)
}

func Phistf(p *Plot, f *Format, vs ...float64) {
	p.once.Do(p.initPlot)
	p.histPlot.series[f.seriesKey] = append(p.histPlot.series[f.seriesKey], vs...)
}
