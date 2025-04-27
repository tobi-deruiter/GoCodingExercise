package models

// ENDPOINTS

// /receipts/{id}/points
type PointsResponse struct {
	points int
}

func NewPointsResponse(points int) PointsResponse {
	return PointsResponse{points: points}
}

// ERROR RESPONSES

// Bad Request
type BadRequest struct {
	description string
}

func NewBadRequest() BadRequest {
	return BadRequest{description: "The receipt is invalid."}
}

// Not Found
type NotFound struct {
	description string
}

func NewNotFound() NotFound {
	return NotFound{description: "No receipt found for that ID."}
}
