package basico

import "math"

type Forma interface {
	Area() float64
}

type Retangulo struct {
	Largura float64
	Altura  float64
}

type Circulo struct {
	Raio float64
}

func Perimetro(retangulo Retangulo) float64 {
	return 2 * (retangulo.Altura + retangulo.Largura)
}

func (r Retangulo) Area() float64 {
	return r.Altura * r.Largura
}

func (r Circulo) Area() float64 {
	return math.Pi * r.Raio * r.Raio
}