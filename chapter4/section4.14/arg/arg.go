package arg

import (
	"flag"
)

var (
	v bool
	vv bool
	vvv bool
	vvvv bool
	vvvvv bool
	vvvvvv bool
	Level uint8
)

func init() {
	flag.BoolVar(&v, "v", false, "verbose level 1")
	flag.BoolVar(&vv, "vv", false, "verbose level 2")
	flag.BoolVar(&vvv, "vvv", false, "verbose level 3")
	flag.BoolVar(&vvvv, "vvvv", false, "verbose level 4")
	flag.BoolVar(&vvvvv, "vvvvv", false, "verbose level 5")
	flag.BoolVar(&vvvvvv, "vvvvvv", false, "verbose level 6")
	flag.Parse()
	if v {
		Level = 1
		return
	}
	if vv {
		Level = 2
		return
	}
	if vvv {
		Level = 3
		return
	}
	if vvvv {
		Level = 4
		return
	}
	if vvvvv {
		Level = 5
		return
	}
	if vvvvvv {
		Level = 6
		return
	}
}
