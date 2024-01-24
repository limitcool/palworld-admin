package handlers

const DefaultPageSize = 20

type Pagination struct {
	Page int `json:"page"`
	Size int `json:"size"`
}
