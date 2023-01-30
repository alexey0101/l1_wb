package main

import (
	"fmt"
	"math"
)

type Point struct {
	x, y float64
}

func (basePoint *Point) Distance(point *Point) float64 {
	return math.Sqrt(math.Pow(basePoint.x-point.x, 2) + math.Pow(basePoint.y-point.y, 2))
}

func (point *Point) X() float64 {
	return point.x
}

func (point *Point) Y() float64 {
	return point.y
}

func NewPoint(x, y float64) *Point {
	return &Point{x: x, y: y}
}

func main() {
	point1 := NewPoint(15.23, 22.4)
	point2 := NewPoint(4, 15)

	fmt.Println(point1.Distance(point2))
}
