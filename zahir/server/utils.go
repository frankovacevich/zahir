package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func returnNotFound(w http.ResponseWriter, r *http.Request) {
	http.NotFound(w, r)
}

func returnError(w http.ResponseWriter, msg string) {
	err := fmt.Errorf("%s", msg)
	http.Error(w, err.Error(), http.StatusInternalServerError)
}

func getRequestID(r *http.Request) (int, error) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	return id, err
}

func returnJson(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
