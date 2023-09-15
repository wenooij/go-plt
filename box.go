package plt

import (
	"fmt"
	"sort"

	"golang.org/x/exp/maps"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/font"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

type boxPlot struct {
	series map[seriesKey]plotter.Values
}

// See plotutil.AddBoxPlots.
func addBoxPlots(plt *plot.Plot, width vg.Length, vs ...interface{}) error {
	var ps []plot.Plotter
	var names []string
	name := ""
	color := 0
	for _, v := range vs {
		switch t := v.(type) {
		case string:
			name = t

		case plotter.Valuer:
			b, err := plotter.NewBoxPlot(width, float64(len(names)), t)
			if err != nil {
				return err
			}
			b.FillColor = plotutil.Color(color)
			color++
			ps = append(ps, b)
			names = append(names, name)
			name = ""

		default:
			panic(fmt.Sprintf("plotutil: AddBoxPlots handles strings and plotter.Valuers, got %T", t))
		}
	}
	plt.Add(ps...)
	plt.NominalX(names...)
	return nil
}

func (b *boxPlot) render(p *plot.Plot) error {
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

	var boxes []any
	for _, k := range keys {
		// TODO: add series name.
		boxes = append(boxes, seriesName(k), b.series[k])
	}

	return addBoxPlots(p, font.Length(width)*vg.Inch, boxes...)
}

func (p *boxPlot) Reset() {
	p.series = make(map[seriesKey]plotter.Values)
}

func Box(vs ...float64) {
	s := runtimeSeriesKey(2)
	Boxf(Series(s), vs...)
}

func Boxf(f *Format, vs ...float64) {
	Pboxf(&globalPlot, f, vs...)
}

func Pboxf(p *Plot, f *Format, vs ...float64) {
	p.once.Do(p.initPlot)
	p.boxPlot.series[f.seriesKey] = append(p.boxPlot.series[f.seriesKey], vs...)
}
