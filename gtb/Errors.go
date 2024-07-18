package gtb

import (
	"fmt"
	"os"
)

func printError(err error, msg ...string) {
	fmt.Printf("\n💔 %s\n   %v\n", msg, err)
}

// EoE : exit on error, if err is not nil
func EoE(err error, msg ...string) {
	if err != nil {
		printError(err, msg...)
		os.Exit(1)
		panic(err)
	}
}
