package restserver

import (
	"context"
	"net/http"
	"strings"
	"time"

	"github.com/julienschmidt/httprouter"
)

var (
	rests = []*Rest{}
)

// Rest represents rest server
type Rest struct {
	config *Config
	server *http.Server
	router *httprouter.Router
}

func (rest *Rest) AddHandler(method, path string, handler http.Handler) {
	rest.router.Handler(method, path, handler)
}

// ListenAndServe starts rest server.
func (rest *Rest) ListenAndServe() {
	go func() {
		// debug.Log("server is on: %s", rest.server.Addr)
		if err := rest.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			panic(err)
		}
	}()
}

// GracefulShutdown shutdowns the rest server.
func GracefulShutdown() {
	for _, rest := range rests {
		if rest.server == nil {
			return
		}

		rest.server.SetKeepAlivesEnabled(false)
		if err := rest.server.Shutdown(context.Background()); err != nil {
			// debug.Log("Server Shutdown Failed:%+s", err)
		}
	}

}

// TODO secure server
func Bootstrap(c []*Config) {
	rests = make([]*Rest, len(c))

	for i, r := range c {
		router := httprouter.New()
		server := newHTTPServer(r, router)

		r.corsDomains = strings.Join(r.CORSDomains, ",")
		r.corsHeaders = strings.Join(r.CORSHeaders, ",")

		rest := &Rest{}
		// inject dependicies here.
		rest.server = server
		rest.router = router
		rest.config = r

		rest.routes()
		rest.ListenAndServe()

		rests[i] = rest
	}
}

func Get() []*Rest {
	return rests
}

func newHTTPServer(c *Config, handler http.Handler) *http.Server {
	return &http.Server{
		Addr:              c.Addr,
		ReadHeaderTimeout: time.Second * time.Duration(c.ReadHeaderTimeoutSecond),
		ReadTimeout:       time.Second * time.Duration(c.ReadTimeoutSecond),
		WriteTimeout:      time.Second * time.Duration(c.WriteTimeoutSecond),
		IdleTimeout:       time.Second * time.Duration(c.IdleTimeoutSecond),

		Handler: handler,
	}
}
