package main

import (
	"GoCodingExercise/router"
	"fmt"
	"log"
	"net/http"
)

var PORT int = 3000

func main() {
	http.HandleFunc("/receipt/process", router.ProcessReceipt)
	http.HandleFunc("/receipt/{id}/points", router.GetPointsById)
	fmt.Printf("Starting server on port %d\n", PORT)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", PORT), nil))
}
