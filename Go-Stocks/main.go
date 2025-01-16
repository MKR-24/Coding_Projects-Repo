package main

import (
	"go-postgres/router"
	"fmt"
	"log"
	"net/http"
)

func main(){
	r:=router.Router()
	fmt.Println("Starting Server...")

	log.Fatal(http.ListenAndServe(":8080",r))
}