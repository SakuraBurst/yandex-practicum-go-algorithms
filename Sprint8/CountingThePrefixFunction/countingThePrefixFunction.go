package countingThePrefixFunction

import (
	"bufio"
	"fmt"
	"io"
)

func CountingThePrefixFunction(r io.Reader, w io.Writer) {
	reader := bufio.NewReader(r)
	//writer := bufio.NewWriter(w)
	fmt.Println(reader)
}
