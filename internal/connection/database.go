package connection

import (
	"database/sql"
	"fmt"
	"github.com/ZumaAkbarID/go-library-fiber/internal/config"
	_ "github.com/lib/pq"
	"log"
)

func GetDatabase(conf config.Database) *sql.DB {
	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable Timezone=%s",
		conf.Host,
		conf.Port,
		conf.User,
		conf.Name,
		conf.Tz)

	db, err := sql.Open("postgres", dsn)

	if err != nil {
		log.Fatal("Failed to open connection to database", err.Error())
	}

	err = db.Ping()
	
	if err != nil {
		log.Fatal("Failed to ping database", err.Error())
	}

	return db
}
