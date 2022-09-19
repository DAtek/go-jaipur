package app

import (
	"fmt"
	"io"
)

const clearScreenString = "\033[2J\033[H"

func input(reader io.Reader, writer io.Writer, prompt string) string {
	result := ""
	fmt.Fprint(writer, prompt)
	fmt.Fscanln(reader, &result)
	return result
}
