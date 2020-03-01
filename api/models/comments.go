package models

import (
	db "zhiji/api/database"
)

type Comments struct {
	ID int `json:"id"`
}

func CommentsNumber() int {
	var num int

	db.Eloquent.Table("comments").Count(&num)

	return num
}
