package types

type Rule struct {
	Method     string
	Url        string
	Bodies     map[string]string `yaml:"body"`
	Range      map[string]map[string]int
	Every      map[string]int    `yaml:"run_every"`
	LogResp    bool              `yaml:"log_response"`
	Prometheus map[string]string
}
