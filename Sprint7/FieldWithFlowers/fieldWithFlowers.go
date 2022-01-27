package fieldWithFlowers

import (
	"bufio"
	"fmt"
	"io"
)

func FieldWithFlowers(r io.Reader, w io.Writer) {
	reader := bufio.NewReader(r)
	//writer := bufio.NewWriter(w)
	fmt.Println(reader)
}
