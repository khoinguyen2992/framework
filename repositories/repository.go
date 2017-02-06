package repositories

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

var (
	connectionStringTemplate = "host=%s port=%s sslmode=%s user=%s password='%s' dbname=%s"

	db *gorm.DB
)

func InitDatabase() (err error) {
	if db == nil {
		db, err = gorm.Open("postgres", getConnectionString())
	}

	return err
}

func CloseDatabase() error {
	return db.Close()
}

func getConnectionString() string {
	return fmt.Sprintf(connectionStringTemplate,
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_SSL_MODE"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"))
}
