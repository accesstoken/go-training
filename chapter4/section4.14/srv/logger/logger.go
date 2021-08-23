package logger

import (
	"github.com/accesstoken/go-log"
	"github.com/accesstoken/go-training/chapter4/section4.14/arg"
	"os"
)

var Logger log.Logger

func init(){
	var logger log.Logger
	logger.Writer = os.Stdout
	Logger = logger.Level(arg.Level)
}
