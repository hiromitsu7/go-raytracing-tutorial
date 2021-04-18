package main

import (
	"fmt"
	"math"
	"math/rand"
	"os"
)

const color float64 = 255.99

func check(e error, s string) {
	if e != nil {
		fmt.Fprintf(os.Stderr, s, e)
		os.Exit(1)
	}
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
	sphere1 := Sphere{Center: Vector{0.0, 1.0, 2.0}, Radius: 1.0, Material: Lambertian{Vector{0.8, 0.3, 0.3}}}
	world.Add(&sphere1)
	sphere2 := Sphere{Center: Vector{2.0, 1.0, 3.0}, Radius: 1.0, Material: Lambertian{Vector{0.8, 0.4, 0.4}}}
	world.Add(&sphere2)
	floor := Floor{level: 0.0, Material: Lambertian{Vector{0.3, 0.4, 0.4}}}
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
				isHit, hit := world.Hit(ray, 0.0, math.MaxFloat64)
				// 当たった場合はそのオブジェクトの色
				if isHit {
					rgb = rgb.Add(hit.Material.Color())
				} else {
					// 何にも当たらなかった場合
					rgb = rgb.Add(Vector{0.0, 0.0, 0.5})
				}
			}

			ir := int(color * rgb.X / float64(samplingCount))
			ig := int(color * rgb.Y / float64(samplingCount))
			ib := int(color * rgb.Z / float64(samplingCount))

			_, err = fmt.Fprintf(f, "%d %d %d\n", ir, ig, ib)
			check(err, "Error writting to file: %v\n")
		}
	}
}
