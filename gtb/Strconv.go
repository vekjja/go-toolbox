package gtb

import (
	"bytes"
	"fmt"
	"log"
	"strconv"
	"strings"
)

// AToUint32 : Convert string to Uint32
func AToUint32(a string) uint32 {

	i, err := strconv.Atoi(a)
	if err != nil {
		log.Fatal(err)
	}

	return uint32(i)
}

// FtoA : convert float32 to string
func FtoA(n float32) string {
	return strconv.FormatFloat(float64(n), 'f', 6, 32)
}

// MapToString : convert map[string]string to string
func MapToString(m map[string]string) string {
	b := new(bytes.Buffer)
	for key, value := range m {
		fmt.Fprintf(b, "  %s: %s\n", key, value)
	}
	return b.String()
}

// SplitMulti : Split String on Multiple Delimiters
func SplitMulti(s string, delims string) []string {
	splitter := func(r rune) bool {
		return strings.ContainsRune(delims, r)
	}
	return strings.FieldsFunc(s, splitter)
}

// StringToInt64 : Convet String to Int64
func StringToInt64(s string) (int64, error) {
	n, err := strconv.Atoi(s)
	if err != nil {
		return 0, err
	}
	return int64(n), nil
}
