package env

import (
	"os"

	"github.com/nzmprlr/highway/asphalt/parse"
)

func Get(key, def string) string {
	env, ok := os.LookupEnv(key)
	if ok {
		return env
	}

	return def
}

func GetInt(key string, def int) int {
	env, ok := os.LookupEnv(key)
	if ok {
		p, err := parse.Int(env)
		if err != nil {
			return def
		}

		return p[0]
	}

	return def
}
