package main

import (
	"fmt"
	"math"
	"math/rand"
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
type Camera struct {
	Position, Direction Vector
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

// 球に当たったか判定し、当たった場合は当たった点、法線ベクトルを返す
func (r Ray) HitSphere(s Sphere) (bool, Vector, Vector) {
	oc := r.Origin.Substruct(s.Center)
	a := r.Direction.Dot(r.Direction)
	b := 2.0 * oc.Dot(r.Direction)
	c := oc.Dot(oc) - s.Radius*s.Radius
	// 二次関数の頂点のy
	discriminant := b*b - 4.0*a*c
	hit := discriminant >= 0.0

	if hit {
		t1 := (-b - math.Sqrt(b*b-4.0*a*c)) / a / 2.0
		hitPoint1 := r.Point(t1)
		sub1 := hitPoint1.Substruct(r.Origin)
		length21 := sub1.Dot(sub1)
		normal1 := hitPoint1.Substruct(s.Center).Normalize()

		t2 := (-b + math.Sqrt(b*b-4.0*a*c)) / a / 2.0
		hitPoint2 := r.Point(t2)
		sub2 := hitPoint2.Substruct(r.Origin)
		length22 := sub2.Dot(sub2)
		normal2 := hitPoint2.Substruct(s.Center).Normalize()

		if length21 <= length22 {
			return true, hitPoint1, normal1
		} else {
			return true, hitPoint2, normal2
		}
	} else {
		return false, Vector{}, Vector{}
	}
}

func (r Ray) HitGround() (bool, Vector) {
	isHit := false
	if r.Direction.Y < 0.0 {
		isHit = true
	}

	hitPoint := Vector{-r.Origin.Y / r.Direction.Y * r.Direction.X, 0.0, -r.Origin.Y / r.Direction.Y * r.Direction.Z}

	return isHit, hitPoint
}

func (r Ray) Color() Vector {

	sphere := Sphere{Center: Vector{0.0, 1.0, 2.0}, Radius: 1.0}

	hitSphere, hitSpherePoint, sPhereNormal := r.HitSphere(sphere)
	hitGround, hitGroundPoint := r.HitGround()

	// 球に当たった場合
	if hitSphere {
		if math.Abs(hitSpherePoint.X) < 0.01 && math.Abs(hitSpherePoint.Y) < 0.01 {
			fmt.Println(hitSpherePoint)
		}
		return Vector{0.5*sPhereNormal.Y + 0.5, 0.0, 0.0}
	}

	// 地面に当たった場合
	if hitGround {
		// 球の中心からの距離の2乗
		len2 := math.Pow(hitGroundPoint.X-sphere.Center.X, 2) + math.Pow(hitGroundPoint.Z-sphere.Center.Z, 2)
		// 半径の2乗
		rad2 := math.Pow(sphere.Radius, 2)

		baseColor := Vector{0.0, 0.5, 0.5}

		floorFactor := (hitGroundPoint.X - math.Floor(hitGroundPoint.X) + hitGroundPoint.Z - math.Floor(hitGroundPoint.Z)) / 2.0
		floorColor := Vector{0.2, 0.2, 0.2}.MultiplyScalar(floorFactor)

		shadowFactor := 1.0
		if len2 <= rad2 {
			shadowFactor = (1 - (rad2-len2)/rad2)
		}

		return baseColor.MultiplyScalar(shadowFactor).Add(floorColor)
	}

	// 何も当たらなかった場合
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

	nx := 800
	ny := 400

	const color float64 = 255.99

	f, err := os.Create("out.ppm")

	defer f.Close()

	check(err, "Error opening file: %v\n")

	_, err = fmt.Fprintf(f, "P3\n%d %d\n255\n", nx, ny)

	check(err, "Error writting to file: %v\n")

	lowerLeft := Vector{-2.0, -1.0, 1.0}
	horizontal := Vector{4.0, 0.0, 0.0}
	vertical := Vector{0.0, 2.0, 0.0}
	pointOfView := Vector{0.0, 0.6, 0.0}

	for j := ny - 1; j >= 0; j-- {
		for i := 0; i < nx; i++ {
			rgb := Vector{}
			samplingCount := 500
			for s := 0; s < samplingCount; s++ {
				// 0〜1の値
				u := (float64(i) + rand.Float64()) / float64(nx)
				v := (float64(j) + rand.Float64()) / float64(ny)
				// 0〜4, 0〜2の座標の点
				position := horizontal.MultiplyScalar(u).Add(vertical.MultiplyScalar(v))
				// -2〜2, -1〜1の座標の点
				direction := lowerLeft.Add(position).Normalize()

				rgb = rgb.Add(Ray{pointOfView, direction}.Color())
			}

			ir := int(color * rgb.X / float64(samplingCount))
			ig := int(color * rgb.Y / float64(samplingCount))
			ib := int(color * rgb.Z / float64(samplingCount))

			_, err = fmt.Fprintf(f, "%d %d %d\n", ir, ig, ib)
			check(err, "Error writting to file: %v\n")
		}
	}
}
