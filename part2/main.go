package main

import (
	"fmt"
	"math"
	"os"
)

//##########################################################
type Vector struct {
	X, Y, Z float64
}

func (v Vector) MultiplyScalar(t float64) Vector {
	return Vector{X: v.X * t, Y: v.Y * t, Z: v.Z * t}
}

func (v Vector) Add(v1 Vector) Vector {
	return Vector{X: v.X + v1.X, Y: v.Y + v1.Y, Z: v.Z + v1.Z}
}

func (v Vector) Substruct(v1 Vector) Vector {
	return Vector{X: v.X - v1.X, Y: v.Y - v1.Y, Z: v.Z - v1.Z}
}

func (v Vector) Dot(v1 Vector) float64 {
	return v.X*v1.X + v.Y*v1.Y + v.Z*v1.Z
}

func (v Vector) Normalize() Vector {
	length := math.Sqrt(v.Dot(v))
	return Vector{X: v.X / length, Y: v.Y / length, Z: v.Z / length}
}

//##########################################################
type Ray struct {
	Origin, Direction Vector
}

func (r Ray) Point(t float64) Vector {
	b := r.Direction.MultiplyScalar(t)
	a := r.Origin
	return a.Add(b)
}

func (r Ray) HitSphere(s Sphere) bool {
	oc := r.Origin.Substruct(s.Center)
	a := r.Direction.Dot(r.Direction)
	b := oc.Dot(r.Direction)
	c := oc.Dot(oc) - s.Radius*s.Radius
	discriminant := b*b - a*c
	return discriminant > 0
}

func (r Ray) Color() Vector {
	sphere := Sphere{Center: Vector{0.0, 0.0, 1.0}, Radius: 0.5}

	if r.HitSphere(sphere) {
		return Vector{1.0, 0.0, 0.0}
	}

	unit := r.Direction.Normalize()
	t := 0.5 * (unit.Y + 1.0)
	white := Vector{1.0, 1.0, 1.0}
	blue := Vector{0.5, 0.7, 1.0}
	return white.MultiplyScalar(1.0 - t).Add(blue.MultiplyScalar(t))
}

//##########################################################
type Sphere struct {
	Center Vector
	Radius float64
}

//##########################################################
func check(e error, s string) {
	if e != nil {
		fmt.Fprintf(os.Stderr, s, e)
		os.Exit(1)
	}
}

func main() {
	nx := 400
	ny := 200

	const color float64 = 255.99

	f, err := os.Create("out.ppm")

	defer f.Close()

	check(err, "Error opening file: %v\n")

	_, err = fmt.Fprintf(f, "P3\n%d %d\n255\n", nx, ny)

	check(err, "Error writting to file: %v\n")

	lowerLeft := Vector{-2.0, -1.0, -1.0}
	horizontal := Vector{4.0, 0.0, 0.0}
	vertical := Vector{0.0, 2.0, 0.0}
	origin := Vector{0.0, 0.0, 0.0}

	for j := ny - 1; j >= 0; j-- {
		for i := 0; i < nx; i++ {
			// 0〜1の値
			u := float64(i) / float64(nx)
			v := float64(j) / float64(ny)
			// 0〜4, 0〜2の座標の点
			position := horizontal.MultiplyScalar(u).Add(vertical.MultiplyScalar(v))
			// -2〜2, -1〜1の座標の点
			direction := lowerLeft.Add(position)

			rgb := Ray{origin, direction}.Color()

			ir := int(color * rgb.X)
			ig := int(color * rgb.Y)
			ib := int(color * rgb.Z)

			_, err = fmt.Fprintf(f, "%d %d %d\n", ir, ig, ib)
			check(err, "Error writting to file: %v\n")
		}
	}
}
