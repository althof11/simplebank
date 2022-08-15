package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/althof3/simplebank/util"
	_ "github.com/lib/pq"
)

var TestQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
    var err error
    config, err := util.LoadConfig("../..")
    if err != nil {
        log.Fatal(err)
    }
    
    testDB, err = sql.Open(config.DBDriver, config.DBSource)
    if err != nil {
        log.Fatal("cannot connect to database")
    }

    TestQueries = New(testDB)

    os.Exit(m.Run())
}
