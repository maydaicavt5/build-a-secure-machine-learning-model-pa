package main

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

// SecureMLParser struct to hold the parser instance
type SecureMLParser struct {
	Model  string `json:"model" validate:"required"`
	Config string `json:"config" validate:"required"`
}

// NewSecureMLParser creates a new instance of the parser
func NewSecureMLParser(model, config string) (*SecureMLParser, error) {
	parser := &SecureMLParser{
		Model:  model,
		Config: config,
	}
	err := validator.New().Struct(parser)
	return parser, err
}

// API struct to hold the API router
type API struct {
	Router *mux.Router
}

// NewAPI creates a new instance of the API
func NewAPI() *API {
	api := &API{
		Router: mux.NewRouter(),
	}
	api.Router.Use(cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
			http.MethodDelete,
		},
		AllowedHeaders: []string{
			"Content-Type",
			"Authorization",
		},
	}).Handler)
	return api
}

// Parse parses a machine learning model
func (api *API) Parse(w http.ResponseWriter, r *http.Request) {
	var parser SecureMLParser
	err := json.NewDecoder(r.Body).Decode(&parser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// TO DO: implement secure ML parsing logic here
	w.Write([]byte("Model parsed successfully"))
}

func main() {
	api := NewAPI()
	api.Router.HandleFunc("/parse", api.Parse).Methods("POST")
	http.ListenAndServe(":8000", api.Router)
}