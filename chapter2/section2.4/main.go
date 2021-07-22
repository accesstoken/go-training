package main

import (
	"fmt"
	"github.com/accesstoken/go-training/chapter2/section2.4/arg"
)

func main() {
	arg.Arg()
	if !arg.Shhh {
		fmt.Printf("Hello %s!\n", arg.Name)
	}
}
