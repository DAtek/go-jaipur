package app

import (
	"fmt"
	"io"
)

func input(reader io.Reader, writer io.Writer, prompt string) string {
	result := ""
	fmt.Fprint(writer, prompt)
	fmt.Fscanln(reader, &result)
	return result
}
