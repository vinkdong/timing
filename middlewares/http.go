package middlewares

import (
	"net/http"
	"strings"
	"time"
	"io/ioutil"
	"github.com/vinkdong/timing/types"
	"github.com/vinkdong/gox/log"
)

type HttpMiddleware struct {
	Name string
	Rule types.Rule
}

func (hm *HttpMiddleware) Process() {
	for body := range hm.Rule.Bodies {
		log.Infof("sending to %s ... ", hm.Rule.Url)
		hm.SendRequest(hm.Rule, body)
	}
}

func (hm *HttpMiddleware) SendRequest(r types.Rule, entity string) {
	client := &http.Client{}
	body := r.Bodies[entity]
	req, err := http.NewRequest(r.Method, r.Url, strings.NewReader(body))
	for k , v := range r.Headers{
		req.Header.Set(k,v)
	}
	if err != nil {
		log.Error(err)
		return
	}
	start := time.Now()
	resp, err := client.Do(req)
	ProcessMiddleware(err, resp, r, entity, start)
	if err != nil {
		return
	}
	data, err := ioutil.ReadAll(resp.Body)
	if r.LogResp {
		log.Infof("%s", data)
	}
}