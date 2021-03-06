package highway

import (
	"fmt"
	"time"

	"github.com/nzmprlr/highway/asphalt/env"
)

const (
	unknown = "unknown"

	envKey   = "APP_ENV"
	EnvProd  = "prod"
	EnvDev   = "dev"
	EnvTest  = "test"
	EnvLocal = "local"
)

var (
	app = initApp()
)

// TODO instance, source usage
type AppInfo struct {
	module  string
	env     string
	id      string
	rev     string
	builtAt string

	startedAt time.Time
}

func (a *AppInfo) String() string {
	return fmt.Sprintf("%s-%s %s #%s %s", a.module, a.env, a.rev, a.id, a.builtAt)
}

func (a *AppInfo) RunTime() string {
	return time.Since(a.startedAt).String()
}

func (a *AppInfo) Inject(module, id, rev, builtAt string) {
	a.module = module
	a.id = id
	a.rev = rev
	a.builtAt = builtAt
}

func initApp() *AppInfo {
	return &AppInfo{
		module:  unknown,
		id:      unknown,
		rev:     unknown,
		builtAt: unknown,

		env: env.Get(envKey, EnvLocal),

		startedAt: time.Now().UTC(),
	}
}

func App() *AppInfo {
	return app
}

func Env() string {
	return App().env
}

func IsEnvProd() bool {
	return Env() == EnvProd
}

func IsEnvDev() bool {
	return !IsEnvProd()
}

func IsEnvLocal() bool {
	return Env() == EnvLocal
}
