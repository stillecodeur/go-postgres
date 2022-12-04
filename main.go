package main

import (
	"fmt"
	"go-postgres/router"
	"log"
	"net/http"
)

func main() {
	r := router.Router()
	fmt.Println("Starting server at port 3000")
	log.Fatal(http.ListenAndServe("localhost:3000", r))
}
