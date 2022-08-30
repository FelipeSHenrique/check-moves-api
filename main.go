package main

import (
	"checkmovies/api"
	"fmt"
	"log"
	"net/http"
)

func main() {

	a := api.API()

	fmt.Println("Escutando na porta 5000")
	log.Fatal(http.ListenAndServe(":5000", a))
}
