package main

import (
	"flag"
	"fmt"
	"image/color"
	"log"

	"github.com/wmwv/cosmo"
	"go-hep.org/x/hep/hplot"
	"gonum.org/v1/plot/vg"
	"gonum.org/v1/plot/vg/draw"
)

func main() {

	var (
		h0  = flag.Float64("H0", 70, "Hubble constant at z=0 [km/s/Mpc]")
		om0 = flag.Float64("Omega0", 0.3, "Matter density at z=0")
	)

	flag.Parse()

	cos := cosmo.FlatLCDM{H0: *h0, Om0: *om0}
	fmt.Printf("%#v\n", cos)

	zs := []float64{0.5, 1, 2, 3}
	var ys []float64
	for _, z := range zs {
		v := cos.DistanceModulus(z)
		fmt.Printf("dist-modulus[z=%e]=%v\n", z, v)
		ys = append(ys, v)
	}

	p := hplot.New()
	p.Title.Text = fmt.Sprintf("%#v", cos)
	p.X.Label.Text = "z"
	p.Y.Label.Text = "Distance Modulus"

	scatter := hplot.NewS2D(hplot.ZipXY(zs, ys))
	scatter.Color = color.RGBA{R: 0xff, A: 0xff}
	scatter.Shape = draw.CircleGlyph{}
	scatter.Radius = 2.5
	p.Add(scatter)
	p.Add(hplot.NewGrid())

	err := p.Save(20*vg.Centimeter, -1, "plot.png")
	if err != nil {
		log.Fatal(err)
	}
}
