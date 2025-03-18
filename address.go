package faker

import (
	"database/sql"
	"embed"
	"fmt"
	"io/fs"
	"io/ioutil"
	"math/rand"
	"sync"
	
	"github.com/evertonmsantos/faker/models"
	_ "modernc.org/sqlite"
)

//go:embed data/ceps.db
var cepsDB embed.FS

var (
	db   *sql.DB
	once sync.Once
)

// GetDB mantém uma única conexão ativa (singleton)
func GetDB() *sql.DB {
	once.Do(func() {
		var err error
		dbPath, err := extractDB()
		if err != nil {
			fmt.Printf("Erro ao extrair o banco de dados: %v\n", err)
			return
		}

		db, err = sql.Open("sqlite", dbPath)
		if err != nil {
			fmt.Printf("Erro ao conectar ao banco: %v\n", err)
			return
		}

		if _, err = db.Exec("PRAGMA journal_mode=WAL;"); err != nil {
			fmt.Printf("Erro ao ativar WAL: %v\n", err)
		}
	})
	return db
}

// extractDB cria um arquivo temporário e copia os dados do embed
func extractDB() (string, error) {
	tempFile, err := ioutil.TempFile("", "ceps-*.db")
	if err != nil {
		return "", err
	}
	defer tempFile.Close()

	data, err := fs.ReadFile(cepsDB, "data/ceps.db")
	if err != nil {
		return "", err
	}

	_, err = tempFile.Write(data)
	if err != nil {
		return "", err
	}

	return tempFile.Name(), nil
}

// CloseDB fecha a conexão com o banco de dados
func CloseDB() {
	if db != nil {
		db.Close()
	}
}

// Endereco retorna um CEP aleatório do banco sem usar OFFSET
func Endereco() (*models.Address, error) {
	database := GetDB()
	if database == nil {
		return nil, fmt.Errorf("banco de dados não inicializado")
	}

	var minID, maxID int
	err := database.QueryRow("SELECT MIN(id), MAX(id) FROM ceps").Scan(&minID, &maxID)
	if err != nil {
		return nil, fmt.Errorf("erro ao obter range de IDs: %v", err)
	}

	randomID := minID
	if minID != maxID {
		randomID = rand.Intn(maxID-minID+1) + minID
	}

	query := `SELECT cep, logradouro, complemento, unidade, bairro, localidade, uf, estado, regiao, ibge, gia, ddd, siafi FROM ceps WHERE id = ?;`
	row := database.QueryRow(query, randomID)

	var address models.Address
	err = row.Scan(
		&address.Cep, &address.Logradouro, &address.Complemento, &address.Unidade,
		&address.Bairro, &address.Localidade, &address.Uf, &address.Estado, &address.Regiao,
		&address.Ibge, &address.Gia, &address.Ddd, &address.Siafi,
	)
	if err != nil {
		return nil, fmt.Errorf("erro ao buscar CEP aleatório: %v", err)
	}

	return &address, nil
}
