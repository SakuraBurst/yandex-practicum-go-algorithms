package numeralPath

import (
	"io"
	"bufio"
	"fmt"
)

func NumeralPath(r io.Reader, w io.Writer){
	reader := bufio.NewReader(r)
	//writer := bufio.NewWriter(w)
	fmt.Println(reader)
}