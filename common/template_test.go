package common

import (
	"testing"
	"fmt"
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