package matrix

import "math"

type Vertex struct {
	X, Y int
}

func (v *Vertex) Sq(X, Y int) int {
	return v.X*v.X + v.Y*v.Y
}
func (v *Vertex) Sqrt(x int) float64 {
	return math.Sqrt(float64(x))
}
