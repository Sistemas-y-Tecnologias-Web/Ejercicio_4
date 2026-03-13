package videogames

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"strings"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) Register(mux *http.ServeMux) {
	mux.HandleFunc("/api/videogames", h.router)
	mux.HandleFunc("/api/videogames/", h.router)
	mux.HandleFunc("/api", h.router)
}

func (h *Handler) router(w http.ResponseWriter, r *http.Request) {
	hasID := strings.Count(r.URL.Path, "/") == 3
	health := strings.Count(r.URL.Path, "/") == 1

	if health {
		switch r.Method {
		case http.MethodGet:
			h.health(w)
		default:
			writeError(w, http.StatusNotFound, "method not allowed!")
		}
		return
	}

	if !hasID {
		switch r.Method {
		case http.MethodGet:
			h.list(w, r)
		case http.MethodPost:
			h.create(w, r)
		default:
			writeError(w, http.StatusMethodNotAllowed, "method not allowed!")
		}
		return
	}

	switch r.Method {
	case http.MethodGet:
		h.getOne(w, r)
	case http.MethodPut:
		h.update(w, r)
	case http.MethodDelete:
		h.delete(w, r)
	default:
		writeError(w, http.StatusMethodNotAllowed, "method not allowed!")
	}
}

func (h *Handler) health(w http.ResponseWriter) {
	writeJSON(w, http.StatusOK, map[string]string{"message": "I'm Alive 😁"})
}

func (h *Handler) list(w http.ResponseWriter, r *http.Request) {
	videogames, err := h.service.GetAll(r.Context())
	if err != nil {
		writeError(w, http.StatusInternalServerError, "could not fetch videogames!")
		return
	}
	if videogames == nil {
		videogames = []Videogame{}
	}
	writeJSON(w, http.StatusOK, videogames)
}

func (h *Handler) getOne(w http.ResponseWriter, r *http.Request) {
	id, ok := extractID(r.URL.Path)
	if !ok {
		writeError(w, http.StatusBadRequest, "invalid id!")
		return
	}

	videogame, err := h.service.GetByID(r.Context(), id)
	if err != nil {
		if errors.Is(err, ErrNotFound) {
			writeError(w, http.StatusNotFound, err.Error())
		} else {
			writeError(w, http.StatusInternalServerError, "could not fetch videogame!")
		}
		return
	}

	writeJSON(w, http.StatusOK, videogame)
}

func (h *Handler) create(w http.ResponseWriter, r *http.Request) {
	var req CreateVideogame
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid body")
		return
	}

	videogame, err := h.service.Create(r.Context(), req)
	if err != nil {
		if errors.Is(err, ErrNameRequired) || errors.Is(err, ErrCategoryRequired) || errors.Is(err, ErrSizeRequired) {
			writeError(w, http.StatusBadRequest, err.Error())
		} else {
			writeError(w, http.StatusInternalServerError, "could not create videogame")
		}
		return
	}

	writeJSON(w, http.StatusCreated, videogame)
}

func (h *Handler) update(w http.ResponseWriter, r *http.Request) {
	var req UpdateVideogame
	id, ok := extractID(r.URL.Path)

	if !ok {
		writeError(w, http.StatusBadRequest, "invalid id")
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid body")
		return
	}

	videogame, err := h.service.Update(r.Context(), id, req)

	if err != nil {
		if errors.Is(err, ErrNotFound) {
			writeError(w, http.StatusNotFound, err.Error())
		} else {
			writeError(w, http.StatusInternalServerError, "could not update videogame")
		}
		return
	}

	writeJSON(w, http.StatusOK, videogame)
}

func (h *Handler) delete(w http.ResponseWriter, r *http.Request) {
	id, ok := extractID(r.URL.Path)
	if !ok {
		writeError(w, http.StatusNotFound, "invalid id!")
		return
	}

	if err := h.service.Delete(r.Context(), id); err != nil {
		if errors.Is(err, ErrNotFound) {
			writeError(w, http.StatusNotFound, err.Error())
		} else {
			writeError(w, http.StatusInternalServerError, "could not delete videogame")
		}
		return
	}

	writeJSON(w, http.StatusOK, map[string]string{"message": "videogame deleted"})
}

func writeError(w http.ResponseWriter, status int, msg string) {
	writeJSON(w, status, map[string]string{"error": msg})
}

func writeJSON(w http.ResponseWriter, status int, payload any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(payload)
}

func extractID(path string) (int, bool) {
	parts := strings.Split(strings.Trim(path, "/"), "/")
	if len(parts) < 3 {
		return 0, false
	}
	id, err := strconv.Atoi(parts[2])
	return id, err == nil
}
