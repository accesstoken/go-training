package arg

import (
	"flag"
)

var (
	Name string
	Shhh bool
)

func Parse() {
	flag.StringVar(&Name, "name", "<name>", "The name to say hello to")
	flag.BoolVar(&Shhh, "shhh", false, "verbose")
	flag.Parse()
}
