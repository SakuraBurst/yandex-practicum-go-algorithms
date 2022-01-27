package horoscopes

import (
	"bufio"
	"fmt"
	"io"
)

func Horoscopes(r io.Reader, w io.Writer) {
	reader := bufio.NewReader(r)
	//writer := bufio.NewWriter(w)
	fmt.Println(reader)
}
