package gtb

import (
	"fmt"
	"os"
)

func printError(err error) {
	fmt.Printf("\n💔 %v\n", err)
}

// EoE : if err is not nil print error and exit
func EoE(err error) {
	if err != nil {
		printError(err)
		os.Exit(1)
		panic(err)
	}
}
