package models

import (
	"crypto/md5"
	"fmt"

	"time"
	db "zhiji/api/database"
	//"errors"
)

/******sql******
CREATE TABLE `customer_login` (
  `customer_id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '用户ID',
  `login_name` varchar(20) NOT NULL COMMENT '用户登录名',
  `password` char(32) NOT NULL COMMENT 'md5加密的密码',
  `user_stats` tinyint(4) NOT NULL DEFAULT '1' COMMENT '用户状态',
  `modified_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最后修改时间',
  PRIMARY KEY (`customer_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='用户登录表'
******sql******/
// CustomerLogin 用户登录表
type CustomerLogin struct {
	CustomerID   int       `gorm:"primary_key;column:customer_id;type:int(10) unsigned;not null"`                    // 用户ID
	LoginName    string    `form:"login_name" binding:"required" gorm:"column:login_name;type:varchar(50);not null"` // 用户登录名
	Password     string    `form:"password" binding:"required" gorm:"column:password;type:char(32);not null"`        // md5加密的密码
	UserStats    int8      `gorm:"column:user_stats;type:tinyint(4);not null"`                                       // 用户状态
	ModifiedTime time.Time `gorm:"column:modified_time;type:timestamp;not null"`                                     // 最后修改时间
}

const KEY string = "asdwdasdresdsxcghjtugbcvf"
const SecretKey = "hello world"

//获取token
func (CU *CustomerLogin) GetToken() string {

	token, _ := CreateToken([]byte(SecretKey), CU.Password, uint(CU.CustomerID), true)

	return token
}
func CheckPassword(pass string) string {

	data := []byte(pass + KEY)

	has := md5.Sum(data)

	md5str1 := fmt.Sprintf("%x", has) //将[]byte转成16进制

	return md5str1
}

//注册生成password
func (CU *CustomerLogin) Register() (err error) {

	encryption := CU

	data := []byte(encryption.Password + KEY)

	has := md5.Sum(data)

	md5str1 := fmt.Sprintf("%x", has) //将[]byte转成16进制

	CU.Password = md5str1

	CU.UserStats = 1

	CU.ModifiedTime = time.Now()

	err = db.Eloquent.Save(&CU).Error

	if err != nil {
		return err
	}

	return nil

}
