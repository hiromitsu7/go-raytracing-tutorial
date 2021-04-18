package main

type World struct {
	objects []Hitable
}

func (w *World) Add(h Hitable) {
	w.objects = append(w.objects, h)
}

func (w *World) Hit(r Ray, tMin float64, tMax float64) (bool, Hit) {
	hitAnything := false
	closestT := tMax
	hit := Hit{}

	// 最も近いもの=実際にrayが当たったものを探す
	for _, object := range w.objects {
		isHit, tempHit := object.Hit(r, tMin, closestT)

		if isHit {
			hitAnything = true
			closestT = tempHit.T
			hit = tempHit
		}
	}
	return hitAnything, hit
}
