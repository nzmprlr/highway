package restserver

import (
	"context"
	"net/http"

	"github.com/nzmprlr/highway/toll"
)

func (rest *Rest) Handler(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rest.HandlerWithMaxBodySize(h, rest.config.MaxRequestBodyByte)(w, r)
	}
}

func (rest *Rest) HandlerWithMaxBodySize(h http.HandlerFunc, maxBodySize int64) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		toll := toll.NewWithID(r.Header.Get(headerKeyDebug))

		// set request info labels
		requestLabels(toll, r)

		// inject toll to request context
		r = r.WithContext(context.WithValue(r.Context(), contextKeyDebug, toll))

		// recover panic with debugged request
		defer func() {
			if p := recover(); p != nil {
				rest.router.PanicHandler(w, r, p)
			}

			toll.End()
		}()

		// enforce a maximum read from the request body
		r.Body = http.MaxBytesReader(w, r.Body, maxBodySize)

		// set cors headers for all requests
		rest.setCORSHeaders(w)

		// continue with tolled request
		h(w, r)
	}
}

func (rest *Rest) setCORSHeaders(w http.ResponseWriter) {
	header := w.Header()

	if allow := header.Get("Allow"); allow != "" {
		header.Set("Access-Control-Allow-Methods", allow)
	}

	header.Set("Access-Control-Allow-Origin", rest.config.corsDomains)
	header.Set("Access-Control-Allow-Headers", rest.config.corsHeaders)
}

func requestLabels(t *toll.Toll, r *http.Request) {
	t.Label("remote-addr", r.RemoteAddr)

	header := r.Header
	t.Label("user-agent", header.Get("User-Agent"))

	if realIP := header.Get("X-Real-IP"); realIP != "" {
		t.Label("x-real-ip", realIP)
	}

	if forwardedIP := header.Get("X-Forwarded-For"); forwardedIP != "" {
		t.Label("x-forwarded-for", forwardedIP)
	}
}
