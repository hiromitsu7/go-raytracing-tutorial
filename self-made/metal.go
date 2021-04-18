package main

type Metal struct {
	C    Vector
	Fuzz float64
}

func (m Metal) Bounce(input Ray, hit Hit) (bool, Ray) {
	direction := reflect(input.Direction, hit.Normal)
	bouncedRay := Ray{hit.Point, direction.Add(RandomUnitVector().MultiplyScalar(m.Fuzz))}
	bounced := direction.Dot(hit.Normal) > 0
	return bounced, bouncedRay
}

func (m Metal) Color(hitPoint Vector) Vector {
	return m.C
}

func reflect(v Vector, n Vector) Vector {
	b := 2 * v.Dot(n)
	return v.Substract(n.MultiplyScalar(b))
}
