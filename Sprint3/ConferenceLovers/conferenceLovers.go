package conferenceLovers

import (
	"io"
	"bufio"
	"fmt"
)

func ConferenceLovers(r io.Reader, w io.Writer){
	reader := bufio.NewReader(r)
	writer := bufio.NewWriter(w)
	fmt.Println(reader, writer)
}