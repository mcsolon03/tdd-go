package basico

import "testing"

func TestSomar(t *testing.T) {
	t.Helper()

	verificarMensagemCorreta := func(t *testing.T, resultado int, esperado int, numeros []int) {
		t.Helper()

		if resultado != esperado {
			t.Errorf("resultado %d, esperado %d, dados %v", resultado, esperado, numeros)
		}
	}

	t.Run("deverá somar uma coleção de 5 numeros", func(t *testing.T) {

		numeros := []int{1, 2, 3, 4, 5}

		resultado := Somar(numeros)
		esperado := 15
		verificarMensagemCorreta(t, resultado, esperado, numeros)
	})
	t.Run("deverá somar uma coleção de qualquer tamanho", func(t *testing.T) {

		numeros := []int{1, 2, 3}

		resultado := Somar(numeros)
		esperado := 6
		verificarMensagemCorreta(t, resultado, esperado, numeros)
	})
}