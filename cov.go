package plt

import (
	"image/color"
	"os"
	"sort"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/font"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
	"gonum.org/v1/plot/vg/draw"
	"gonum.org/v1/plot/vg/vgimg"
	"gonum.org/v1/plot/vg/vgsvg"
)

type covPlot struct {
	n int

	scatters [][]plotter.XYs  // i,j.
	hists    []plotter.Values // i,j; i=j.

	i   int
	buf []float64

	series []seriesKey
}

func makeGlyphStyleMix(h1, h2 *plotter.Histogram, p plotter.XY, s1, s2 draw.GlyphStyle) draw.GlyphStyle {
	width1, width2 := h1.Width, h2.Width

	i := sort.Search(len(h1.Bins), func(i int) bool {
		b := h1.Bins[i]
		return p.X < b.Min
	})
	if i == len(h1.Bins) {
		i--
	}
	w1 := h1.Bins[i].Weight

	i = sort.Search(len(h2.Bins), func(i int) bool {
		b := h2.Bins[i]
		return p.Y < b.Min
	})
	if i == len(h2.Bins) {
		i--
	}
	w2 := h2.Bins[i].Weight

	c1 := s1.Color.(color.RGBA)
	c2 := s2.Color.(color.RGBA)

	f := (w1 * width1) / (w1*width1 + w2*width2)
	c := color.RGBA{
		R: uint8(float64(c1.R)*f + float64(c2.R)*(1-f)),
		G: uint8(float64(c1.G)*f + float64(c2.G)*(1-f)),
		B: uint8(float64(c1.B)*f + float64(c2.B)*(1-f)),
		A: uint8(float64(c1.A)*f + float64(c2.A)*(1-f)),
	}

	d := s1.Shape
	if w2 > w1 {
		d = s2.Shape
	}

	return draw.GlyphStyle{
		Color:  c,
		Radius: vg.Points(2),
		Shape:  d,
	}
}

func (c *covPlot) render() error {
	if c == nil {
		return nil
	}

	n := c.n
	if n == 0 {
		return nil
	}

	plots := make([][]*plot.Plot, n)
	for i := 0; i < n; i++ {
		plots[i] = make([]*plot.Plot, n)
		for j := 0; j < n; j++ {
			plots[i][j] = plot.New()
		}
	}

	hists := make([]*plotter.Histogram, n)
	for i := 0; i < n; i++ {
		h, err := makeHist(c.hists[i], i)
		if err != nil {
			return err
		}
		h.Color = h.FillColor
		hists[i] = h
		plots[i][i].Add(h)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i == j {
				continue // Skip histogram.
			}

			p, err := plotter.NewScatter(c.scatters[i][j])
			if err != nil {
				return err
			}
			p.GlyphStyleFunc = func(i, j int) func(int) draw.GlyphStyle {
				return func(idx int) draw.GlyphStyle {
					return makeGlyphStyleMix(hists[i], hists[j], c.scatters[i][j][idx], draw.GlyphStyle{
						Color:  plotutil.Color(i),
						Radius: vg.Points(1),
						Shape:  draw.CircleGlyph{},
					}, draw.GlyphStyle{
						Color:  plotutil.Color(j),
						Radius: vg.Points(1),
						Shape:  draw.CircleGlyph{},
					})
				}
			}(i, j)
			plots[i][j].Add(p)
		}
	}

	// Formatting.
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			// Hack NominalAxes to remove ticks and numbers.
			plots[i][j].NominalX("")
			plots[i][j].NominalY("")
			if i == 0 {
				plots[i][j].Title.Text = seriesName(c.series[j])
			}
			if j == 0 {
				plots[i][j].Y.Label.Text = seriesName(c.series[i])
			}
		}
	}

	const svg = false

	var canvas vg.CanvasSizer

	if svg {
		canvas = vgsvg.New(1*font.Length(n)*vg.Inch, 1*font.Length(n)*vg.Inch)
	} else {
		canvas = vgimg.New(1*font.Length(n)*vg.Inch, 1*font.Length(n)*vg.Inch)
	}
	dc := draw.New(canvas)

	t := draw.Tiles{
		Rows: n,
		Cols: n,
	}
	canvases := plot.Align(plots, t, dc)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if plots[i][j] != nil {
				plots[i][j].Draw(canvases[i][j])
			}
		}
	}

	name := "cov.png"
	if svg {
		name = "cov.svg"
	}
	w, err := os.Create(name)
	if err != nil {
		return err
	}
	defer w.Close()
	if svg {
		if _, err := canvas.(*vgsvg.Canvas).WriteTo(w); err != nil {
			return err
		}
	} else {
		png := vgimg.PngCanvas{Canvas: canvas.(*vgimg.Canvas)}
		if _, err := png.WriteTo(w); err != nil {
			return err
		}
	}

	return nil
}

func CovInit(n int) {
	PcovInit(&globalPlot, n)
}

func PcovInit(p *Plot, n int) {
	s := make([][]plotter.XYs, n)
	for i := 0; i < n; i++ {
		s[i] = make([]plotter.XYs, n)
	}
	p.covPlot = &covPlot{
		n:        n,
		scatters: s,
		hists:    make([]plotter.Values, n),
		series:   make([]seriesKey, n),
		buf:      make([]float64, n),
	}
}

func Cov(v float64) {
	s := runtimeSeriesKey(2)
	Covf(Series(s), v)
}

func Covf(f *Format, v float64) {
	Pcovf(&globalPlot, f, v)
}

func Pcov(p *Plot, v float64) {
	s := runtimeSeriesKey(2)
	Pcovf(p, Series(s), v)
}

func Pcovf(p *Plot, f *Format, v float64) {
	if p.covPlot == nil {
		panic("Pcovf called before PcovInit")
	}
	c := p.covPlot
	if c.n <= p.i {
		panic("Pcovf called too many times before PcovFlush")
	}
	p.covPlot.series[p.i] = f.seriesKey
	p.covPlot.buf[p.i] = v
	p.i++
}

func CovFlush() {
	PcovFlush(&globalPlot)
}

func PcovFlush(p *Plot) {
	if p.covPlot == nil {
		panic("PcovFlush plot called before CovInit or PcovInit")
	}
	c := p.covPlot
	if c.i < c.n {
		panic("PcovFlush called before all values buffered")
	}

	n := c.n
	buf := c.buf

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i == j {
				c.hists[i] = append(c.hists[i], buf[i])
				continue
			}
			c.scatters[i][j] = append(c.scatters[i][j], plotter.XY{
				X: buf[i],
				Y: buf[j],
			})
		}
	}

	// Reset counter.
	p.i = 0
}
