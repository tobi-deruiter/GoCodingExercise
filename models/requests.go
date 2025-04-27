package models

// /receipts/{id}/points
type PointsRequest struct {
	id  string
	idx int
}

func NewPointsRequest(id string, idx int) PointsRequest {
	return PointsRequest{id: id, idx: idx}
}

func (pReq PointsRequest) Idx() int {
	return pReq.idx
}
