package models

import "math"

type Lingkaran struct {
	JariJari, Keliling, Luas float64
}

func (ling *Lingkaran) HitungLuas() {
	ling.Luas = math.Pi * math.Pow(ling.JariJari, 2)
}
func (ling *Lingkaran) HitungKeliling() {
	ling.Keliling = math.Pi * (ling.JariJari * 2)
}
