package main

import (
	"math"
	"math/rand"
)

type Vector struct {
	X, Y, Z float64
}

func RandomUnitVector() Vector {
	r := Vector{rand.Float64(), rand.Float64(), rand.Float64()}
	return r.Normalize()
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

func (v Vector) Length() float64 {
	return math.Sqrt(v.Dot(v))
}
