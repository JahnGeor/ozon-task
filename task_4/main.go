package main

import (
	"bufio"
	"io"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)

	task(in, out)
}

func task(in io.Reader, out io.Writer) {

}
