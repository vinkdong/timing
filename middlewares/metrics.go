package middlewares

import (
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/VinkDong/asset-alarm/log"
	"net/http"
)
var (
	httpSendSuccess = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Help: "HTTP send success counter",
			Name: "timing_request_success"}, []string{"url", "method", "entity", "respCode"})
	httpSendFail = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Help: "HTTP send fail counter",
			Name: "timing_request_fail"}, []string{"url", "method", "entity", "respCode"})
)

func prometheusInit() {
	prometheus.MustRegister(httpSendSuccess)
	prometheus.MustRegister(httpSendFail)
}

func startPrometheus(addr *string)  {
	http.Handle("/metrics", promhttp.Handler())
	log.Info(http.ListenAndServe(*addr, nil))
}