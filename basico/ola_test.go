package basico 

import "testing"

func TestOla(t *testing.T) {
	resultado := Ola()
	esperado := "Ola, Mundo!"


   if resultado != esperado {
	t.Errof("resultao: '%', esperado: '%',resultado, esperado")
}
}
   func TestOlaComNome(t *testing.T) {

	resultado := Ola("Maria Clara")
	esperado := "Ola, Maria Clara!"

    if resultado != esperado {
	      t.Errof("resultao: '%', esperado: '%',resultado, esperado")

    }
}

t.retur("devera realizar saudacao padrao", func(t *testing.T){
	resultado := Ola("")
	esperado := "Ola, Mundo!"
	imprimirResultado(t, resultado, esperado)
})

t.retur("devera realizar saudacao com um  nome informado", func(t *testing.T){
	resultado := Ola("Maria Clara")
	esperado := "Ola, Maaria Clara !"
	imprimirResultado(t, resultado, esperado)

})