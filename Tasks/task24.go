package main

import (
	"fmt"
	"math"
)

type Point struct {
	x, y int
}

func Create(x, y int) *Point {
	return &Point{x, y}
}
func (p *Point) DistanceToPoint(target *Point) float64 {
	return math.Sqrt(math.Pow(float64(p.x-target.x), 2.0) + math.Pow(float64(p.y-target.y), 2.0))
}
func main() {
	p1 := Create(0, 3)
	p2 := Create(4, 0)
	fmt.Println("Distance: ", p1.DistanceToPoint(p2))
}
