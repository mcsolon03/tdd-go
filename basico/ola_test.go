package basico 

import "testing"

func TestOla(t *testing.T) {

	verificarMensagemCorreta := func(t *testing.T, resultado, esperado string) {
		t.Helper()

		if resultado != esperado {
			t.Errorf("resultado '%s', esperado '%s'", resultado, esperado)
		}

	}

	t.Run("diz 'Olá Mundo' quando o nome não for informado", func(t *testing.T) {
		resultado := Ola("", "")
		esperado := "Olá Mundo!"
		verificarMensagemCorreta(t, resultado, esperado)
	})

	t.Run("diz o nome correto, quando o nome for informado", func(t *testing.T) {
		resultado := Ola("Glêsio", "")
		esperado := "Olá Glêsio!"
		verificarMensagemCorreta(t, resultado, esperado)
	})
	t.Run("realizando saudação em espanhol", func(t *testing.T) {
		resultado := Ola("Glêsio", "espanhol")
		esperado := "Holla Glêsio!"
		verificarMensagemCorreta(t, resultado, esperado)
	})

	t.Run("realizando saudação em frances", func(t *testing.T) {
		resultado := Ola("Glêsio", "frances")
		esperado := "Bonjour Glêsio!"
		verificarMensagemCorreta(t, resultado, esperado)
	})
}