package gtb

import (
	"encoding/base64"
	"math/rand"
	"time"
)

// Int : return pseudo random integer in range
func RandInt(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min) + min
}

// Float : return pseudo random float32
func RandFloat() float32 {
	rand.Seed(time.Now().UnixNano())
	return rand.Float32()
}

// Token : generate random token, use case: oath2
func RandToken() string {
	b := make([]byte, 32)
	rand.Read(b)
	return base64.StdEncoding.EncodeToString(b)
}
