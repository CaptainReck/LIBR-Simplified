package main

import (
	"fmt"
	"libr-simplified/db"
	"libr-simplified/router"
	"net/http"
)

func main() {
	db.InitDB()
	fmt.Println("LIBR-Simplified")
	r := router.Router()
	fmt.Println("Server starting...")
	http.ListenAndServe(":4000", r)
}
