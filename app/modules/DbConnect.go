package modules

import (
	"database/sql"
	"fmt"
	"log"

	//mysql for driver of mysql
	_ "github.com/go-sql-driver/mysql"
)

//DbConnect is a class of library connection to mysql
type DbConnect struct{}

var db *sql.DB

//Init for configure database connection
func (d DbConnect) Init() {
	config := new(Configuration).GetConfig()

	var err error
	mysqlInfo := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&multiStatements=true",
		config.GetString("DB_USER"), config.GetString("DB_PASS"), config.GetString("DB_HOST"), config.GetString("DB_PORT"), config.GetString("DATABASE"))

	db, err = sql.Open("mysql", mysqlInfo)
	if err != nil {
		log.Fatal(err)
	}
}

//GetDB is for get current database after init
func (d DbConnect) GetDB() *sql.DB {
	return db
}
