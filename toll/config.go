package toll

var (
	c = &Config{}
)

type Config struct {
	Quiet  bool   `json:"quiet"`
	Indent string `json:"indent"`
}

func Bootstrap(config *Config) {
	c = config
}
