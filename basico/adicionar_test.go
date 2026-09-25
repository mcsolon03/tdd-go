package basico

import "testing"

func TestAdicionar(t *testing.T) {
	t.Helper()

	verificarMensagemCorreta := func(t *testing.T, resultado, esperado int64) {
		t.Helper()

		if resultado != esperado {
			t.Errorf("resultado '%d', esperado '%d'", resultado, esperado)
		}
	}

	t.Run("deverá somar dois numero positivos", func(t *testing.T) {
		resultado := Adiciona(2, 2)
		esperado := 4
		verificarMensagemCorreta(t, resultado, int64(esperado))
	})
}