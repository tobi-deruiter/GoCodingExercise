package middleware

import (
	"GoCodingExercise/models"
	"encoding/json"
	"net/http"
)

/*
	PUBLIC
*/

// Check for bad requests, returns status and Receipt if ok; BadRequest if not
func IsValidReceipt(w http.ResponseWriter, r *http.Request) (int, interface{}) {
	receipt := models.Receipt{}
	err := json.NewDecoder(r.Body).Decode(&receipt)
	if err != nil {
		return 400, models.NewBadRequest()
	}
	return 200, receipt
}

// Check if ID is valid, returns status and PointsRequest if ok; NotFound if not
func IsValidId(w http.ResponseWriter, r *http.Request) (int, interface{}) {
	id := r.PathValue("id")
	for i, r := range models.GetReceipts() {
		if r.Id() == id {
			return 200, models.NewPointsRequest(r.Id(), i)
		}
	}
	return 404, models.NewNotFound()
}
