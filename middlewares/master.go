package middlewares

import (
	"net/http"
	"strconv"
	"github.com/VinkDong/asset-alarm/log"
)

type Rule struct {
	Method     string
	Url        string
	Bodies     map[string]string `yaml:"body"`
	Range      map[string]map[string]int
	Every      map[string]int    `yaml:"run_every"`
	LogResp    bool              `yaml:"log_response"`
	Prometheus map[string]string
}

var (
	sysEnableMetrics = false
)

func InitMiddleware(enableMetrics *bool, addr *string) {
	if *enableMetrics{
		go startPrometheus(addr)
		prometheusInit()
		sysEnableMetrics = true
	}
}

func ProcessMiddleware(err error, resp *http.Response, rule interface{}, entity string) {
	r, _ := rule.(Rule)
	url := r.Url
	method := r.Method
	if err != nil {
		log.Error(err)
		httpSendFail.WithLabelValues(url, method, entity, "CLIENT_ERR_019")
		return
	}
	resqCode := resp.StatusCode
	if sysEnableMetrics {
		if resqCode >= 200 && resqCode < 400 {
			httpSendSuccess.WithLabelValues(url, method, entity, strconv.Itoa(resqCode)).Inc()
		} else {
			httpSendFail.WithLabelValues(url, method, entity, strconv.Itoa(resqCode)).Inc()
		}
	}
}
