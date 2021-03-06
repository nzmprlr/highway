package restserver

import (
	"encoding/json"
	"net/http"

	"github.com/nzmprlr/highway/toll"

	"github.com/julienschmidt/httprouter"
)

type Request interface {
	Parse(*http.Request) error
	Validate() error
}

func ParseRequest(r *http.Request, i Request) error {
	if err := i.Parse(r); err != nil {
		return err
	}

	return i.Validate()
}

func ParseRequestParams(r *http.Request) map[string]string {
	params := map[string]string{}
	if r == nil {
		return params
	}

	p := httprouter.ParamsFromContext(r.Context())
	for _, v := range p {
		params[v.Key] = v.Value
	}

	return params
}

func ParseRequestBodyJSON(r *http.Request, v interface{}) error {
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	return decoder.Decode(v)
}

func ParseToll(r *http.Request) *toll.Toll {
	t, ok := r.Context().Value(contextKeyDebug).(*toll.Toll)
	if !ok {
		return toll.New()
	}

	return t
}
