package anagramGrouping

import (
	"io"
	"bufio"
	"fmt"
)

func AnagramGrouping(r io.Reader, w io.Writer){
	reader := bufio.NewReader(r)
	//writer := bufio.NewWriter(w)
	fmt.Println(reader)
}