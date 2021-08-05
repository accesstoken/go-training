package arg

import (
	"flag"
)

var (
	Verbose bool
)

func init() {
	flag.BoolVar(&Verbose, "v", false, "verbose")
	flag.Parse()
}
