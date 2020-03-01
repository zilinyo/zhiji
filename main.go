package main

import (
	db "zhiji/api/database"
	router2 "zhiji/api/router"
)

func main() {
	defer db.Eloquent.Close()

	router := router2.InitRouter()

	router.Run(":8081")
}
