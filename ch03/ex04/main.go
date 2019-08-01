package main

import (
	"log"
	"net/http"
	"image/color"
	"math"
	"fmt"
	"strconv"
	"net/url"
)

var palette = []color.Color{color.White, color.Black}

const (
	xyrange       = 30.0                // axis ranges (-xyrange..+xyrange)
	angle         = math.Pi / 6         // angle of x, y axes (=30°)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)
var xyscale, zscale float64
var width, height, cells int

func main() {
	handler := func(w http.ResponseWriter, r *http.Request) {
		m, err := url.ParseQuery(r.URL.RawQuery)
		check(err)
		width, _= strconv.Atoi(m["width"][0])
		check(err)
		height, err = strconv.Atoi(m["height"][0])
		check(err)
		cells, err = strconv.Atoi(m["cells"][0])
		check(err)

		xyscale	= float64(width) / 2 / xyrange // pixels per x or y unit
		zscale	= float64(height) * 0.1      // pixels per z unit

		w.Header().Set("Content-Type", "image/svg+xml")
		_, err = fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
				"style='stroke: grey; fill: white; stroke-width: 0.7' "+
				"width='%d' height='%d'>", width, height)

		check(err)

		for i := 0; i < cells; i++ {
			for j := 0; j < cells; j++ {
				ax, ay := corner(i+1, j)
				bx, by := corner(i, j)
				cx, cy := corner(i, j+1)
				dx, dy := corner(i+1, j+1)
				if isFinite(ax) && isFinite(ay) && isFinite(bx) && isFinite(by) &&
					isFinite(cx) && isFinite(cy) && isFinite(dx) && isFinite(dy) {
						fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
						ax, ay, bx, by, cx, cy, dx, dy)
				} else {
					fmt.Printf("!!!!!!!!!!!! skip !!!!!!!!!!!!!!!!!!!!!!!!")
				}
			}
		}
		_, err = fmt.Println("</svg>")

		check(err)
	}

	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
	return 
}

func corner(i, j int) (float64, float64) {
	x := xyrange * (float64(i)/float64(cells) - 0.5)
	y := xyrange * (float64(j)/float64(cells) - 0.5)

	z := f(x, y)

	sx := float64(width)/2 + (x-y)*cos30*xyscale
	sy := float64(height)/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}

//TODO:return  z?
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

func ys(x0, y0, x1, y1, x float64) float64 {
	return y0 + (x - x0)/(x1 - x0) * (y1 - y0)
}


func getColor(z, zmax, zmin float64) string {
	b := 255
	r := 255
	if z > 0 {
		r = int(ys(0, 255, zmax, 0, z))
	} else {
		b = int(ys(zmin, 0, 0, 255, z))
	}

	return "#" + string(fmt.Sprintf("%02x", b)) + "00" + string(fmt.Sprintf("%02x", r))
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}