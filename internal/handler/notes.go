package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/aandermai/notes-server/internal/model"
	"github.com/aandermai/notes-server/internal/storage"
)

// Хендлер для отображения всех заметок
func showNotesHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(storage.Notes)
	if err != nil {
		http.Error(w, "Ошибка при отправке JSON", http.StatusInternalServerError)
		return
	}
}

// Хендлер для создания заметки
func createNoteHandler(w http.ResponseWriter, r *http.Request) {
	var req model.CreateNoteRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	note := model.Note{
		ID:      len(storage.Notes) + 1,
		Title:   req.Title,
		Content: req.Content,
	}

	storage.Notes = append(storage.Notes, note)

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(note)
	if err != nil {
		http.Error(w, "Ошибка во время отправки JSON", http.StatusInternalServerError)
		return
	}
}

func searchNoteHandler(w http.ResponseWriter, r *http.Request) {
	urlSlice := strings.Split(r.URL.Path, "/")
	noteNumber, err := strconv.Atoi(urlSlice[len(urlSlice)-1])

	if err != nil {
		http.Error(w, "Invalid Note ID", http.StatusBadRequest)
		return
	}

	for _, note := range storage.Notes {
		if note.ID == noteNumber {
			w.Header().Set("Content-Type", "application/json")
			err = json.NewEncoder(w).Encode(note)
			if err != nil {
				http.Error(w, "Ошибка отправки данных", http.StatusInternalServerError)
				return
			}
			return
		}
	}

	http.Error(w, "Note not found", http.StatusNotFound)
}

func NotesHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		if r.URL.Path == "/notes" {
			showNotesHandler(w, r)
		} else {
			searchNoteHandler(w, r)
		}

	case http.MethodPost:
		createNoteHandler(w, r)

	default:
		http.Error(w, "Nethod not allowed", http.StatusMethodNotAllowed)
		return
	}
}
