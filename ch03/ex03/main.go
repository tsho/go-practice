package main

import (
	"fmt"
	"math"
)

const (
	width, height = 600, 320            // canvas size in pixels
	cells         = 100                 // number of grid cells
	xyrange       = 30.0                // axis ranges (-xyrange..+xyrange)
	xyscale       = width / 2 / xyrange // pixels per x or y unit
	zscale        = height * 0.4        // pixels per z unit
	angle         = math.Pi / 6         // angle of x, y axes (=30°)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

func main() {
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j)
			bx, by := corner(i, j)
			cx, cy := corner(i, j+1)
			dx, dy := corner(i+1, j+1)
			r, g, b := color(i, j)
			if isFinite(ax) && isFinite(ay) && isFinite(bx) && isFinite(by) &&
				isFinite(cx) && isFinite(cy) && isFinite(dx) && isFinite(dy) {
					fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g' fill=\"rgb(%d,%d,%d)\"/>\n",
					ax, ay, bx, by, cx, cy, dx, dy, r, g, b)
			} else {
				fmt.Printf("!!!!!!!!!!!! skip !!!!!!!!!!!!!!!!!!!!!!!!")
			}
			
		}
	}
	fmt.Println("</svg>")
}

func corner(i, j int) (float64, float64) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	z := f(x, y)

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)

	return math.Sin(r) / r
}

func isFinite(f float64) bool {
	if math.IsInf(f, 0) {
		return false
	}
	if math.IsNaN(f) {
		return false
	}
	return true
}

func color(i, j int) (int, int, int) {
	z := f(xyrange*(float64(i)/cells-0.5), xyrange*(float64(j)/cells-0.5))
	r := 255 * math.Abs(z)
	g := 255 - r
	b := 255 - r
	return int(r), int(g), int(b)
}