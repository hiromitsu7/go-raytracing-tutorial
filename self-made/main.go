package main

import (
	"fmt"
	"math"
	"math/rand"
	"os"
)

const colorMax float64 = 255.99

func check(e error, s string) {
	if e != nil {
		fmt.Fprintf(os.Stderr, s, e)
		os.Exit(1)
	}
}

func color(ray Ray, world World, depth int) Vector {
	isHit, hit := world.Hit(ray, 0.0, math.MaxFloat64)
	// 当たった場合はそのオブジェクトの色
	if isHit {
		if depth < 10 {
			bounced, bouncedRay := hit.Bounce(ray, hit)
			if bounced {
				newColor := color(bouncedRay, world, depth+1)
				return hit.Material.Color(hit.Point).Multiply(newColor)
			}
		}
		return Vector{}
	} else {
		// 何にも当たらなかった場合
		return gradient(ray)
	}
}

func gradient(r Ray) Vector {
	v := r.Direction.Normalize()

	t := 0.5 * (v.Y + 1.0)

	white := Vector{1.0, 1.0, 1.0}
	yellow := Vector{0.99, 0.92, 0.69}
	return white.MultiplyScalar(1.0 - t).Add(yellow.MultiplyScalar(t))
}

func main() {

	nx := 600
	ny := 300

	f, err := os.Create("out.ppm")
	defer f.Close()
	check(err, "Error opening file: %v\n")
	_, err = fmt.Fprintf(f, "P3\n%d %d\n255\n", nx, ny)
	check(err, "Error writting to file: %v\n")

	world := World{}
	sphere1 := Sphere{Center: Vector{-3.0, 1.0, 5.0}, Radius: 1.0, Material: Lambertian{Vector{0.8, 0.3, 0.3}}}
	world.Add(&sphere1)

	sphere2 := Sphere{Center: Vector{3.0, 1.0, 4.0}, Radius: 1.0, Material: Lambertian{Vector{0.8, 0.4, 0.4}}}
	world.Add(&sphere2)

	sphere3 := Sphere{Center: Vector{8.0, 2.0, 9.0}, Radius: 2.0, Material: Metal{Vector{1.0, 0.0, 0.0}, 0.0}}
	world.Add(&sphere3)

	sphere4 := Sphere{Center: Vector{2.0, 2.0, 9.0}, Radius: 2.0, Material: Lambertian{Vector{0.0, 1.0, 0.0}}}
	world.Add(&sphere4)

	sphere5 := Sphere{Center: Vector{-2.0, 2.0, 9.0}, Radius: 2.0, Material: Metal{Vector{1.0, 1.0, 1.0}, 0.0}}
	world.Add(&sphere5)

	sphere6 := Sphere{Center: Vector{-20.0, 4.0, 20.0}, Radius: 4.0, Material: Lambertian{Vector{0.0, 1.0, 1.0}}}
	world.Add(&sphere6)

	floor := Floor{level: 0.0, Material: LambertianTile{Vector{0.4, 0.5, 1.0}}}
	world.Add(&floor)

	camera := NewCamera()

	for j := ny - 1; j >= 0; j-- {
		for i := 0; i < nx; i++ {
			rgb := Vector{}
			samplingCount := 50
			for s := 0; s < samplingCount; s++ {
				// 0〜1の値
				u := (float64(i) + rand.Float64()) / float64(nx)
				v := (float64(j) + rand.Float64()) / float64(ny)

				ray := camera.RayAt(u, v)
				rgb = rgb.Add(color(ray, world, 0))
			}

			ir := int(colorMax * rgb.X / float64(samplingCount))
			ig := int(colorMax * rgb.Y / float64(samplingCount))
			ib := int(colorMax * rgb.Z / float64(samplingCount))

			_, err = fmt.Fprintf(f, "%d %d %d\n", ir, ig, ib)
			check(err, "Error writting to file: %v\n")
		}
	}
}
