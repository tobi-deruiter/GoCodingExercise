package main

import (
	"GoCodingExercise/router"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/receipt/process", router.ProcessReceipt)
	http.HandleFunc("/receipt/{id}/points", router.GetPointsById)
	log.Fatal(http.ListenAndServe(":3000", nil))
}
