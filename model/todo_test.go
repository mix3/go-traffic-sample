package model

import (
	"os"
	"reflect"
	"testing"

	_ "github.com/mattn/go-sqlite3"
	"github.com/naoina/genmai"
)

func testDB() (*genmai.DB, error) {
	switch os.Getenv("DB") {
	case "mysql":
		return genmai.New(&genmai.MySQLDialect{}, os.Getenv("DSN"))
	case "postgres":
		return genmai.New(&genmai.PostgresDialect{}, os.Getenv("DSN"))
	default:
		return genmai.New(&genmai.SQLite3Dialect{}, ":memory:")
	}
}

func TestCRUD(t *testing.T) {
	// mocking initDB
	backDB := initDB
	initDB = func() *genmai.DB {
		db, err := testDB()
		if err != nil {
			t.Fatal(err)
		}
		if _, err := db.DB().Exec("DROP TABLE IF EXISTS todo"); err != nil {
			t.Fatal(err)
		}
		err = db.CreateTable(&Todo{})
		if err != nil {
			t.Fatal(err)
		}
		//db.SetLogOutput(os.Stdout)
		return db
	}
	defer func() {
		db = nil
		initDB = backDB
	}()

	// empty
	actual, err := TodoList()
	if err != nil {
		t.Fatal(err)
	}
	expected := []Todo{}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Expect %q but %q", expected, actual)
	}

	// create
	for _, f := range []func() error{
		func() error { return TodoCreate("TODO-1") },
		func() error { return TodoCreate("TODO-2") },
		func() error { return TodoCreate("TODO-3") },
	} {
		if err := f(); err != nil {
			t.Fatal(err)
		}
	}
	actual, err = TodoList()
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
	if err := TodoSwitch(2); err != nil {
		t.Fatal(err)
	}
	actual, err = TodoList()
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
	if err := TodoDelete(1); err != nil {
		t.Fatal(err)
	}
	actual, err = TodoList()
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

	// delete all
	for _, f := range []func() error{
		func() error { return TodoCreate("TODO-4") },
		func() error { return TodoCreate("TODO-5") },
		func() error { return TodoCreate("TODO-6") },
		func() error { return TodoSwitch(4) },
		func() error { return TodoSwitch(5) },
	} {
		if err := f(); err != nil {
			t.Fatal(err)
		}
	}
	actual, err = TodoList()
	if err != nil {
		t.Fatal(err)
	}
	expected = []Todo{
		{Id: 2, Title: "TODO-2", Completed: true},
		{Id: 3, Title: "TODO-3", Completed: false},
		{Id: 4, Title: "TODO-4", Completed: true},
		{Id: 5, Title: "TODO-5", Completed: true},
		{Id: 6, Title: "TODO-6", Completed: false},
	}
	if err := TodoDeleteAll(); err != nil {
		panic(err)
		t.Fatal(err)
	}
	actual, err = TodoList()
	if err != nil {
		t.Fatal(err)
	}
	expected = []Todo{
		{Id: 3, Title: "TODO-3", Completed: false},
		{Id: 6, Title: "TODO-6", Completed: false},
	}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Expect %q but %q", expected, actual)
	}
}
