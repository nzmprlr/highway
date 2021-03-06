package restserver

type typeContextKey string

const (
	headerKeyDebug  string         = "x-debug-id"
	contextKeyDebug typeContextKey = "ctx-key-debug"
)

type Config struct {
	Addr string `json:"addr"`

	Pprof  string `json:"pprof"`
	Health bool   `json:"health"`

	MaxRequestBodyByte int64 `json:"maxRequestBodyByte"`

	ReadHeaderTimeoutSecond uint `json:"readHeaderTimeoutSecond"`
	ReadTimeoutSecond       uint `json:"readTimeoutSecond"`
	WriteTimeoutSecond      uint `json:"writeTimeoutSecod"`
	IdleTimeoutSecond       uint `json:"idleTimeoutSecond"`

	CORSDomains []string `json:"corsDomains"`
	corsDomains string   `json:"-"`
	CORSHeaders []string `json:"corsHeaders"`
	corsHeaders string   `json:"-"`
}
