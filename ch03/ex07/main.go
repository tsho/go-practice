//ニュートン法の利用

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
			z := complex(x, y)
			// Image point (px, py) represents complex value z.
			img.Set(px, py, newton(z))
		}
	}
	png.Encode(os.Stdout, img) // NOTE: ignoring errors
}

// f(x) = x^4 - 1
//
// z' = z - f(z)/f'(z)
//    = z - (z^4 - 1) / (4 * z^3)
//    = z - (z - 1/z^3) / 4
func newton(z complex128) color.Color {
	const iterations = 100
	const contrast = 15
	for i := uint8(0); i < iterations; i++ {
		z -= (z - 1/(z*z*z)) / 4
		if cmplx.Abs(z*z*z*z-1) < 1e-6 {
			return color.Gray{255 - contrast*i}
		}
	}
	return color.Black
}
