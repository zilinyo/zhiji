package models

import (
	db "zhiji/api/database"
)

type Visitors struct {
	ID  int `json:"id"`
	Num int `json:"num"` //总的点击量
}

func VisitorsNumber() Visitors {
	var num Visitors

	db.Eloquent.Table("visitors").Select("sum(clicks) as num").Scan(&num)

	return num
}
