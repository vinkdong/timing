package log

import (
	"fmt"
	"time"
	"log"
)

func Info(l interface{}) {
	info := fmt.Sprintf("%s [%s] %v", time.Now().String(), "INFO", l)
	log.Println(info)
}

func Infof(l string, a ...interface{}) {
	tmp := fmt.Sprintf(l, a...)
	info := fmt.Sprintf("%s [%s] %s", time.Now().String(), "INFO", tmp)
	log.Println(info)
}

func Error(l interface{}) {
	err := fmt.Sprintf("%s [%s] %v", time.Now().String(), "Error", l)
	log.Println(err)
}

func Errorf(format string, a ...interface{}) {
	tmp := fmt.Sprintf(format, a...)
	err := fmt.Sprintf("%s [%s] %s", time.Now().String(), "Error", tmp)
	log.Println(err)
}
