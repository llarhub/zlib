package main

import (
	"fmt"

	"github.com/llarhub/zlib"
)

func main() {
	fmt.Printf("%08x\n", zlib.Crc32ZString(0, "Hello world"))
}
