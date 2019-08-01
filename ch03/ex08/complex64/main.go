// complex128
// ex05と同じ内容

package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)

var palette = []color.Color{
	color.RGBA{0xff, 0xb2, 0xb2, 0xFF},
	color.RGBA{0xff, 0xa8, 0xd3, 0xFF},
	color.RGBA{0xff, 0xad, 0xff, 0xFF},
	color.RGBA{0xd6, 0xad, 0xff, 0xFF},
	color.RGBA{0xad, 0xad, 0xff, 0xFF},
	color.RGBA{0xad, 0xd6, 0xff, 0xFF},
}


func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex64(x, y)
			// Image point (px, py) represents complex value z.
			img.Set(px, py, mandelbrot(z))
		}
	}
	png.Encode(os.Stdout, img) // NOTE: ignoring errors
}

func mandelbrot(z complex64) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex64
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return palette[n % uint8(len(palette))]
		}
	}
	return color.RGBA{0xad, 0xff, 0xff, 0xFF}
}

//!-

// Some other interesting functions:

func acos(z complex64) color.Color {
	v := cmplx.Acos(z)
	blue := uint8(real(v)*128) + 127
	red := uint8(imag(v)*128) + 127
	return color.YCbCr{192, blue, red}
}

func sqrt(z complex64) color.Color {
	v := cmplx.Sqrt(z)
	blue := uint8(real(v)*128) + 127
	red := uint8(imag(v)*128) + 127
	return color.YCbCr{128, blue, red}
}

