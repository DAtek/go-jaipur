package app

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

const clearScreenString = "\033[2J\033[H"

func input(reader *bufio.Reader, writer io.Writer, prompt string) string {
	fmt.Fprint(writer, prompt)
	line, _ := reader.ReadString('\n')
	return strings.Trim(line, "\n")
}
