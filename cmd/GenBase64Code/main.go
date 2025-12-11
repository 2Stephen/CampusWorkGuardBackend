package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
)

func main() {
	bytes := make([]byte, 64)
	rand.Read(bytes)
	fmt.Println(base64.StdEncoding.EncodeToString(bytes))
}
