package faker

import (
	"math/rand"
	"strconv"
)

// Cpf gera um CPF válido, podendo ser formatado ou não.
func Cpf(formatado bool) string {
	var cpf [11]int

	for {
		allSame := true
		for i := 0; i < 9; i++ {
			cpf[i] = rand.Intn(10)
			if i > 0 && cpf[i] != cpf[0] {
				allSame = false
			}
		}
		if !allSame {
			break
		}
	}

	sum := 0
	for i := 0; i < 9; i++ {
		sum += cpf[i] * (10 - i)
	}
	resto := sum % 11
	cpf[9] = 0
	if resto >= 2 {
		cpf[9] = 11 - resto
	}

	sum = 0
	for i := 0; i < 10; i++ {
		sum += cpf[i] * (11 - i)
	}
	resto = sum % 11
	cpf[10] = 0
	if resto >= 2 {
		cpf[10] = 11 - resto
	}

	cpfStr := ""
	for _, digito := range cpf {
		cpfStr += strconv.Itoa(digito)
	}

	if formatado {
		return cpfStr[:3] + "." + cpfStr[3:6] + "." + cpfStr[6:9] + "-" + cpfStr[9:]
	}
	return cpfStr
}
