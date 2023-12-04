package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/davidklz/ds-app-backend/types"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func (s *Server) handleVariableRequests(w http.ResponseWriter, r *http.Request) error {
	switch r.Method {
	case "GET":
		return s.handleGetAllVariablen(w, r)
	case "POST":
		return s.handleSaveVariable(w, r)
	default:
		return fmt.Errorf("method not allowed %s", r.Method)
	}
}

func (s *Server) handleVariableRequestsById(w http.ResponseWriter, r *http.Request) error {
	switch r.Method {
	case "GET":
		return s.handleGetVariable(w, r)
	case "DELETE":
		return s.handleDeleteVariable(w, r)
	default:
		return fmt.Errorf("method not allowed %s", r.Method)
	}
}

func (s *Server) handleGetVariable(w http.ResponseWriter, r *http.Request) error {
	id, err := uuid.Parse(mux.Vars(r)["id"])
	if err != nil {
		return err
	}

	v, err := s.store.GetVariable(id)
	if err != nil {
		return err
	}

	return WriteJson(w, http.StatusOK, v)
}

func (s *Server) handleGetAllVariablen(w http.ResponseWriter, r *http.Request) error {
	vars, err := s.store.GetAllVariablen()
	if err != nil {
		return err
	}

	return WriteJson(w, http.StatusOK, vars)
}

func (s *Server) handleSaveVariable(w http.ResponseWriter, r *http.Request) error {
	variable := &types.Variable{}
	json.NewDecoder(r.Body).Decode(variable)
	if err := s.store.SaveVariable(variable); err != nil {
		return err
	}

	return WriteJson(w, http.StatusOK, map[string]string{"saved": variable.Name})
}

func (s *Server) handleDeleteVariable(w http.ResponseWriter, r *http.Request) error {
	id, err := uuid.Parse(mux.Vars(r)["id"])
	if err != nil {
		return err
	}

	err = s.store.DeleteVariable(id)
	if err != nil {
		return err
	}

	return WriteJson(w, http.StatusOK, map[string]uuid.UUID{"deleted": id})
}
