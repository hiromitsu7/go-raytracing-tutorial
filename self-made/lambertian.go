package main

type Lambertian struct {
	C Vector
}

// ランダムな方向に反射
func (l Lambertian) Bounce(input Ray, hit Hit) (bool, Ray) {
	direction := hit.Normal.Add(RandomUnitVector()).Normalize()
	return true, Ray{hit.Point, direction}
}

func (l Lambertian) Color() Vector {
	return l.C
}
