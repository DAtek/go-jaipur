package app

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

const clearScreenString = "\033[2J\033[H"

// const clearScreenString = "\n"

func input(reader io.Reader, writer io.Writer, prompt string) string {
	fmt.Fprint(writer, prompt)
	scanner := bufio.NewScanner(reader)
	scanner.Scan()
	line := scanner.Text()
	return strings.Trim(line, "\n")
}
