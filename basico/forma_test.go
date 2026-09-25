package basico

import (
	"math"
	"testing"
)

func TestFormas(t *testing.T) {
	verificarAreaCorreta := func(t *testing.T, forma Forma, esperado float64) {
		t.Helper()

		resultado := forma.Area()

		if math.Abs(resultado-esperado) > 0.000001 {
			t.Errorf(
				"resultado '%.2f', esperado '%.2f'",
				resultado,
				esperado,
			)
		}
	}

	t.Run("deverá calcular o perímetro de um retângulo", func(t *testing.T) {
		retangulo := Retangulo{
			Largura: 10.0,
			Altura:  10.0,
		}

		resultado := Perimetro(retangulo)
		esperado := 40.0

		if resultado != esperado {
			t.Errorf(
				"resultado '%.2f', esperado '%.2f'",
				resultado,
				esperado,
			)
		}
	})

	t.Run("deverá calcular a área de um retângulo", func(t *testing.T) {
		retangulo := Retangulo{
			Largura: 12.0,
			Altura:  6.0,
		}

		verificarAreaCorreta(t, retangulo, 72.0)
	})

	t.Run("deverá calcular a área de um círculo", func(t *testing.T) {
		circulo := Circulo{
			Raio: 10.0,
		}

		verificarAreaCorreta(t, circulo, math.Pi*100)
	})
}