package highway

import (
	"fmt"
	"strings"

	"github.com/nzmprlr/highway/lane/restserver"
	"github.com/nzmprlr/highway/toll"
)

const (
	version = "v0.0.0"
)

var (
	c = &Config{}
)

type Config struct {
	HideBanner bool `json:"hideBanner"`

	Toll toll.Config          `json:"toll"`
	Rest []*restserver.Config `json:"rest"`
}

func Bootstrap(config *Config) {
	c = config

	toll.Bootstrap(&c.Toll)
	restserver.Bootstrap(c.Rest)

	printBanner()
}

func printBanner() {
	if c.HideBanner {
		return
	}

	fmt.Println(`
   __ ___      __
  / // (_)__ _/ / _    _____ ___ __
 / _  / / _ ` + "`" + `/ _ \ |/|/ / _ ` + "`" + `/ // /
/_//_/_/\_, /_//_/__,__/\_,_/\_, /
        ___/                 ___/
_  ,--.  ______/   ,   /___________
 ~'O---O'` + formatPrintVersion() + `
-----------------------------------
` + App().String())

	for _, r := range c.Rest {
		fmt.Println("> rest", r.Addr, r.Pprof, r.CORSDomains, r.CORSHeaders)
	}
}

func formatPrintVersion() string {
	return strings.Repeat(" ", 26-len(version)) + version
}
