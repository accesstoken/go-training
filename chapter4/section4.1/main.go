package main

import (
	"fmt"
	"github.com/accesstoken/go-training/chapter4/section4.1/logger"
	"os"
)

func main() {
	var log logger.Logger
	log.Writer = os.Stdout
	//log.Writer = ioutil.Discard
	log.Begin()

	var msg string = "Hello world!"
	fmt.Println(msg)
	log.Log("I said:", msg)

	log.End()
}