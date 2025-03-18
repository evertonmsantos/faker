package faker

import (
	"fmt"
	"math/rand"
	"strconv"
)

// GeneratePhone gera um número de telefone brasileiro aleatório, podendo ser formatado ou não.
func Phone(formatado bool) string {
	validDDDs := []string{
		"11", "12", "13", "14", "15", "16", "17", "18", "19",
		"21", "22", "24", "27", "28",
		"31", "32", "33", "34", "35", "37", "38",
		"41", "42", "43", "44", "45", "46",
		"47", "48", "49",
		"51", "53", "54", "55",
		"61", "62", "63", "64", "65", "66", "67", "68", "69",
		"71", "73", "74", "75", "77", "79",
		"81", "82", "83", "84", "85", "86", "87", "88", "89",
		"91", "92", "93", "94", "95", "96", "97", "98", "99",
	}
	ddd := validDDDs[rand.Intn(len(validDDDs))]

	phone := "9"
	for i := 0; i < 8; i++ {
		phone += strconv.Itoa(rand.Intn(10))
	}

	if formatado {
		return fmt.Sprintf("(%s) %s-%s", ddd, phone[:5], phone[5:])
	}
	return ddd + phone
}
