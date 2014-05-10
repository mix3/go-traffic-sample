package model

import (
	"fmt"

	"github.com/naoina/genmai"
	"github.com/pilu/traffic"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

func DBGet() *genmai.DB {
	driver := fmt.Sprintf("%s", traffic.GetVar("driver"))
	dsn := fmt.Sprintf("%s", traffic.GetVar("dsn"))
	var d genmai.Dialect
	switch driver {
	case "mysql":
		d = &genmai.MySQLDialect{}
	case "postgres":
		d = &genmai.PostgresDialect{}
	case "sqlite3":
		d = &genmai.SQLite3Dialect{}
	default:
		panic(fmt.Errorf("kocha: genmai: unsupported driver type: %v", driver))
	}
	db, err := genmai.New(d, dsn)
	if err != nil {
		panic(err)
	}
	db.CreateTableIfNotExists(&Todo{})
	return db
}
