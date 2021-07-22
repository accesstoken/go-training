package arg

import (
	"flag"
)

var (
	Name string
	Shhh bool
)


func init() {
	flag.StringVar(&Name, "name", "<name>", "The name to say hello to")
	flag.BoolVar(&Shhh, "shhh", false, "verbose")
	flag.Parse()
}
