package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/davidklz/ds-app-backend/types"
	"github.com/gorilla/mux"
)

func (s *Server) handleDatentypRequests(w http.ResponseWriter, r *http.Request) error {
	switch r.Method {
	case "GET":
		return s.handleGetAllDatentypen(w, r)
	case "POST":
		return s.handleSaveDatentyp(w, r)
	default:
		return fmt.Errorf("method not allowed %s", r.Method)
	}
}

func (s *Server) handleDatentypRequestsById(w http.ResponseWriter, r *http.Request) error {
	switch r.Method {
	case "GET":
		return s.handleGetDatentyp(w, r)
	case "DELETE":
		return s.handleDeleteDatentyp(w, r)
	default:
		return fmt.Errorf("method not allowed %s", r.Method)
	}
}

func (s *Server) handleGetDatentyp(w http.ResponseWriter, r *http.Request) error {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		return err
	}

	v, err := s.store.GetDatentyp(id)
	if err != nil {
		return err
	}

	return WriteJson(w, http.StatusOK, v)
}

func (s *Server) handleControltypRequests(w http.ResponseWriter, r *http.Request) error {
	switch r.Method {
	case "GET":
		return s.handleGetAllControltypen(w, r)
	case "POST":
		return s.handleSaveControltyp(w, r)
	default:
		return fmt.Errorf("method not allowed %s", r.Method)
	}
}

func (s *Server) handleControltypRequestsById(w http.ResponseWriter, r *http.Request) error {
	switch r.Method {
	case "GET":
		return s.handleGetControltyp(w, r)
	case "DELETE":
		return s.handleDeleteControltyp(w, r)
	default:
		return fmt.Errorf("method not allowed %s", r.Method)
	}
}

func (s *Server) handleGetControltyp(w http.ResponseWriter, r *http.Request) error {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		return err
	}

	v, err := s.store.GetControltyp(id)
	if err != nil {
		return err
	}

	return WriteJson(w, http.StatusOK, v)
}

func (s *Server) handleGetAllDatentypen(w http.ResponseWriter, r *http.Request) error {
	dtypes, err := s.store.GetAllDatentypen()
	if err != nil {
		return err
	}

	return WriteJson(w, http.StatusOK, dtypes)
}

func (s *Server) handleSaveDatentyp(w http.ResponseWriter, r *http.Request) error {
	dtype := &types.Datentyp{}
	json.NewDecoder(r.Body).Decode(dtype)
	if err := s.store.SaveDatentyp(dtype); err != nil {
		return err
	}

	return WriteJson(w, http.StatusOK, types.DefaultResponse{Success: true, Message: "saved datentyp"})
}

func (s *Server) handleDeleteDatentyp(w http.ResponseWriter, r *http.Request) error {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		return err
	}

	err = s.store.DeleteDatentyp(id)
	if err != nil {
		return err
	}

	return WriteJson(w, http.StatusOK, types.DefaultResponse{Success: true, Message: "deleted datentyp"})
}

func (s *Server) handleGetAllControltypen(w http.ResponseWriter, r *http.Request) error {
	ctypes, err := s.store.GetAllControltypen()
	if err != nil {
		return err
	}

	return WriteJson(w, http.StatusOK, ctypes)
}

func (s *Server) handleSaveControltyp(w http.ResponseWriter, r *http.Request) error {
	ctype := &types.Controltyp{}
	json.NewDecoder(r.Body).Decode(ctype)
	if err := s.store.SaveControltyp(ctype); err != nil {
		return err
	}

	return WriteJson(w, http.StatusOK, types.DefaultResponse{Success: true, Message: "saved controltyp"})
}

func (s *Server) handleDeleteControltyp(w http.ResponseWriter, r *http.Request) error {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		return err
	}

	err = s.store.DeleteControltyp(id)
	if err != nil {
		return err
	}

	return WriteJson(w, http.StatusOK, types.DefaultResponse{Success: true, Message: "deleted controltyp"})
}
