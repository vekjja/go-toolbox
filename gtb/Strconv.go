package gtb

import (
	"bytes"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/dustin/go-humanize"
)

// AToFHR : Convert string to human readable float
func AToHRF(a string) string {
	f, err := strconv.ParseFloat(a, 64)
	if err != nil {
		EoE(err)
	}
	return humanize.Commaf(f)
}

// AToI64 : Convert String to Int64
func AToI64(s string) (int64, error) {
	n, err := strconv.Atoi(s)
	if err != nil {
		return 0, err
	}
	return int64(n), nil
}

// AToU32 : Convert string to Uint32
func AToU32(a string) uint32 {

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
