package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/davidklz/ds-app-backend/types"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func (s *Server) handleFormularRequests(w http.ResponseWriter, r *http.Request) error {
	switch r.Method {
	case "GET":
		return s.handleGetAllFormulare(w, r)
	case "POST":
		return s.handleSaveFormular(w, r)
	default:
		return fmt.Errorf("method not allowed %s", r.Method)
	}
}

func (s *Server) handleFormularRequestsById(w http.ResponseWriter, r *http.Request) error {
	switch r.Method {
	case "GET":
		return s.handleGetFormular(w, r)
	case "DELETE":
		return s.handleDeleteFormular(w, r)
	default:
		return fmt.Errorf("method not allowed %s", r.Method)
	}
}

func (s *Server) handleGetFormular(w http.ResponseWriter, r *http.Request) error {
	id, err := uuid.Parse(mux.Vars(r)["id"])
	if err != nil {
		return err
	}

	form, err := s.store.GetFormular(id)
	if err != nil {
		return err
	}

	return WriteJson(w, http.StatusOK, form)
}

func (s *Server) handleGetAllFormulare(w http.ResponseWriter, r *http.Request) error {
	forms, err := s.store.GetAllFormulare()
	if err != nil {
		return err
	}

	return WriteJson(w, http.StatusOK, forms)
}

func (s *Server) handleSaveFormular(w http.ResponseWriter, r *http.Request) error {
	form := &types.Formular{}
	json.NewDecoder(r.Body).Decode(form)
	if err := s.store.SaveFormular(form); err != nil {
		return err
	}

	return WriteJson(w, http.StatusOK, map[string]string{"saved": form.Name})
}

func (s *Server) handleDeleteFormular(w http.ResponseWriter, r *http.Request) error {
	id, err := uuid.Parse(mux.Vars(r)["id"])
	if err != nil {
		return err
	}

	err = s.store.DeleteFormular(id)
	if err != nil {
		return err
	}

	return WriteJson(w, http.StatusOK, map[string]uuid.UUID{"deleted": id})
}
