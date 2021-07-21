package main

import (
	"github.com/accesstoken/go-training/chapter1/section1.9/cfg"
	"github.com/accesstoken/go-training/chapter1/section1.9/lib/formletter"
)

func main() {
	formletter.Formletter(cfg.Name, cfg.Weather, cfg.Snack)
}
