package connection

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/shellrean/golang-base-project-cean-directory/internal/config"
	"log"
)

func GetDatabase(conf config.Database) *sql.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=%s",
		conf.Host,
		conf.User,
		conf.Pass,
		conf.Name,
		conf.Port,
		conf.Tz,
	)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal("failed open connection to db: ", err.Error())
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("failed open connection to db: ", err.Error())
	}

	return db
}
