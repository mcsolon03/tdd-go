package basico

func Somar(numeros []int) int {
	total := 0

	for _, numero := range numeros {
		total += numero
	}

	return total
}