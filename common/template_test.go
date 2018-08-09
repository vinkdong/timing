package common

import (
	"testing"
	"fmt"
	"github.com/vinkdong/gox/vtime"
	"time"
	"regexp"
)

func TestCheckExistTemplate(t *testing.T) {
	origin := "this is a test  vk.a.list{1..2} and this is vk.random{1..3} and this is vk.random "
	fmt.Println(RegisterTemplate(origin))
}

func TestExecute(t *testing.T){

}

func TestSplitFucArgs(t *testing.T)  {
	SplitFucArgs("la go la")
}

func TestVTemplate_Execute(t *testing.T) {
	vt := VTemplate{Time: vtime.Time{Time: time.Now()}}
	result := vt.Execute(`this is a test string {{ .RenderRelativeTime "now-15h" "2006-01-02" }}`)
	reg := regexp.MustCompile("this is a test string \\d{4}-\\d{2}-\\d{2}")
	if !reg.MatchString(result){
		t.Errorf("not much, got: %s",result)
	}
}