package globalReplacement

import (
	"bufio"
	"fmt"
	"io"
)

func GlobalReplacement(r io.Reader, w io.Writer) {
	reader := bufio.NewReader(r)
	//writer := bufio.NewWriter(w)
	fmt.Println(reader)
}