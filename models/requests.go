package models

// /receipts/{id}/points
type PointsRequest struct {
	id string
}

func NewPointsRequest(id string) PointsRequest {
	return PointsRequest{id: id}
}
