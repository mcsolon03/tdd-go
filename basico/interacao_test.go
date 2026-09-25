package basico

import "testing"

// func TestIteracao(t *testing.T) {
// 	t.Helper()

// 	verificarMensagemCorreta := func(t *testing.T, esperado, resultado string) {
// 		t.Helper()

// 		if resultado != esperado {
// 			t.Errorf("esperado '%s', mas obteve '%s'", esperado, resultado)
// 		}
// 	}

// 	t.Run("deverá contar o total de repetição", func(t *testing.T) {
// 		resultado := Repetir("a")
// 		esperado := "aaaaa"
// 		verificarMensagemCorreta(t, esperado, resultado)
// 	})

// 	t.Run("deverá contar o total de repetição", func(t *testing.T) {
// 		resultado := Repetir("a")
// 		esperado := "aaaaa"
// 		verificarMensagemCorreta(t, esperado, resultado)
// 	})
// }

func BenchmarkRepetir(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repetir("a")
	}
}