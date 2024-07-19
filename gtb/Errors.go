package gtb

import (
	"errors"
	"fmt"
	"os"
)

func printError(err error) {
	fmt.Printf("\n💔 %v\n", err)
}

// EoE : if err is not nil print error and exit
func EoE(err error, msg ...string) {
	if err != nil {
		if len(msg) > 0 {
			errMsg := ""
			for i, m := range msg {
				if i > 0 {
					errMsg += "   " + m + "\n"
				} else {
					errMsg += Red(m) + "\n"
				}
			}
			err = errors.New(errMsg)
		}
		printError(err)
		os.Exit(1)
		panic(err)
	}
}
