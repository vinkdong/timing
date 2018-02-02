package middlewares

import (
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/vinkdong/asset-alarm/log"
	"net/http"
)

const (
	metricPrefix        = "timing_request"
	reqSuccessTotalName = metricPrefix + "_success_total"
	reqFailTotalName    = metricPrefix + "_fail_total"
	reqDurationName     = metricPrefix + "_duration_seconds"
)

var (
	httpSendSuccess = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Help: "HTTP send success counter",
			Name: reqSuccessTotalName}, []string{"url", "method", "entity", "respCode"})
	httpSendFail = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Help: "HTTP send fail counter",
			Name: reqFailTotalName}, []string{"url", "method", "entity", "respCode"})
	reqDurationHistogram = &prometheus.HistogramVec{}
)

func initHistogram(buckets []float64) {
	reqDurationHistogram = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name:    reqDurationName,
		Help:    "How long it took to process the request.",
		Buckets: buckets,
	}, []string{"url", "method", "entity", "respCode"})
	prometheus.MustRegister(reqDurationHistogram)
}

func prometheusInit() {
	prometheus.MustRegister(httpSendSuccess)
	prometheus.MustRegister(httpSendFail)
}

func startPrometheus(addr *string)  {
	http.Handle("/metrics", promhttp.Handler())
	log.Info(http.ListenAndServe(*addr, nil))
}