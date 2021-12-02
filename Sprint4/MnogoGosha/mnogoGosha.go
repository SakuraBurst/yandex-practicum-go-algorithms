package mnogoGosha

import (
	"io"
	"bufio"
	"fmt"
)

func MnogoGosha(r io.Reader, w io.Writer){
	reader := bufio.NewReader(r)
	//writer := bufio.NewWriter(w)
	fmt.Println(reader)
}