package perimeterOfTriangle

import (
	"io"
	"bufio"
	"fmt"
)

func PerimeterOfTriangle(r io.Reader, w io.Writer){
	reader := bufio.NewReader(r)
	writer := bufio.NewWriter(w)
	fmt.Println(reader, writer)
}