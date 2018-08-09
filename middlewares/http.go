package middlewares

import (
	"net/http"
	"strings"
	"time"
	"io/ioutil"
	"github.com/vinkdong/timing/types"
	"github.com/vinkdong/gox/log"
	"github.com/vinkdong/timing/common"
	"compress/gzip"
	"io"
)

type HttpMiddleware struct {
	Name     string
	Rule     *types.Rule
	Template common.VTemplate
}

func (hm *HttpMiddleware) Init(rule *types.Rule) {
	hm.Rule = rule
}

func (hm *HttpMiddleware) Process() {
	// render template
	hm.Rule.Url = hm.Template.Execute(hm.Rule.Url)
	for body := range hm.Rule.Bodies {
		log.Infof("%s to %s ... ", hm.Rule.Method, hm.Rule.Url)
		hm.SendRequest(*hm.Rule, body)
	}
}

/**
send http request
 */
func (hm *HttpMiddleware) SendRequest(r types.Rule, entity string) {
	client := &http.Client{}
	body := r.Bodies[entity]

	//render template
	body = hm.Template.Execute(body)
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
	
	if r.LogResp {
		var r io.Reader
		switch resp.Header.Get("Content-Encoding") {
		case "gzip":
			r , err = gzip.NewReader(resp.Body)
			if err != nil{
				log.Error(err)
			}
		default:
			r = resp.Body
		}

		data, err := ioutil.ReadAll(r)
		if err != nil{
			log.Error(err)
		}
		log.Infof("%s", data)
	}
}