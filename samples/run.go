package main

import (
	"bytes"
	"crypto/rand"
	"encoding/base64"
	"fmt"
)

func main() {
	c := 10
	b := make([]byte, c)
	_, err := rand.Read(b)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Println(b)
	// The slice should now contain random bytes instead of only zeroes.
	fmt.Println(bytes.Equal(b, make([]byte, c)))

	fmt.Println(base64.URLEncoding.EncodeToString(b))
}
