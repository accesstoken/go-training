package arg

import (
	"flag"
)

var (
	V bool
)

func Arg() {

	flag.BoolVar(&V, "v", false, "verbose")
	flag.Parse()
}
