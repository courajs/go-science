package main

import (
	"fmt"
	"math"
	"math/rand"
	"sort"
)

type Point struct {
	x, y int
}

// Magnitude calculates distance from the origin
func (p Point) Magnitude() float64 {
	return math.Sqrt(float64(p.x*p.x + p.y*p.y))
}

func (p Point) String() string {
	return fmt.Sprintf("{%v %v}:%.4g", p.x, p.y, p.Magnitude())
}

type Points []Point

func (p Points) Last() Point {
	return p[len(p)-1]
}

// Sort interface
func (p Points) Len() int {
	return len(p)
}
func (p Points) Less(i, j int) bool {
	return p[i].Magnitude() < p[j].Magnitude()
}
func (p Points) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func getPoints() []Point {
	points := make([]Point, 50)
	for i := range points {
		points[i] = Point{rand.Intn(50), rand.Intn(50)}
	}
	return points
}

func kNearest(points []Point, k int) []Point {
	result := Points(make([]Point, k))

	i := 0

	for i < k {
		result[i] = points[i]
		i++
	}

	sort.Sort(result)

	for i < len(points) {
		if points[i].Magnitude() < result[k-1].Magnitude() {
			result[k-1] = points[i]
			sort.Sort(result)
		}
		i++
	}

	return result
}

func main() {
	p := getPoints()
	// fmt.Println(p)
	fmt.Println(kNearest(p, 4))
}
