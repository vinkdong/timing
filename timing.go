package main

import (
	"io/ioutil"
	"flag"
	"fmt"
	"github.com/vinkdong/gox/log"
	"gopkg.in/yaml.v2"
	"time"
	"bytes"
	"os"
	"github.com/vinkdong/timing/middlewares"
	"github.com/vinkdong/timing/types"
)

const CLR_0 = "\x1b[30;1m"
const CLR_N = "\x1b[0m"
const CLR_R = "\x1b[31;1m"
const CLR_G = "\x1b[32;1m"
const CLR_Y = "\x1b[33;1m"
const CLR_B = "\x1b[34;1m"
const CLR_M = "\x1b[35;1m"
const CLR_C = "\x1b[36;1m"
const CLR_W = "\x1b[37;1m"
const VERSION = "v0.1.0"

var (
	conf       = flag.String("conf", "", "Timing request config file")
	help       = flag.Bool("help", false, "Show help information")
	enableMetrics = flag.Bool("enable_metrics", false, "Provide prometheus metrics")
	addr       = flag.String("addr", ":9800", "The address to listen on for HTTP requests.")
	buckets    = flag.String("buckets", "0.1, 0.3, 1.2, 5.0", "Buckets holds Prometheus Buckets")

)




func main() {
	flag.Parse()
	if *help == true {
		showHelp()
	}
	ruleList := make([]types.Rule,0)
	middlewares.InitMiddleware(enableMetrics,addr,buckets)
	parseYaml(&ruleList, *conf)
	for _, r := range ruleList {
		var p = r
		p.Started = time.Now().UnixNano()
		go dealSend(&p)
	}
	for{
		time.Sleep(time.Hour*10)
	}
}

func dealSend(r *types.Rule) {
	for {
		if checkTimeIn(r) {
			var middleware middlewares.Middleware
			middleware = middlewares.SelectMiddleware(r)
			middleware.Process()
		}
		d := getSleepTime(r)
		if d > 0 {
			log.Infof("sleepy for %s", d.String())
			time.Sleep(d)
		}
	}
}

func getSleepTime(r *types.Rule) time.Duration {
	every := r.Every
	var duration time.Duration
	if val, ok := every["seconds"]; ok {
		duration = time.Duration(val) * time.Second
	}
	if val, ok := every["microsecond"]; ok {
		duration = time.Duration(val) * time.Microsecond
	}
	if val, ok := every["minutes"]; ok {
		duration += time.Duration(val) * time.Minute
	}
	if val, ok := every["hours"]; ok {
		duration += time.Duration(val) * time.Hour
	}
	if val, ok := every["day"]; ok {
		duration += time.Duration(val) * time.Hour * 24
	}
	return duration
}

func checkTimeIn(r *types.Rule) bool{
	if r.Skip{
		return false
	}
	current := time.Now().UTC()
	if rH := r.Range["month"]; rH != nil {
		currentMonth := int(current.Month())
		if !checkCondition(currentMonth, rH){
			r.LogNotIn("month")
			return false
		}

	}
	if rH := r.Range["weekday"]; rH != nil {
		currentWeekday := int(current.Weekday())
		if !checkCondition(currentWeekday, rH){
			r.LogNotIn("weekday")
			return false
		}

	}
	if rH := r.Range["hour"]; rH != nil {
		currentHour := current.Hour()
		if !checkCondition(currentHour, rH){
			r.LogNotIn("hour")
			return false
		}

	}
	if rH := r.Range["minute"]; rH != nil {
		currentMinute := current.Minute()
		if !checkCondition(currentMinute, rH){
			r.LogNotIn("minute")
			return false
		}
	}
	if rH := r.Range["second"]; rH != nil {
		currentSecond := current.Second()
		if !checkCondition(currentSecond, rH){
			r.LogNotIn("second")
			return false
		}
	}
	return true
}

func checkCondition(current int, condition map[string]int) bool {
	if v, ok := condition["gt"]; ok {
		if current <= v{
			return false
		}
	}

	if v, ok := condition["gte"]; ok {
		if current < v{
			return false
		}
	}

	if v, ok := condition["lt"]; ok {
		if current >= v{
			return false
		}
	}

	if v, ok := condition["lte"]; ok {
		if current > v{
			return false
		}
	}

	if v, ok := condition["eq"]; ok {
		if current != v{
			return false
		}
	}

	return true
}



func showHelp() {
	fmt.Printf(`%sTiming Request %sis used send request by timing
--conf  point a config file,must be yaml
--help  show help information
--version     show version
Need more refer at  %shttps://github.com/VinkDong/TimingRequest%s
`, CLR_Y, CLR_N, CLR_C, CLR_N)
}

func parseYaml(rList *[]types.Rule, filePath string) {
	data, err := readFile(filePath)
	if err != nil {
		log.Errorf("Read config file %s error", filePath)
	}
	dataList := bytes.Split(data, []byte("\n---"))
	for _, single := range dataList {
		r := types.Rule{}
		err = yaml.Unmarshal(single, &r)
		if err != nil {
			log.Errorf("Parse config %s error", filePath)
			os.Exit(128)
		}
		*rList = append(*rList, r)
	}
}

func readFile(filePath string) ([]byte, error) {
	return ioutil.ReadFile(filePath)
}