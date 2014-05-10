package model

import "github.com/naoina/genmai"

func TodoList(db *genmai.DB) ([]Todo, error) {
	var results []Todo
	err := db.Select(&results)
	return results, err
}

func TodoCreate(db *genmai.DB, title string) error {
	_, err := db.Insert(&Todo{
		Title:     title,
		Completed: false,
	})
	return err
}

func TodoSwitch(db *genmai.DB, id int64) error {
	var results []Todo
	if err := db.Select(&results, db.Where("id", "=", id)); err != nil {
		return err
	}
	obj := results[0]
	obj.Completed = !obj.Completed
	_, err := db.Update(&obj)
	return err
}

func TodoDelete(db *genmai.DB, id int64) error {
	obj := Todo{Id: id}
	_, err := db.Delete(&obj)
	return err
}
