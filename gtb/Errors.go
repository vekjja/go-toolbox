package gtb

import (
	"fmt"
	"os"
)

func FormatError(err error, msg ...string) error {
	if len(msg) > 0 {
		errMsg := ""
		for _, m := range msg {
			errMsg += Red(m)
		}
		err = fmt.Errorf(errMsg)
	}
	return fmt.Errorf("\n💔 %s", Red(err.Error()))
}

// EoE : if err is not nil format and print error then exit
func EoE(err error, msg ...string) {
	if err != nil {
		fmt.Println(FormatError(err, msg...))
		os.Exit(1)
	}
}
