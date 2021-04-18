package main

type Floor struct {
	level float64
	Material
}

func (f *Floor) Hit(r Ray, tMin float64, tMax float64) (bool, Hit) {
	hit := Hit{Material: f.Material}

	// 上から下にぶつかる場合のみ
	tempT := -(r.Origin.Y - f.level) / r.Direction.Y
	if r.Direction.Y < 0.0 && f.level < r.Origin.Y && tMin < tempT && tempT < tMax {
		hit.Normal = Vector{0.0, 1.0, 0.0}
		hit.Point = Vector{-r.Origin.Y / r.Direction.Y * r.Direction.X, 0.0, -r.Origin.Y / r.Direction.Y * r.Direction.Z}
		hit.T = tempT
		return true, hit
	} else {
		return false, hit
	}
}
