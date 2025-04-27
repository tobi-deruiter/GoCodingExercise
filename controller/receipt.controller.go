package controller

import (
	"GoCodingExercise/models"
	"encoding/json"
	"net/http"
)

// POST REQUESTS

var receipts []models.Receipt

func processReceipt(w http.ResponseWriter, r *http.Request) {
	// add to receipts
}

// GET REQUESTS

func getPoints(w http.ResponseWriter, r *http.Request, status int) {
	pointsRequest := models.PointsRequest{}

	err := json.NewDecoder(r.Body).Decode(&pointsRequest)
	if err != nil {

	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	// jsonData, err := json.Marshal()
}
