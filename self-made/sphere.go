package main

import "math"

type Sphere struct {
	Center Vector
	Radius float64
	Material
}

func (s *Sphere) Hit(r Ray, tMin float64, tMax float64) (bool, Hit) {
	oc := r.Origin.Substract(s.Center)
	a := r.Direction.Dot(r.Direction)
	b := 2.0 * oc.Dot(r.Direction)
	c := oc.Dot(oc) - s.Radius*s.Radius
	// 二次関数の頂点のy
	discriminant := b*b - 4.0*a*c
	isHit := discriminant >= 0.0

	hit := Hit{Material: s.Material}

	if isHit {
		// 二次方程式の解は2つある
		// 値が小さい方の解を選ぶ
		t1 := (-b - math.Sqrt(b*b-4.0*a*c)) / a / 2.0
		// t2 := (-b + math.Sqrt(b*b-4.0*a*c)) / a / 2.0

		if tMin < t1 && t1 < tMax {
			hit.T = t1
			hitPoint := r.Point(t1)
			normal := hitPoint.Substract(s.Center).Normalize()
			hit.Normal = normal
			hit.Point = hitPoint
			return true, hit
		}
	}
	return false, Hit{}
}
