package server

import (
	"database/sql"
	"net/http"
	"strconv"
	"strings"

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

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch {
	case r.Method == http.MethodPost && r.URL.Path == "/notes":
		h.CreateNote(w, r)

	case r.Method == http.MethodGet && r.URL.Path == "/notes":
		h.ListNotes(w, r)

	case strings.HasPrefix(r.URL.Path, "/notes/"):
		idStr := strings.TrimPrefix(r.URL.Path, "/notes/")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "invalid note ID", http.StatusBadRequest)
			return
		}

		switch r.Method {
		case http.MethodGet:
			h.GetNote(w, r, int32(id))
		case http.MethodPut:
			h.UpdateNote(w, r, int32(id))
		case http.MethodDelete:
			h.DeleteNote(w, r, int32(id))
		default:
			http.NotFound(w, r)
		}

	default:
		http.NotFound(w, r)
	}
}
