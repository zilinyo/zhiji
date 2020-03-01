package models

import (
	db "zhiji/api/database"
)

type Users struct {
	ID           uint   `gorm:"primary_key"`
	Username     string `gorm:"column:username;type:varchar(20)" json:"username"`
	Password     string `gorm:"column:password;type:varchar(100)" json:"password"`
	Access_Token string `gorm:"column:access_token;type:varchar(255)" json:"access_token"`
}

func (Users) TableName() string {
	return "gb_api_user"
}

func List() *Users {
	var user Users
	db.Eloquent.First(&user)
	return &user
	//var list []Users
	//err := db.Eloquent.First(&Users{}).Error
	//
	//return &list, err
}

func UsersNumber() int {
	var num int

	db.Eloquent.Table("gb_api_user").Count(&num)

	return num
}

const SecretKey = "Hello World"

//func TestCreateToken(t *testing.T)  {
//	token, _ := CreateToken([]byte(SecretKey), "YDQ", 2222, true)
//	fmt.Println(token)
//}
