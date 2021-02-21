package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math"
	"os"
)

const (
	width     = 1000
	height    = 1000
	a         = 1.0
	k         = 0.05
	rotations = 20
)

func toX(x int) float64 {
	return float64(x - (width / 2))
}

func main() {

	upLeft := image.Point{0, 0}
	lowRight := image.Point{width, height}
	black := color.RGBA{0, 0, 0, 255}
	white := color.RGBA{255, 255, 255, 255}

	fmt.Println("Creating image...")
	img := image.NewRGBA(image.Rectangle{upLeft, lowRight})
	for Px := 0; Px < width; Px++ {
		for Py := 0; Py < height; Py++ {
			img.Set(Px, Py, white)
		}
	}
	for i := 0; i < 360*rotations; i++ {
		x := a * math.Exp(float64(i)*k*math.Pi/180.0) * math.Cos(float64(i)*math.Pi/180.0)
		y := a * math.Exp(float64(i)*k*math.Pi/180.0) * math.Sin(float64(i)*math.Pi/180.0)
		Px := int(mapVal(x, -100.0, 100.0, 0.0, float64(width)-1))
		Py := int(mapVal(y, -100.0, 100, float64(height)-1, 0.0))
		img.Set(Px, Py, black)
		fmt.Printf("x: %v, y: %v, Px: %v, Py: %v\n", x, y, Px, Py)
	}

	// Encode as PNG.
	f, err := os.Create("images/image.png")
	if err != nil {
		fmt.Println(err)
	}
	if err = png.Encode(f, img); err != nil {
		fmt.Println(err)
	}
}

func mapVal(x, imin, imax, omin, omax float64) float64 {
	if x < imin {
		x = imin
	} else if x > imax {
		x = imax
	}
	return (x-imin)*(omax-omin)/(imax-imin) + omin
}
