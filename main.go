package main

import (
	"io"
	"os"
)

const octo = 8

func main() {
	goggles := make([]io.Writer, octo)

	for i := range goggles {
		goggles[i] = os.Stdout
	}

	writer := io.MultiWriter(goggles...)

	io.Copy(writer, os.Stdin)
}
