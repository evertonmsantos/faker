package faker

import (
	"math/rand"
	"regexp"
	"strconv"
	"strings"
)

// Email gera um e-mail aleatório baseado no nome completo.
func Email() string {
	domains := []string{"gmail.com", "yahoo.com", "outlook.com", "hotmail.com", "protonmail.com"}
	re := regexp.MustCompile(`[^a-z0-9.]`)
	name := re.ReplaceAllString(strings.ReplaceAll(strings.ToLower(Fullname()), " ", " "), "")
	return name + strconv.Itoa(rand.Intn(1000)) + "@" + domains[rand.Intn(len(domains))]
}
