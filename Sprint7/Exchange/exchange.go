package exchange

import (
	"bufio"
	"fmt"
	"io"
)

func Exchange(r io.Reader, w io.Writer) {
	reader := bufio.NewReader(r)
	//writer := bufio.NewWriter(w)
	fmt.Println(reader)
}
