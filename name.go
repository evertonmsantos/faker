package faker

import (
	"math/rand"

	"github.com/evertonmsantos/faker/data"
)

// FirstName retorna um nome aleatório, podendo ser masculino ou feminino.
func FirstName() string {
	names := append(data.Men(), data.Woman()...)
	return names[rand.Intn(len(names))]
}

// Surname retorna um sobrenome aleatório da lista de sobrenomes disponíveis.
func Surname() string {
	surnames := data.Surname()
	return surnames[rand.Intn(len(surnames))]
}

// Fullname gera um nome completo composto por um primeiro nome e dois sobrenomes.
func Fullname() string {
	return FirstName() + " " + Surname() + " " + Surname()
}
