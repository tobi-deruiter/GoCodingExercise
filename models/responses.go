package models

// ENDPOINTS

// /receipts/process
type ProcessResponse struct {
	Id string `json:"id"`
}

func NewProcessResponse(id string) ProcessResponse {
	return ProcessResponse{Id: id}
}

// /receipts/{id}/points
type PointsResponse struct {
	Points int `json:"points"`
}

func NewPointsResponse(points int) PointsResponse {
	return PointsResponse{Points: points}
}

// ERROR RESPONSES

type Error struct {
	Description string `json:"description"`
}

func NewBadRequest() Error {
	return Error{Description: "The receipt is invalid."}
}

func NewNotFound() Error {
	return Error{Description: "No receipt found for that ID."}
}

func NewError(err string) Error {
	return Error{Description: err}
}
