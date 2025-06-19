package server

import (
	"database/sql"

	"exemplar-api/internal/data"
)

type Handler struct {
	Q *data.Queries
}

func NewHandler(db *sql.DB) *Handler {
	return &Handler{
		Q: data.New(db),
	}
}
