package repository

import (
	"fmt"
	//. "zhiji/api/models"
	"reflect"
	. "zhiji/api/vidate"
)

func ProductRes(p *ProducVidate) bool {
	value := reflect.ValueOf(p)
	for i := 0; i < value.NumField(); i++ {
		fmt.Printf("Field %d: %v\n", i, value.Field(i))
	}

	//var 	product =new(ProductInfo)
	//		for k,v:=range product{
	//			for w,l:=range p{
	//				if k=w{
	//					v=l
	//				}
	//			}
	//}

	return true
}
