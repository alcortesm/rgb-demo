package main

import (
	"image"
	"math"
	"testing"
)

func testModulus(p image.Point, expected float64, t *testing.T) {
	returned := modulus(p)
	if expected != returned {
		t.Error("expected ", expected, ", got ", returned)
	}
}

func TestModulus(t *testing.T) {
	testModulus(image.Point{0, 0}, 0.0, t)
	testModulus(image.Point{3, 4}, 5.0, t)
	testModulus(image.Point{1, 1}, math.Sqrt2, t)
	testModulus(image.Point{-1, 1}, math.Sqrt2, t)
	testModulus(image.Point{1, -1}, math.Sqrt2, t)
	testModulus(image.Point{-1, -1}, math.Sqrt2, t)
	testModulus(image.Point{-4, 3}, 5.0, t)
}

func testDistance(x1, y1, x2, y2 int, expected float64, t *testing.T) {
	a := image.Point{x1, y1}
	b := image.Point{x2, y2}
	returned := distance(a, b)
	if returned != expected {
		t.Error("expected ", expected, ", got ", returned)
	}
}

func TestDistance(t *testing.T) {
	testDistance(0, 0, 0, 0, 0.0, t)
	testDistance(1, 1, 0, 0, math.Sqrt2, t)
	testDistance(2, 3, 5, 7, 5.0, t)
	testDistance(-1, 2, 2, 6, 5.0, t)
	testDistance(-2, 3, -2, 3, 0.0, t)
}

func testContains(c circle, p image.Point, expected bool, t *testing.T) {
	returned := c.contains(p)
	if returned != expected {
		t.Error("circle = ", c, ", point = ", p,
			", expected ", expected, ", got ", returned)
	}
}

func TestContains(t *testing.T) {
	testContains(circle{image.Point{0, 0}, 0},
		image.Point{0, 0}, true, t)
	testContains(circle{image.Point{0, 0}, 0},
		image.Point{1, 0}, false, t)
	testContains(circle{image.Point{0, 0}, 1},
		image.Point{1, 0}, true, t)
	testContains(circle{image.Point{0, 0}, 1},
		image.Point{0, 1}, true, t)
	testContains(circle{image.Point{0, 0}, 1},
		image.Point{1, 1}, false, t)
	testContains(circle{image.Point{0, 0}, 1},
		image.Point{-1, 0}, true, t)
	testContains(circle{image.Point{0, 0}, 1},
		image.Point{0, -1}, true, t)
	testContains(circle{image.Point{0, 0}, 1},
		image.Point{-1, -1}, false, t)
}
