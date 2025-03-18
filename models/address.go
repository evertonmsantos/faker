package models

// Struct Address para representar os dados do CEP
type Address struct {
	Cep         string
	Logradouro  string
	Complemento string
	Unidade     string
	Bairro      string
	Localidade  string
	Uf          string
	Estado      string
	Regiao      string
	Ibge        string
	Gia         string
	Ddd         string
	Siafi       string
}
