package action

import (
	"github.com/jmoiron/sqlx"
	"os"
)

var db *sqlx.DB

// init
func Init() (err error) {
	datasource := os.Getenv("DATASOURCE")
	db, err = sqlx.Connect("mysql", datasource)
	return
}

func Final() {
	db.Close()
}
