package goshaInTheRestaurant

import (
	"bufio"
	"fmt"
	"io"
)

func GoshaInTheRestaurant(r io.Reader, w io.Writer) {
	reader := bufio.NewReader(r)
	//writer := bufio.NewWriter(w)
	fmt.Println(reader)
}
