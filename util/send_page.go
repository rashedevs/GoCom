package util

import (
	"net/http"
)

type Pagination struct {
	Page       int64 `json:"page"`
	Limit      int64 `json:"limit"`
	TotalItems int64 `json:"totalItems"`
	TotalPages int64 `json:"totalPages"`
}

type PaginatedData struct {
	Data       any        `json:"data"`
	Pagination Pagination `json:"pagination"`
}

func SendPage(w http.ResponseWriter, data any, page, limit, count int64) {
	paginatedData := PaginatedData{
		Data: data,
		Pagination: Pagination{
			Page:       page,
			Limit:      limit,
			TotalItems: count,
			TotalPages: count / limit,
		},
	}
	SendData(w, http.StatusOK, paginatedData)
}
