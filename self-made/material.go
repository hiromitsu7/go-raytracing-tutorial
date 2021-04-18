package main

type Material interface {
	Bounce(input Ray, hit Hit) (bool, Ray)
	Color(hitPoint Vector) Vector
}
