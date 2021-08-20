package main

import "github.com/accesstoken/go-training/chapter4/section4.14/srv/logger"

func main(){
	var lg = logger.Get()
	lg.Alert("This is an Alert.")
	lg.Error("This is an Error.")
	lg.Warn("This is a Warn.")
	lg.Highlight("This is a Highlight.")
	lg.Inform("This is an Inform.")
	lg.Log("This is a Log.")
	lg.Trace("This is a Trace.")
}
