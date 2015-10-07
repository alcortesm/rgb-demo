// Copyright 2015 Alberto Cortés.

// This program generates an RGB demo image: a square image with a
// black background and three colored circles (red, green and blue)
// interseting in a certain pattern. The resulting image is stored in
// a file called "rgb.png".
package main

import (
	"errors"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math"
	"os"
)

func main() {
	const (
		side = 400
		file = "rgb.png"
	)

	// create an image
	i, err := newImg(side)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// create an output file
	// or exit if it already exits
	f, err := os.OpenFile(file,
		os.O_CREATE|os.O_WRONLY|os.O_EXCL, 0666)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// write image to the output file
	if err := png.Encode(f, i); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	os.Exit(0)
}

// A simple rectangular image.
// A zero value img is NOT safe, use newImg method as a ctor
// Implements image.Image.
type img struct {
	side   int
	circle circle
}

// Calling this method "new" would be confusing because of the
// "new()" builtin function, so I have called it "newImg".
//
// As I want to return nil imgs on error, I have to return a
// pointer. For consistency, the rest of the methods also has
// pointer receivers.
func newImg(side int) (*img, error) {
	if side < 1 {
		return nil, errors.New("image side must be greater than 0")
	}
	i := new(img)
	i.side = side
	i.circle = circle{image.Point{200, 200}, 200}
	return i, nil
}

// to implement image.Image
func (i *img) ColorModel() color.Model {
	return color.RGBAModel
}

// to implement image.Image
func (i *img) Bounds() image.Rectangle {
	return image.Rectangle{
		image.Point{0, 0},
		image.Point{i.side, i.side},
	}
}

// to implement image.Image
// TODO: substitute white circle with a red, green and blue circles
func (i *img) At(x, y int) color.Color {
	p := image.Point{x, y}
	if i.circle.contains(p) {
		return color.White
	}
	return color.Black
}

func modulus(p image.Point) float64 {
	return math.Sqrt(float64(p.X*p.X + p.Y*p.Y))
}

func distance(a, b image.Point) float64 {
	return modulus(a.Sub(b))
}

// A simple circle
type circle struct {
	center image.Point
	radius int
}

func (c *circle) contains(p image.Point) bool {
	d := distance(c.center, p)
	return d <= float64(c.radius)
}
