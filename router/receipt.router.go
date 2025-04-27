package router

import (
	"GoCodingExercise/controller"
	"GoCodingExercise/middleware"
	"GoCodingExercise/models"
	"encoding/json"
	"net/http"
)

// Send Json Response Data
func sendJson(w http.ResponseWriter, data any, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	jsonData, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	_, err = w.Write(jsonData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

/*
	ENDPOINTS
*/

// /receipts/process
func ProcessReceipt(w http.ResponseWriter, r *http.Request) {
	status, res := middleware.IsValidReceipt(w, r)
	if receipt, ok := res.(models.Receipt); ok && status == 200 {
		sendJson(w, controller.ProcessReceipt(receipt), status)
		return
	}
	sendJson(w, res, status)
}

// /receipts/{id}/points
func GetPointsById(w http.ResponseWriter, r *http.Request) {
	status, res := middleware.IsValidId(w, r)
	if pReq, ok := res.(models.PointsRequest); ok && status == 200 {
		sendJson(w, controller.GetPoints(pReq), status)
		return
	}
	sendJson(w, res, status)
}
