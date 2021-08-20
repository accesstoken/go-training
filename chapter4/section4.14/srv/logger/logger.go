package logger

import (
	"github.com/accesstoken/go-log"
	"github.com/accesstoken/go-training/chapter4/section4.14/arg"
	"os"
)

func Get() log.Logger{
	var logger log.Logger
	logger.Writer = os.Stdout
	leveledLogger := logger.Level(arg.Level)
	return leveledLogger
}
