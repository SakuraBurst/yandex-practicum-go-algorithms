package frequentWord

import (
	"bufio"
	"fmt"
	"io"
)

func FrequentWord(r io.Reader, w io.Writer) {
	reader := bufio.NewReader(r)
	//writer := bufio.NewWriter(w)
	fmt.Println(reader)
}