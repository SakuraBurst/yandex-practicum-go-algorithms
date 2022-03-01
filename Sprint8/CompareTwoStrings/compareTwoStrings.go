package compareTwoStrings

import (
	"bufio"
	"fmt"
	"io"
)

func CompareTwoStrings(r io.Reader, w io.Writer) {
	reader := bufio.NewReader(r)
	//writer := bufio.NewWriter(w)
	fmt.Println(reader)
}
