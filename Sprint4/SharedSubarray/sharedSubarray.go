package sharedSubarray

import (
	"io"
	"bufio"
	"fmt"
)

func SharedSubarray(r io.Reader, w io.Writer){
	reader := bufio.NewReader(r)
	//writer := bufio.NewWriter(w)
	fmt.Println(reader)
}