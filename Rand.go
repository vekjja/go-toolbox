package toolbox

import (
	"encoding/base64"
	"math/rand"
)

// Int : return pseudo random integer in range
func RandInt(min, max int) int {
	return rand.Intn(max-min) + min
}

// Float : return pseudo random float32
func RandFloat() float32 {
	return rand.Float32()
}

// Token : generate random token, use case: oath2
func RandToken() string {
	b := make([]byte, 32)
	return base64.StdEncoding.EncodeToString(b)
}
