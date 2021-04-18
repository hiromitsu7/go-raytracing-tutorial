package main

import "math"

type LambertianTile struct {
	C Vector
}

// ランダムな方向に反射
func (l LambertianTile) Bounce(input Ray, hit Hit) (bool, Ray) {
	direction := hit.Normal.Add(RandomUnitVector()).Normalize()
	return true, Ray{hit.Point, direction}
}

func (l LambertianTile) Color(hitPoint Vector) Vector {
	factor := math.Sin(hitPoint.X*math.Pi) * math.Sin(hitPoint.Z*math.Pi)

	if 0 < factor {
		return l.C.MultiplyScalar(0.5)
	} else {
		return l.C
	}
}
