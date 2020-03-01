package models

import (
	db "zhiji/api/database"
)

type Articles struct {
	ID int `json:"id"`
}

func ArticlesNumber() int {
	var num int

	db.Eloquent.Table("articles").Count(&num)

	return num
}
