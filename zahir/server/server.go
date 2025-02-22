package server

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"zahir/data"
	"zahir/player"

	"github.com/gorilla/mux"
)

var Router *mux.Router

func RunServer() {
	createRouter()
	fmt.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func createRouter() {
	Router = mux.NewRouter()

	Router.HandleFunc("/v1/sources", getSources).Methods("GET")
	Router.HandleFunc("/v1/sources/{id}", getSource).Methods("GET")
	Router.HandleFunc("/v1/sources", createSource).Methods("POST")

	Router.HandleFunc("/v1/sequences", getSequences).Methods("GET")
	Router.HandleFunc("/v1/sequences/{id}", getSequenceData).Methods("GET")
	Router.HandleFunc("/v1/sequences/{id}/run", runSequence).Methods("POST")
	Router.HandleFunc("/v1/sequences/{id}/values", setVariableValues).Methods("POST")

	Router.HandleFunc("/v1/player/start", startPlayer).Methods("POST")
	Router.HandleFunc("/v1/player/stop", stopPlayer).Methods("POST")

	Router.HandleFunc("/v1/ws", wsState)
	http.Handle("/", Router)
}

// /v1/sources
func getSources(w http.ResponseWriter, r *http.Request) {
	sources, err := data.GetSources()
	if err != nil {
		returnError(w, "Failed to get sources")
		return
	}
	returnJson(w, sources)
}

// v1/sources/{id}
func getSource(w http.ResponseWriter, r *http.Request) {
	id, err := getRequestID(r)
	if err != nil {
		returnNotFound(w, r)
		return
	}
	source, err := data.GetSource(id)
	if err != nil {
		returnNotFound(w, r)
		return
	}
	returnJson(w, source)
}

// /v1/sources (POST)
func createSource(w http.ResponseWriter, r *http.Request) {
	err := data.CreateSource()
	if err != nil {
		returnError(w, "Failed to create source")
		return
	}
}

// /v1/sequences
func getSequences(w http.ResponseWriter, r *http.Request) {
	sequences, err := data.GetSequences()
	if err != nil {
		returnError(w, "Failed to get sequences")
		return
	}
	returnJson(w, sequences)
}

// /v1/sequences/{id}
func getSequenceData(w http.ResponseWriter, r *http.Request) {
	sequenceID, err := getRequestID(r)
	if err != nil {
		returnNotFound(w, r)
		return
	}
	sequence, err := data.GetSequence(sequenceID)
	if err != nil {
		returnNotFound(w, r)
		return
	}
	sources, err := data.GetSequenceSources(sequenceID)
	if err != nil {
		returnNotFound(w, r)
		return
	}
	variableValues, err := data.GetVariableValues(sequenceID)
	if err != nil {
		returnNotFound(w, r)
		return
	}

	response := map[string]interface{}{
		"sequence": sequence,
		"sources":  sources,
		"values":   variableValues,
	}
	returnJson(w, response)
}

// /v1/sequences/{id}/run
func runSequence(w http.ResponseWriter, r *http.Request) {
	id, err := getRequestID(r)
	if err != nil {
		returnNotFound(w, r)
		return
	}
	_, err = data.GetSequence(id)
	if err != nil {
		returnNotFound(w, r)
		return
	}
	player.RunSequence(id)
}

// /v1/sequences/{id}/values
func setVariableValues(w http.ResponseWriter, r *http.Request) {
	type bodyContent struct {
		VariableID int           `json:"variable_id"`
		Values     []interface{} `json:"values"`
	}

	sequenceID, err := getRequestID(r)
	if err != nil {
		returnNotFound(w, r)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		returnError(w, "Failed to read body")
		return
	}
	defer r.Body.Close()

	var bc bodyContent
	err = json.Unmarshal(body, &bc)
	if err != nil {
		returnError(w, "Failed to decode body")
		return
	}

	data.SetVariableValues(sequenceID, bc.VariableID, bc.Values)
}

// /v1/player/start
func startPlayer(w http.ResponseWriter, r *http.Request) {
	player.SetRunning(true)
}

// /v1/player/stop
func stopPlayer(w http.ResponseWriter, r *http.Request) {
	player.SetRunning(false)
}
