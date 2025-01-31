package toolbox

import "strings"

// Confirm : return confirmation based on user input
func Confirm(q string) bool {
	print(q + " (Y/n):")
	a := Input()
	var res bool
	switch strings.ToLower(a) {
	case "":
		fallthrough
	case "y":
		fallthrough
	case "yes":
		res = true
	case "n":
		fallthrough
	case "no":
		res = false
	default:
		return Confirm(q)
	}
	return res
}
