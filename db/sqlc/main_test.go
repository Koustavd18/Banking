package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/Koustavd18/Banking/utils"
	_ "github.com/lib/pq"
)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	config, err := utils.LoadConfig("../..")
	if err != nil {
		log.Fatal("error reading env", err)
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)

	if err != nil {
		log.Fatal("[Error] : Cannot connect to db", err)
	}

	testDB = conn
	testQueries = New(conn)

	os.Exit(m.Run())
}
