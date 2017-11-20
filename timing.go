package main

import (
	"io/ioutil"
	"flag"
	"fmt"
	"github.com/VinkDong/asset-alarm/log"
	"gopkg.in/yaml.v2"
	"net/http"
	"strings"
	"time"
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
	prometheus = flag.Bool("metrics", false, "Provide prometheus metrics")
)

type Rule struct {
	Method string
	Url    string
	Bodies map[string]string `yaml:"body"`
	Range  map[string]map[string]int
	Every  map[string]int `yaml:"run_every"`
}

func main() {
	flag.Parse()
	if *help == true {
		showHelp()
	}
	r := &Rule{}
	parseYaml(r, *conf)
	for {
		if checkTimeIn(r) {
			for body := range r.Bodies {
				log.Infof("sending to %s ... ", r.Url)
				sendRequest(r.Method, r.Url, r.Bodies[body])
			}
		}
		d := getSleepTime(r)
		log.Infof("sleepy for %s",d.String())
		time.Sleep(d)
	}
}

func getSleepTime(r *Rule) time.Duration {
	every := r.Every
	var duration time.Duration
	if val, ok := every["seconds"]; ok {
		duration = time.Duration(val) * time.Second
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

func checkTimeIn(r *Rule) bool{
	current := time.Now().UTC()
	if rH := r.Range["hour"]; rH != nil {
		currentHour := current.Hour()
		if !checkCondition(currentHour, rH){
			log.Infof("hour %d not in request period",currentHour)
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

func sendRequest(method,url,body string){
	client  := &http.Client{}
	req, err := http.NewRequest(method,url,strings.NewReader(body))
	if err != nil {
		log.Error(err)
		return
	}
	resp, err := client.Do(req)
	if err != nil{
		log.Error(err)
		return
	}
	data , err := ioutil.ReadAll(resp.Body)
	log.Infof("%s",data)
}

func showHelp() {
	fmt.Printf(`%sTiming Request %sis used send request by timing
--conf  point a config file,must be yaml
--help  show help information
--version     show version
Need more refer at  %shttps://github.com/VinkDong/TimingRequest%s
`, CLR_Y, CLR_N, CLR_C, CLR_N)
}

func parseYaml(r *Rule, filePath string) {
	data, err := readFile(filePath)
	if err != nil {
		log.Errorf("Read config file %s error", filePath)
	}
	err = yaml.Unmarshal(data, &r)
	if err != nil {
		log.Errorf("Parse config %s error", filePath)
	}
}

func readFile(filePath string) ([]byte, error) {
	return ioutil.ReadFile(filePath)
}