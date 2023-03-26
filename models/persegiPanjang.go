package models

type PersegiPanjang struct {
	Panjang, Lebar, Luas, Keliling int
}

func (pp *PersegiPanjang) HitungLuas() {
	pp.Luas = pp.Panjang * pp.Lebar
}
func (pp *PersegiPanjang) HitungKeliling() {
	pp.Keliling = 2 * (pp.Lebar + pp.Panjang)
}
