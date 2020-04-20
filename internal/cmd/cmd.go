package cmd

import (
	"fmt"
	"os"

	flag "github.com/spf13/pflag"
)

var (
	help bool
	test bool
)

func init() {
	flag.BoolVarP(&help, "help", "h", false, "show help")
	flag.BoolVarP(&test, "test", "t", true, "test-flag")
	flag.Parse()
	if help {
		flag.PrintDefaults()
		os.Exit(0)
	}
}

func Run() {
	fmt.Println(test)
}
