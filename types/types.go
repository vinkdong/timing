package types

import (
	"github.com/vinkdong/asset-alarm/log"
)

type Rule struct {
	Method     string
	Url        string
	Bodies     map[string]string `yaml:"body"`
	Headers    map[string]string `yaml:"header"`
	Range      map[string]map[string]int
	Every      map[string]int    `yaml:"run_every"`
	LogResp    bool              `yaml:"log_response"`
	Prometheus map[string]string
	Check      []Check           `yaml: "check"`
}

type Check struct {
	Type string `yaml:"type"`
	Rule string `yaml:"rule"`
}

func (r *Rule) LogNotIn(period string) {
	log.Infof("%s%s%s's %s not in request period", "\x1b[33;1m", r.Url, "\x1b[0m", period)
}