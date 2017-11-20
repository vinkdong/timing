package main

import (
	"io/ioutil"
	"flag"
	"fmt"
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
	conf = flag.String("conf","","Timing request config file")
	help = flag.Bool("help",false,"Show help information")
	prometheus = flag.Bool("metrics",false,"Provide prometheus metrics")
)
func main()  {
	flag.Parse()
	if *help == true{
		showHelp()
	}
}

func showHelp() {
	fmt.Printf(`%sTiming Request %sis used send request by timing
--conf  point a config file,must be yaml
--help  show help information
--version     show version
Need more refer at  %shttps://github.com/VinkDong/TimingRequest%s
`, CLR_Y, CLR_N, CLR_C, CLR_N)
}

func parseYaml(filePath string){

}

func readFile(filePath string) ([]byte, error) {
	return ioutil.ReadFile(filePath)
}