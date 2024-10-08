package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Server struct {
	listenPort string
}

// writting Json
func WriteJson (w http.ResponseWriter, status int, value any) error {
	w.WriteHeader(status)
	w.Header().Set("content-type", "application/json")

	return json.NewEncoder(w).Encode(value)
}

// type for the apiHandler functions I've made
type apiFunc func (http.ResponseWriter, *http.Request) error

type apiError struct {
	Error string 
}

// decorating apiFunction into httpHandler, to remove the retunr type error
func makeHTTPHandleFunc (apiFunction apiFunc) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request) {
		if err := apiFunction(w, r); err != nil {
			WriteJson(w, http.StatusBadRequest, apiError{Error: err.Error()})
		}
	}
}

// constructor
func NewServer (listenAddr string) (*Server) {
	return &Server{listenPort: listenAddr}
}

// router
func (s *Server) Run () {
	router := mux.NewRouter()

	router.HandleFunc("/account", makeHTTPHandleFunc(s.handleAccount))

	router.HandleFunc("/account/{id}", makeHTTPHandleFunc(s.handleAccount))

	log.Printf("Server Listening on port: %s", s.listenPort)
	http.ListenAndServe(s.listenPort, router)
}

// to handle the type of request, mux doesn't allow POST, PUT, DELETE, PATCH........
func (s *Server) handleAccount (w http.ResponseWriter, r *http.Request) error {
	method := r.Method
	switch method {
	case "GET": 
		return s.GetAccount(w, r)

	case "POST": 
		return s.CreateAccount(w, r)

	case "DELETE":
		return s.DeleteAccount(w, r)
	}

	return fmt.Errorf("method not supported %s", method)
}

func (s *Server) GetAccount (w http.ResponseWriter, r *http.Request) error {
	id := mux.Vars(r)["id"]
	fmt.Println(id)
	// Account := NewAccount("Ryan", "Gosling")

	return WriteJson(w, http.StatusOK, &Account{})
}

func (s * Server) CreateAccount (w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *Server) DeleteAccount (w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *Server) handleTransaction (w http.ResponseWriter, r *http.Request) error{
	return nil
}
