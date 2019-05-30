package plot

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image"
	"os"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg/draw"
	"gonum.org/v1/plot/vg/vgimg"
)

// Point represents a coordination in the graph.
type Point struct {
	X float64
	Y float64
}

// Scatter prints a scatter graph.
func Scatter(data []Point) error {
	xys := make(plotter.XYs, len(data))
	for i := range data {
		xys[i] = plotter.XY(data[i])
	}
	sc, err := plotter.NewScatter(xys)
	if err != nil {
		return err
	}

	p, err := plot.New()
	if err != nil {
		return err
	}

	p.Add(plotter.NewGrid())
	p.Add(sc)
	const w, h = 1000, 1000
	i := image.NewRGBA(image.Rect(0, 0, w, h))
	canvas := vgimg.PngCanvas{
		Canvas: vgimg.NewWith(
			vgimg.UseImage(i),
		),
	}
	p.Draw(draw.New(canvas))

	var buf bytes.Buffer
	canvas.WriteTo(&buf)
	fmt.Fprint(os.Stdout, "\x1b]1337;File=inline=1:")
	wc := base64.NewEncoder(base64.StdEncoding, os.Stdout)
	canvas.WriteTo(wc)
	wc.Close()
	fmt.Fprint(os.Stdout, "\a")
	return nil
}
