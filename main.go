package main

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"log"
	"os"
)

func main() {
	path := "random2.bin"

	buf, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err = buf.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	r := bufio.NewReader(buf)

	i := 0
	for i < 1000 {
		var n uint8
		err := binary.Read(r, binary.LittleEndian, &n)
		if err != nil {
			fmt.Println("Error reading file:", err)
			return
		}
		fmt.Print(fmt.Sprintf("%v,", n))
		i += 1
	}
}
