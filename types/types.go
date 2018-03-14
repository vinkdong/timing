package types

import (
	"github.com/vinkdong/gox/log"
	"net/http"
	"github.com/bitly/go-simplejson"
	"text/template"
	"bytes"
	"fmt"
)

type Rule struct {
	Method     string
	Url        string
	Bodies     map[string]string `yaml:"body"`
	Headers    map[string]string `yaml:"header"`
	Range      map[string]map[string]int
	Every      map[string]int    `yaml:"run_every"`
	Count      int64             `yaml:"run_count"`
	Thread     int16             `yaml:"run_thread"`
	LogResp    bool              `yaml:"log_response"`
	Prometheus map[string]string
	Check      []Checker         `yaml: "checker"`
	Type       string
	Database   Database          `yaml:"database"`
	Sql        TSql              `yaml: "sql"`
	Executed   int64
	Skip       bool
	Started    int64
}

type Database struct {
	Type     string
	Host     string
	Port     int16
	Username string
	Password string
	Database string
}

type TSql struct {
	Execute []string
	Query   []string
}

type Checker struct {
	Type string   `yaml:"type"`
	Name string   `yaml:"name"`
	Rule []string `yaml:"rule"`
}

func (c *Checker) Check(response http.Response) {

}

func (c *Checker) CheckJson(js *simplejson.Json) {

	templ, err := template.New(c.Name).Parse("")
	if err != nil{
		log.Errorf("checker %s init failed", c.Name)
	}
	var tpl bytes.Buffer

	templ.Execute(&tpl, js)
	fmt.Println(tpl.String())
}

func (r *Rule) LogNotIn(period string) {
	log.Infof("%s%s%s's %s not in request period", "\x1b[33;1m", r.Url, "\x1b[0m", period)
}