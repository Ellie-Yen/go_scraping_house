package utils

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func SaveFile(filename string, content string) {
	file, err := os.Create(filename)
	if err != nil {
		fmt.Printf("Error creating file: %v\n", err)
		return
	}
	defer file.Close()

	bytesWritten, err := io.Copy(file, strings.NewReader(content))
	if err != nil {
		fmt.Printf("Error writing to file: %v\n", err)
		return
	}
	fmt.Printf("Successfully saved response to %s (%d bytes written)\n", filename, bytesWritten)
}
