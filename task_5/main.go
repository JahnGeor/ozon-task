package main

import (
	"bufio"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)

	taskInteractor(in, out)
}

func taskInteractor(in *bufio.Reader, out *bufio.Writer) {

}
