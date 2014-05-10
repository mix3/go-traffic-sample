package model

import (
	"fmt"
	"os"
	"reflect"
	"testing"

	_ "github.com/mattn/go-sqlite3"
	"github.com/naoina/genmai"
)

func testDB(dsn ...string) (*genmai.DB, error) {
	switch os.Getenv("DB") {
	case "mysql":
		return genmai.New(&genmai.MySQLDialect{}, "travis@/go_traffic_sample_test")
	case "postgres":
		return genmai.New(&genmai.PostgresDialect{}, "user=postgres dbname=go_traffic_sample_test sslmode=disable")
	default:
		var DSN string
		switch len(dsn) {
		case 0:
			DSN = ":memory:"
		case 1:
			DSN = dsn[0]
		default:
			panic(fmt.Errorf("too many arguments"))
		}
		return genmai.New(&genmai.SQLite3Dialect{}, DSN)
	}
}

func newTestDB(t *testing.T) *genmai.DB {
	db, err := testDB()
	if err != nil {
		t.Fatal(err)
	}
	err = db.CreateTable(&Todo{})
	if err != nil {
		t.Fatal(err)
	}
	return db
}

func TestCRUD(t *testing.T) {
	db := newTestDB(t)
	defer db.Close()

	// empty
	actual, err := TodoList(db)
	if err != nil {
		t.Fatal(err)
	}
	expected := []Todo{}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Expect %q but %q", expected, actual)
	}

	// create
	for _, f := range []func() error{
		func() error { return TodoCreate(db, "TODO-1") },
		func() error { return TodoCreate(db, "TODO-2") },
		func() error { return TodoCreate(db, "TODO-3") },
	} {
		if err := f(); err != nil {
			t.Fatal(err)
		}
	}
	actual, err = TodoList(db)
	if err != nil {
		t.Fatal(err)
	}
	expected = []Todo{
		{Id: 1, Title: "TODO-1", Completed: false},
		{Id: 2, Title: "TODO-2", Completed: false},
		{Id: 3, Title: "TODO-3", Completed: false},
	}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Expect %q but %q", expected, actual)
	}

	// update
	if err := TodoSwitch(db, 2); err != nil {
		t.Fatal(err)
	}
	actual, err = TodoList(db)
	if err != nil {
		t.Fatal(err)
	}
	expected = []Todo{
		{Id: 1, Title: "TODO-1", Completed: false},
		{Id: 2, Title: "TODO-2", Completed: true},
		{Id: 3, Title: "TODO-3", Completed: false},
	}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Expect %q but %q", expected, actual)
	}

	// delete
	if err := TodoDelete(db, 1); err != nil {
		t.Fatal(err)
	}
	actual, err = TodoList(db)
	if err != nil {
		t.Fatal(err)
	}
	expected = []Todo{
		{Id: 2, Title: "TODO-2", Completed: true},
		{Id: 3, Title: "TODO-3", Completed: false},
	}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Expect %q but %q", expected, actual)
	}
}
