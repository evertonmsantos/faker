package faker

import (
	"math/rand"
	"strconv"
)

// Cnpj gera um CNPJ válido, podendo ser formatado ou não.
func Cnpj(formatado bool) string {
	var cnpj [14]int

	for i := 0; i < 8; i++ {
		cnpj[i] = rand.Intn(10)
	}
	cnpj[8], cnpj[9], cnpj[10], cnpj[11] = 0, 0, 0, 1

	multiplicador1 := []int{5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2}
	sum := 0
	for i := 0; i < 12; i++ {
		sum += cnpj[i] * multiplicador1[i]
	}
	resto := sum % 11
	cnpj[12] = 0
	if resto >= 2 {
		cnpj[12] = 11 - resto
	}

	multiplicador2 := []int{6, 5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2}
	sum = 0
	for i := 0; i < 13; i++ {
		sum += cnpj[i] * multiplicador2[i]
	}
	resto = sum % 11
	cnpj[13] = 0
	if resto >= 2 {
		cnpj[13] = 11 - resto
	}

	cnpjStr := ""
	for _, digito := range cnpj {
		cnpjStr += strconv.Itoa(digito)
	}

	if formatado {
		return cnpjStr[:2] + "." + cnpjStr[2:5] + "." + cnpjStr[5:8] + "/" + cnpjStr[8:12] + "-" + cnpjStr[12:]
	}
	return cnpjStr
}
