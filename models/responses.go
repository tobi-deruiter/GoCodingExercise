package models

// ENDPOINTS

// /receipts/process
type ProcessResponse struct {
	id string
}

func NewProcessResponse(id string) ProcessResponse {
	return ProcessResponse{id: id}
}

// /receipts/{id}/points
type PointsResponse struct {
	points int
}

func NewPointsResponse(points int) PointsResponse {
	return PointsResponse{points: points}
}

// ERROR RESPONSES

type Error struct {
	description string
}

func NewBadRequest() Error {
	return Error{description: "The receipt is invalid."}
}

func NewNotFound() Error {
	return Error{description: "No receipt found for that ID."}
}

func NewError(err string) Error {
	return Error{description: err}
}
