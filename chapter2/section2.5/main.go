package main

import (
	"fmt"
	"github.com/accesstoken/go-training/chapter2/section2.5/arg"
)

func main() {
	if !arg.Shhh {
		fmt.Printf("Hello %s!\n", arg.Name)
	}
}
