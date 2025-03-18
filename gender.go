package faker

import "math/rand"

// Gender gera um sexo aleatório.
func Gender() string {
	genders := []string{
		"Masculino", "Feminino", "Não-Binário", "Gênero Fluido", "Agênero",
		"Bigênero", "Demiboy", "Demigirl", "Two-Spirit", "Pangênero",
	}
	return genders[rand.Intn(len(genders))]
}
