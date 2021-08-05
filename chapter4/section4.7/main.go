package main

import (
	"github.com/accesstoken/go-log"
	"go-training/chapter4/section4.7/arg"
	"io"
	"os"
)

func main() {
	var logger log.Logger
	if arg.Verbose {
		logger.Writer = os.Stdout
	} else {
		logger.Writer = io.Discard
	}
	logger.Begin()
	logger.End()
}
