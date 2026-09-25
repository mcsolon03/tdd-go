package basico

import "fmt"

const espanhol = "espanhol"
const frances = "frances"
const prefixoOlaPortugues = "Olá"
const prefixoOlaEspanhol = "Holla"
const prefixoOlaFrances = "Bonjour"

func Ola(nome string, idioma string) string {

	if nome == "" {
		nome = "Mundo"
	}

	prefixo := prefixoOlaPortugues

	switch idioma {
	case espanhol:
		prefixo = prefixoOlaEspanhol
	case frances:
		prefixo = prefixoOlaFrances
	}

	return fmt.Sprintf("%s %s!", prefixo, nome)
}