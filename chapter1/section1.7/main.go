package main

import (
	"fmt"
	"github.com/accesstoken/go-training/chapter1/section1.9/cfg"
)

func formletter(name string, weather string, snack string){
	fmt.Printf("Hello %s!\nThe weather today is %s.\nToday's snack will be %s.\n", name, weather, snack)
}

func main(){
	formletter(cfg.Name, cfg.Weather, cfg.Snack)
}
