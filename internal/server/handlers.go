package server

import (
	"database/sql"
	"encoding/json"
	"exemplar-api/internal/data"
	"net/http"
)

func (h *Handler) CreateNote(w http.ResponseWriter, r *http.Request) {

	var req CreateRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	note, err := h.Q.CreateNote(r.Context(), data.CreateNoteParams{
		Title:   req.Title,
		Content: sql.NullString{String: req.Content, Valid: req.Content != ""},
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(note)
}

func (h *Handler) ListNotes(w http.ResponseWriter, r *http.Request) {
	notes, err := h.Q.ListNotes(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(notes)
}

func (h *Handler) GetNote(w http.ResponseWriter, r *http.Request, id int) {

	note, err := h.Q.GetNote(r.Context(), int32(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(note)
}

func (h *Handler) UpdateNote(w http.ResponseWriter, r *http.Request, id int) {

	var req UpdateRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	note, err := h.Q.UpdateNote(r.Context(), data.UpdateNoteParams{
		ID:      int32(id),
		Title:   req.Title,
		Content: sql.NullString{String: req.Content, Valid: req.Content != ""},
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(note)
}

func (h *Handler) DeleteNote(w http.ResponseWriter, r *http.Request, id int) {
	err := h.Q.DeleteNote(r.Context(), int32(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
