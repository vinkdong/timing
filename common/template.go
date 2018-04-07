package common

import (
	"regexp"
	"fmt"
	"strings"
)

const templateSpec = "vk.[a-z].[a-z]+(\\{[a-z0-9A-Z\\.]+\\})?"
const TEMPLATE  = " vk.[a-z]+(.[a-z]+)?(\\{[a-z0-9A-Z\\.]+\\})? "
const templateFuc = "(list|random)"

type VTemplate struct {
	origin   string
	result   string
	register map[string]string
}

type Register struct {
	start string
	end   string
	next  string
}

func (t *VTemplate) Execute(origin string) string {
	if checkTemplate(origin){
		t.execute(origin)
	}
	return origin
}

func (t *VTemplate) execute(origin string) string {
	val, pos := RegisterTemplate(origin)
	for index, v := range val {
		fmt.Println(pos,index,v)
	}
	return ""
}

func (t *VTemplate) splitTemplate(template string) {
	reg := regexp.MustCompile(templateSpec)
	if reg.MatchString(template) {
		t.processSpec(template)
	} else {
		process(template)
	}
}



func (t *VTemplate) processSpec(template string) {
	tmpList := strings.Split(template,".")
	register := tmpList[1]
	t.register[register] = ""
}



func process(template string) {

}


// split as list{} and list
func SplitFucArgs(mixed string){
	reg := regexp.MustCompile(templateFuc)
	for k,v := range "abcdefg"{
		fmt.Println(k,v,reg)
	}
}

func checkTemplate(origin string) bool {
	reg := regexp.MustCompile(TEMPLATE)
	return reg.MatchString(origin)
}

func randomInt(start, end int) {

}

func randomString(string, end string) {

}

func RegisterTemplate(origin string) ([]string, [][]int) {
	reg := regexp.MustCompile(TEMPLATE)
	return  reg.FindAllString(origin, -1), reg.FindAllStringIndex(origin, -1)
}
