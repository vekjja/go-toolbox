package toolbox

import (
	"encoding/json"
	"fmt"

	"github.com/TylerBrock/colorjson"
)

// PrettyJson takes a JSON byte slice and prints it in a pretty, colorized format
func PrettyJson(data []byte) {
	// Unmarshal JSON into a map
	var obj map[string]interface{}
	EoE(json.Unmarshal(data, &obj))

	// Create a colorized JSON formatter
	formatter := colorjson.NewFormatter()
	formatter.Indent = 2

	// Marshal the object to a colorized JSON string
	s, err := formatter.Marshal(obj)
	EoE(err)

	// Print the colorized JSON
	fmt.Println(string(s))
}
