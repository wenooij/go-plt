package plt

import (
	"sync"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/vg"
)

type Plot struct {
	*plot.Plot
	once sync.Once

	*barPlot
	*boxPlot
	*candlePlot
	*linePlot
	*scatterPlot
	*histPlot
	*covPlot
}

func (p *Plot) initPlot() {
	p.Plot = plot.New()

	if p.barPlot == nil {
		p.barPlot = new(barPlot)
	}
	if p.boxPlot == nil {
		p.boxPlot = new(boxPlot)
	}
	if p.candlePlot == nil {
		p.candlePlot = new(candlePlot)
	}
	if p.linePlot == nil {
		p.linePlot = new(linePlot)
	}
	if p.scatterPlot == nil {
		p.scatterPlot = new(scatterPlot)
	}
	if p.histPlot == nil {
		p.histPlot = new(histPlot)
	}
	if p.covPlot == nil {
		p.covPlot = new(covPlot)
	}

	p.barPlot.Reset()
	p.boxPlot.Reset()
	p.candlePlot.Reset()
	p.linePlot.Reset()
	p.scatterPlot.Reset()
	p.histPlot.Reset()
}

func (p *Plot) Flush() {
	p.once.Do(p.initPlot)

	for _, render := range []func(p *plot.Plot) error{
		p.candlePlot.render,
		p.boxPlot.render,
		p.barPlot.render,
		p.linePlot.render,
		p.scatterPlot.render,
		p.histPlot.render,
	} {
		if err := render(p.Plot); err != nil {
			panic(err)
		}
	}

	if err := p.Save(4*vg.Inch, 4*vg.Inch, "points.png"); err != nil {
		panic(err)
	}

	// Cov uses a seperate plot.
	p.covPlot.render()
}

var globalPlot Plot

func Flush() {
	globalPlot.Flush()
}
