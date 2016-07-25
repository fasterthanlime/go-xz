package main

import (
	"fmt"
	"io"
	"os"

	"github.com/fasterthanlime/go-xz"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Printf("Usage: %s in out\n", os.Args[0])
		os.Exit(1)
	}
	srcFile := os.Args[1]
	dstFile := os.Args[2]

	f, err := os.Open(srcFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer f.Close()

	xr := xz.NewReader(f)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	of, err := os.Create(dstFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer of.Close()

	wb, err := io.CopyBuffer(of, xr, make([]byte, 4096))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("Wrote %d bytes to %s\n", wb, dstFile)
	os.Exit(0)
}
