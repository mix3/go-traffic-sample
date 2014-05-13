package model

import "github.com/naoina/genmai"

type Todo struct {
	Id        int64  `db:"pk" json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func TodoList() ([]Todo, error) {
	db := GetDB()
	var results []Todo
	err := db.Select(&results, db.OrderBy("id", genmai.ASC))
	return results, err
}

func TodoCreate(title string) error {
	db := GetDB()
	_, err := db.Insert(&Todo{
		Title:     title,
		Completed: false,
	})
	return err
}

func TodoSwitch(id int64) error {
	db := GetDB()
	var results []Todo
	if err := db.Select(&results, db.Where("id", "=", id)); err != nil {
		return err
	}
	obj := results[0]
	obj.Completed = !obj.Completed
	_, err := db.Update(&obj)
	return err
}

func TodoDelete(id int64) error {
	db := GetDB()
	obj := Todo{Id: id}
	_, err := db.Delete(&obj)
	return err
}

func TodoDeleteAll() error {
	db := GetDB()
	var targets []Todo
	if err := db.Select(&targets, db.Where("completed", "=", 1)); err != nil {
		return err
	}
	if _, err := db.Delete(&targets); err != nil {
		return err
	}
	return nil
}
