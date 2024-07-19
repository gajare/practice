package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("key :", key())
}

func key() string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	buf := make([]byte, 16)
	for i := range buf {
		buf[i] = byte(r.Intn(256))
	}
	return fmt.Sprintf("%x\n", buf)
}
