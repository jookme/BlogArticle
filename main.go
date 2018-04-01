package main

import (
	db "BlogArticle/database"
)

func main() {
	defer db.SqlDb.Close()
	router := initRouter()
	router.Run(":8000")
}
