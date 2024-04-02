package calculator

import (
	"math"
)

func CalculatorMath(num1 float64, num2 float64, value string) (result float64) {

	// Verificar se os parâmetros de entrada estão corretos
	if value == "" {
		panic("O valor não foi informado")
	}

	for {
		switch value {
		case "Somar":
			result = num1 + num2
			return
		case "Subtrair":
			result = num1 - num2
			return
		case "Multiplicar":
			result = num1 * num2
			return
		case "Dividir":
			if num2 == 0 {
				panic("O divisor não pode ser zero")
			}
			result = num1 / num2
			return
		case "Potencia":
			result = math.Pow(num1, num2)
			return
		}
	}
}
