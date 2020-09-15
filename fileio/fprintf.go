package fileio

import (
	"fmt"
	"os"
)

func fprintfWrite(filename string) {
	destination, err := os.Create(filename)
	if err != nil {
		fmt.Println("os.Create:", err)
		os.Exit(1)
	}
	defer destination.Close()

	fmt.Fprintf(destination, "[%s]: ", filename)
	fmt.Fprintf(destination, "Using fmt.Fprintf in %s\n", filename)
}
