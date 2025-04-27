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
func IsValidReceipt(w http.ResponseWriter, r *http.Request) (int, any) {
	receiptRequest := models.Receipt{}
	err := json.NewDecoder(r.Body).Decode(&receiptRequest)
	if err != nil {

	}
	return 400, models.NewBadRequest()
}

// Check if ID is valid, returns status and PointsRequest if ok; NotFound if not
func IsValidId(w http.ResponseWriter, r *http.Request) (int, any) {
	id := r.PathValue("id")
	for _, r := range models.GetReceipts() {
		if r.Id == id {
			return 200, models.NewPointsResponse(r.Points)
		}
	}
	return 404, models.NewNotFound()
}
