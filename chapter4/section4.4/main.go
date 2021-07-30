package main

import (
	"fmt"
	"github.com/accesstoken/go-log"
	"os"
)

func main() {
	var logger log.Logger
	logger.Writer = os.Stdout
	//logger.Writer = ioutil.Discard
	logger.Begin()
	var name string = "Joe"
	logger.Logf("The name was %q\n", name)
	var msg string = "Hello world!"
	fmt.Println(msg)
	logger.Log("I said:", msg)
	logger.End()
}