package vilidate

type OrderForm struct {
	StoreCode     string `form:"store_code" binding:"required"`
	Recipients    string `form:"recipients" binding:"required"`
	Province      string `form:"province" binding:"required"`
	City          string `form:"city" binding:"required"`
	Area          string `form:"area" binding:"required"`
	Address       string `form:"address" binding:"required"`
	Postcodes     string `form:"postcodes" binding:"required"`
	Tel           string `form:"tel" binding:"required"`
	wh_trade_code string `form:"wh_trade_code" binding:"required"`
	Str           string `form:"str" binding:"required"`
}
