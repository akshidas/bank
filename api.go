package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func writeJson(w http.ResponseWriter, status int, v any) error {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(v)
}

type ApiError struct {
	Error string
}

type apiFunc func(http.ResponseWriter, *http.Request) error

func makeHttpHandlerFunc(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			//handle the error
			writeJson(w, http.StatusBadRequest, ApiError{Error: err.Error()})
		}
	}
}

type ApiServer struct {
	listeningPort string
}

func newApiServer(listerAddress string) *ApiServer {
	return &ApiServer{
		listeningPort: listerAddress,
	}
}

func (s *ApiServer) Run() {
	router := mux.NewRouter()

	router.HandleFunc("/account", makeHttpHandlerFunc(s.handleAccount))

	log.Println("JSON API server running on port:", 3000)
	err := http.ListenAndServe(":3000", router)
	log.Println(err)
}

func (s *ApiServer) handleAccount(w http.ResponseWriter, r *http.Request) error {

	if r.Method == "GET" {
		return s.handleGetAccount(w, r)
	}
	return fmt.Errorf("Method Not Allowed: %s", r.Method)
}

func (s *ApiServer) handleGetAccount(w http.ResponseWriter, r *http.Request) error {
	account := newAccount("John", "Doe")
	return writeJson(w, http.StatusOK, account)
}

// func (s *ApiServer) handleCreateAccount(w http.ResponseWriter, r *http.Request) error {
//
// 	return nil
// }
//
// func (s *ApiServer) handleDeleteAccount(w http.ResponseWriter, r *http.Request) error {
//
// 	return nil
// }
