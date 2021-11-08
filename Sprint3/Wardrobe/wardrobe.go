package wardrobe

import (
	"io"
	"bufio"
	"fmt"
)

func Wardrobe(r io.Reader, w io.Writer){
	reader := bufio.NewReader(r)
	writer := bufio.NewWriter(w)
	fmt.Println(reader, writer)
}