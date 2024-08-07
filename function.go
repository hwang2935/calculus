package calculus

import (
	"errors"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
)

type Domain interface {
	Contains(Point) bool
	Bounds() []float64 // x1min, x1max, x2min, x2max, ...
	Dim() int
}

type Function interface {
	Domain
	Map(Point) Point
}

type RealValued interface {
	Function
	Eval(Point) float64
}

func GraphRealValue(f RealValued, width, height int, filename string) error {
	dim := f.Dim()
	if dim != 1 && dim != 2 {
		return errors.New("GraphRealValued: Dimension must be either 1 or 2.")
	}
	switch dim {
	case 1:
		return Graph2D(f, width, height, filename)
	case 2:
		return Graph3D(f, width, height, filename)
	}
	return nil
}

func Graph2D(f RealValued, w, h int, fn string) error {
	img := image.NewRGBA(image.Rect(0, 0, w, h))
	bgColor := color.RGBA{255, 255, 255, 255}
	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			img.Set(x, y, bgColor)
		}
	}
	for i := 0; i < w; i++ {
		img.Set(i, h/2, color.RGBA{0, 0, 0, 255})
	}
	for j := 0; j < h; j++ {
		img.Set(w/2, j, color.RGBA{0, 0, 0, 255})
	}
	bdd := f.Bounds()
	dx := (bdd[1] - bdd[0]) / float64(w)
	for i := 0; i < w; i++ {
		x := bdd[0] + float64(i)*dx
		y := f.Eval(Point{x})
		j := int((bdd[3]- y) / (bdd[3] - bdd[2]) * float64(h))
		fmt.Println(i,j)
		img.Set(i, j, color.RGBA{0, 0, 0, 255})
	}
	file, err := os.Create(fn)
	if err != nil {
		return err
	}
	defer file.Close()
	err = png.Encode(file, img)
	if err != nil {
		return err
	}
	fmt.Println(err)
	return nil
}

func Graph3D(f RealValued, w, h int, fn string) error {
	img := image.NewRGBA(image.Rect(0, 0, w, h))
	bgColor := color.RGBA{255, 255, 255, 255}
	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			img.Set(x, y, bgColor)
		}
	}
	for i := 0; i < w; i++ {
		img.Set(i, h/2, color.RGBA{0, 0, 0, 255})
	}
	for j := 0; j < h; j++ {
		img.Set(w/2, j, color.RGBA{0, 0, 0, 255})
	}
	bdd := f.Bounds()
	dx := (bdd[1] - bdd[0]) / float64(w)
	dy := (bdd[3] - bdd[2]) / float64(h)
	t1, t2 := 1., 1.
	for i := 0; i < w; i++ {
		for j := 0; j < h; j++ {
			x := bdd[0] + float64(i)*dx
			y := bdd[2] + float64(j)*dy
			p := Point{x, y, f.Eval(Point{x, y})}
			p = Rotation3D(0, t1).Transform(p)
			p = Rotation3D(1, t2).Transform(p)
			a, b := Pixel(p[0], p[1], bdd, w, h)
			img.Set(a, b, color.RGBA{0, 0, 0, 255})
		}
	}
	file, err := os.Create(fn)
	if err != nil {
		return err
	}
	defer file.Close()
	err = png.Encode(file, img)
	if err != nil {
		return err
	}
	return nil
}

func Pixel(x, y float64, bdd []float64, w, h int) (int, int) {
	a := int((x - bdd[0]) / (bdd[1] - bdd[0]) * float64(w))
	b := int((y - bdd[2]) / (bdd[3] - bdd[2]) * float64(h))
	return a, b
}
