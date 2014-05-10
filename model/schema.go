package model

type Todo struct {
	Id        int64  `db:"pk" json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}
