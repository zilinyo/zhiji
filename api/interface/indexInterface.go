package _interface

import (
	. "zhiji/api/models"
)

/*
@Time : 2020/4/15 9:04 AM
@Author : apple
@File : indexInterface
@Software: GoLand
*/
//首页接口
type MiniIndex interface {
	Banner() []ProductInfo //banner图

	Channel() []ProductCategory //品牌分类

	NewGoodsList() []ProductInfo //新品推荐
}
