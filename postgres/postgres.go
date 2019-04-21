package postgres

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

type BigSerial int64

var singletonDB *sqlx.DB
var once sync.Once

func SharedDB() *sqlx.DB {
	once.Do(func() {
		singletonDB = newDB()
	})
	return singletonDB
}

func postgresURL() (url string) {
	return os.Getenv("DATABASE_URL")
}

func newDB() *sqlx.DB {
	db, err := sqlx.Connect("postgres", postgresURL())
	if err != nil {
		logrus.Errorf("Postgres : Error Connecting to DB => %v", err)
		panic(err)
	}
	return db
}

func TableNames() []string {
	return []string{
		"users",
	}
}

func CreateTables() error {
	logrus.Info("Postgres : Initializing Schema")

	fileContents := []string{}
	schemaDirectory := "/postgres/schema"
	for _, tableName := range TableNames() {
		fullFileName := filepath.Join(schemaDirectory, tableName+".sql")
		fileContent, err := ioutil.ReadFile(fullFileName)
		if err != nil {
			logrus.Errorf("Postgres : Error Initializing Schema =>\n%v", err)
			return err
		}
		fileContents = append(fileContents, string(fileContent))
	}

	for _, fileContent := range fileContents {
		_, err := SharedDB().Exec(fileContent)
		if err != nil {
			logrus.Errorf("Postgres: Could not create table: %v =>\n%v", fileContent, err)
			return err
		}
		logrus.Debugf("Postgres: Successfully created table =>\n%v", fileContent)
	}

	return nil
}
