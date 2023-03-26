package models

type Persegi struct {
	Sisi, Luas, Keliling int
}

func (p *Persegi) HitungKeliling() {
	p.Keliling = p.Sisi * 4
}
func (p *Persegi) HitungLuas() {
	p.Luas = p.Sisi * p.Sisi
}
