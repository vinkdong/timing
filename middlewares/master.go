package middlewares

import (
	"net/http"
	"strconv"
	"github.com/vinkdong/gox/log"
	"strings"
	"github.com/vinkdong/timing/types"
	"time"
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

type Middleware interface {
	Process()
	Init(rule types.Rule)
}

var (
	sysEnableMetrics = false
)

func InitMiddleware(enableMetrics *bool, addr, buckets *string) {
	if *enableMetrics {
		go startPrometheus(addr)
		prometheusInit()
		initHistogram(ConvStringListToFloat64List(
			strings.Split(strings.Replace(*buckets, " ", "", -1), ",")))
		sysEnableMetrics = true
	}
}

func ConvStringListToFloat64List(sList []string) []float64 {
	fList := make([]float64, 0)
	for _, v := range sList {
		unit, err := strconv.ParseFloat(v,64)
		if err != nil {
			return nil
		}
		fList = append(fList, unit)
	}
	return fList
}

func ProcessMiddleware(err error, resp *http.Response, r types.Rule, entity string, start time.Time) {
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
		reqDurationHistogram.WithLabelValues(r.Url,r.Method,entity,strconv.Itoa(resqCode)).Observe(time.Since(start).Seconds())
	}
}

func SelectMiddleware(rule types.Rule) Middleware {
	var hd Middleware
	if rule.Type == "database" {
		switch rule.Database.Type {
		case "mysql":
			hd = new(MysqlMiddleware)
			break
		default:
			log.Errorf("can't find database of type %s", rule.Database.Type)
		}
	}
	if rule.Type == "" || rule.Type == "http" {
		hd = new(HttpMiddleware)
	}
	hd.Init(rule)
	return hd
}
