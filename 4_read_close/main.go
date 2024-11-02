package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

type ReadCloser interface {
	Read(b []byte) (n int, err error)
	Close() error
}

func ReadAndClose(r ReadCloser, buf []byte) (n int, err error) {
	for len(buf) > 0 && err == nil {
		var nr int
		nr, err = r.Read(buf)
		n += nr
		buf = buf[nr:]
	}
	r.Close()
	return
}

func main() {
	file, err := os.Open("example.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	buf := make([]byte, 1024)
	n, err := ReadAndClose(file, buf)
	if err != nil && err != io.EOF {
		log.Fatal(err)
	}

	fmt.Printf("Read %d bytes: %s\n", n, buf[:n])
}

// https://research.swtch.com/interfaces
