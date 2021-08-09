package main

import (
	"github.com/accesstoken/go-log"
	"os"
)

func main() {
	var logger log.Logger
	logger.Writer = os.Stdout

	simpleLogger := logger
	simpleLogger.Log("Hello world!")

	prefixedLogger := logger.Prefix("apple", "banana", "cherry")

	prefixedLogger.Log("Hello world with prefixes!")

	doublePrefixedLogger := prefixedLogger.Prefix("date")

	doublePrefixedLogger.Log("I am here with more prefixes!")
}
