package restserver

import (
	"fmt"
	"net/http"
	"net/url"
	"runtime/debug"

	// add default pprof handlers to default http serve mux.
	_ "net/http/pprof"
)

func (rest *Rest) routes() {
	rest.Panic(rest.panic())
	rest.NotFound(rest.notFound())
	rest.MethodNotAllowed(rest.methodNotAllowed())
	rest.Options(rest.options())

	rest.health()
	rest.pprof()
}

func (rest *Rest) Panic(handler func(http.ResponseWriter, *http.Request, interface{})) {
	rest.router.PanicHandler = handler
}
func (rest *Rest) panic() func(http.ResponseWriter, *http.Request, interface{}) {
	return func(w http.ResponseWriter, r *http.Request, panic interface{}) {
		toll := ParseToll(r)
		toll.Label("panic", fmt.Sprintf("%s\n%s", panic, debug.Stack()))

		toll.SetStatus(http.StatusInternalServerError)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (rest *Rest) NotFound(handler http.HandlerFunc) { rest.router.NotFound = handler }
func (rest *Rest) notFound() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		toll := ParseToll(r)
		defer toll.End()
		defer toll.Metric(toll.Metrics.Controller())

		requestLabels(toll, r)
		toll.Log("method: %s, uri: %s", r.Method, r.RequestURI)

		toll.SetStatus(http.StatusNotFound)
		w.WriteHeader(http.StatusNotFound)
	}
}

func (rest *Rest) MethodNotAllowed(handler http.HandlerFunc) { rest.router.MethodNotAllowed = handler }
func (rest *Rest) methodNotAllowed() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		toll := ParseToll(r)
		defer toll.End()
		defer toll.Metric(toll.Metrics.Controller())

		requestLabels(toll, r)
		toll.Log("method: %s, uri: %s", r.Method, r.RequestURI)

		toll.SetStatus(http.StatusMethodNotAllowed)
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (rest *Rest) Options(handler http.HandlerFunc) { rest.router.GlobalOPTIONS = handler }
func (rest *Rest) options() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		toll := ParseToll(r)
		defer toll.End()
		defer toll.Metric(toll.Metrics.Controller())

		requestLabels(toll, r)

		// options handler
		if r.Header.Get("Access-Control-Request-Method") != "" {
			// Set CORS headers
			rest.setCORSHeaders(w)
		}

		// Adjust status code to 204
		toll.SetStatus(http.StatusNoContent)
		w.WriteHeader(http.StatusNoContent)
	}
}

func (rest *Rest) health() {
	if !rest.config.Health {
		return
	}

	rest.AddHandler(http.MethodGet, "/health", http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			toll := ParseToll(r)
			defer toll.End()
			defer toll.Metric(toll.Metrics.Controller())

			requestLabels(toll, r)
		}))
}

func (rest *Rest) pprof() {
	if rest.config.Pprof == "" {
		return
	}

	rest.AddHandler(http.MethodGet, rest.config.Pprof+"/*param", http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			u, err := url.Parse("/debug/pprof" + ParseRequestParams(r)["param"])
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			r.URL = u
			r.RequestURI = u.RequestURI()
			http.DefaultServeMux.ServeHTTP(w, r)
		}))
}
