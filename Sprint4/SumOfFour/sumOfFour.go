package sumOfFour

import (
	"io"
	"bufio"
	"fmt"
)

func SumOfFour(r io.Reader, w io.Writer){
	reader := bufio.NewReader(r)
	//writer := bufio.NewWriter(w)
	fmt.Println(reader)
}