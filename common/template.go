package common

import (
	"regexp"
	"fmt"
)

type Template struct {
	origin string
	result string
	register map[int64]string
}

type Register struct {

}

func CheckTemplateReplace(origin string) string {
	return ""
}

func randomInt(start, end int) {

}

func randomString(string, end string) {

}

func RegisterTemplate(origin string)  {
	reg := regexp.MustCompile(" vk.[a-z]+(\\{[a-z0-9A-Z\\.]+\\})? ")
	fmt.Println(reg.FindAllString(origin,-1))
	fmt.Println(reg.FindAllStringIndex(origin,-1))
}