package bulbs

import (
	"io"
	"bufio"
	"fmt"
)

func Bulbs(r io.Reader, w io.Writer){
	reader := bufio.NewReader(r)
	//writer := bufio.NewWriter(w)
	fmt.Println(reader)
}