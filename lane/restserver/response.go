package restserver

import (
	"encoding/json"
	"net/http"

	"github.com/nzmprlr/highway/toll"
)

type Response interface {
	Format() error
}

func respondJSON(status uint16, w http.ResponseWriter, r interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(int(status))

	return json.NewEncoder(w).Encode(r)
}

func RespondJSON(t *toll.Toll, w http.ResponseWriter, response Response) error {
	err := response.Format()
	if err != nil {
		return err
	}

	return respondJSON(t.Status, w, response)
}

func RespondErrorJSON(t *toll.Toll, w http.ResponseWriter, response interface{}) {
	if err, ok := response.(error); ok {
		response = struct {
			Error string `json:"error"`
		}{err.Error()}
	}

	err := respondJSON(t.Status, w, response)
	if err != nil {
		panic(err)
	}
}
