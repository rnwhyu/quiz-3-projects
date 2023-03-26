package models

type Segitiga struct {
	Alas, Tinggi, Luas, Keliling float64
}

func (s3 *Segitiga) HitungKeliling() {
	s3.Keliling = s3.Alas * 3
}
func (s3 *Segitiga) HitungLuas() {
	s3.Luas = (s3.Alas * s3.Tinggi) / 2
}
