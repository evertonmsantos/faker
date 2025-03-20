# Fake Data Generator (GO)

Este projeto é um gerador de dados falsos em **Golang**, útil para testes, preenchimento de bases de dados fictícias e desenvolvimento de aplicações.

---

## 📌 Funcionalidades

O **Fake Data Generator** possui as seguintes funções:

### 1️⃣ Nome Completo
```go
faker.Fullname()
```
📌 **Descrição**: Gera um nome completo aleatório, composto por primeiro nome e dois sobrenomes.  
📌 **Exemplo de saída**:  
```txt
Maria Fernanda Laurenco Frachin
```

---

### 2️⃣ E-mail
```go
faker.Email()
```
📌 **Descrição**: Gera um e-mail aleatório baseado em um nome fictício, incluindo um número aleatório e um domínio popular.  
📌 **Exemplo de saída**:  
```txt
rebecapegovaropagot442@protonmail.com
```

---

### 3️⃣ Telefone
```go
faker.Phone(formatado bool)
```
📌 **Descrição**: Gera um número de telefone celular válido no Brasil. Pode ser retornado formatado (`(XX) 9XXXX-XXXX`) ou não (`XX9XXXXXXXX`).  
📌 **Exemplo de saída**:  
```txt
(12) 97206-4710
```

---

### 4️⃣ CPF
```go
faker.Cpf(formatado bool)
```
📌 **Descrição**: Gera um **CPF válido**, com dígitos verificadores corretos. Pode ser formatado (`XXX.XXX.XXX-XX`) ou não (`XXXXXXXXXXX`).  
📌 **Exemplo de saída**:  
```txt
528.049.906-48
```

---

### 5️⃣ CNPJ
```go
faker.Cnpj(formatado bool)
```
📌 **Descrição**: Gera um **CNPJ válido**, seguindo as regras oficiais de validação. Pode ser formatado (`XX.XXX.XXX/XXXX-XX`) ou não (`XXXXXXXXXXXXXX`).  
📌 **Exemplo de saída**:  
```txt
17.550.122/0001-54
```

---

### 6️⃣ Endereço (CEP)
```go
faker.Endereco()
```
📌 **Descrição**: Retorna um endereço aleatório com informações como CEP, logradouro, bairro, cidade, estado e código IBGE.  
📌 **Exemplo de saída**:  
```txt
&{60840-030 Rua João Bezerra   Messejana Fortaleza CE Ceará Nordeste 2304400  85 1389}
```

---

## ⚠️ Aviso Legal
Os dados gerados por este projeto são **totalmente fictícios** e não possuem qualquer vínculo com informações reais. **Não utilize para fraudes ou qualquer atividade ilegal.**

---

## 🚀 Como Usar

### 1️⃣ Instalar o pacote
```sh
go get github.com/evertonmsantos/faker
```

### 2️⃣ Exemplo de uso

Crie um arquivo `main.go` e adicione o seguinte código:

```go
package main

import (
	"github.com/evertonmsantos/faker"
	"fmt"
)

func main() {
	fmt.Println(faker.Fullname())
	fmt.Println(faker.Email())
	fmt.Println(faker.Phone(true))
	fmt.Println(faker.Cpf(true))
	fmt.Println(faker.Cnpj(true))
	fmt.Println(faker.Address())
}
```

### 3️⃣ Executar o projeto
```sh
go run main.go
```

Isso irá gerar dados aleatórios conforme as funções listadas acima.

---

## 📜 Licença
Este projeto é de código aberto e pode ser utilizado livremente para fins educacionais e de testes.

