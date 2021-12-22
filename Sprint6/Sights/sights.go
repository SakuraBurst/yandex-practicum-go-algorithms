package sights

import (
	"io"
	"bufio"
	"fmt"
)

func Sights(r io.Reader, w io.Writer){
	reader := bufio.NewReader(r)
	//writer := bufio.NewWriter(w)
	fmt.Println(reader)
}