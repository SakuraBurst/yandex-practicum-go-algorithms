package insertingRows

import (
	"bufio"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

type InsertedSubStr struct {
	subStr string
	index  int
}

type InsertedSubStrs []InsertedSubStr

func (iss InsertedSubStrs) Len() int {
	return len(iss)
}

func (iss InsertedSubStrs) Less(i, j int) bool {
	return iss[i].index < iss[j].index
}

func (iss InsertedSubStrs) Swap(i, j int) {
	iss[i], iss[j] = iss[j], iss[i]
}

func InsertingRows(r io.Reader, w io.Writer) {
	reader := bufio.NewReader(r)

	s, _ := reader.ReadString('\n')
	s = strings.TrimSpace(s)
	ns, _ := reader.ReadString('\n')
	n, _ := strconv.Atoi(strings.TrimSpace(ns))
	insertedSubStrs := make(InsertedSubStrs, 0, n)
	for i := 0; i < n; i++ {
		subStrInfo, _ := reader.ReadString('\n')
		info := strings.Fields(subStrInfo)
		index, _ := strconv.Atoi(info[1])
		insertedSubStrs = append(insertedSubStrs, InsertedSubStr{
			subStr: info[0],
			index:  index,
		})
	}
	sort.Sort(insertedSubStrs)
	insertedStringIndex := 0
	builder := strings.Builder{}
	i := 0
	for i < len(s) {
		if insertedStringIndex < n && i == insertedSubStrs[insertedStringIndex].index {
			builder.WriteString(insertedSubStrs[insertedStringIndex].subStr)
			insertedStringIndex++
		} else {
			builder.WriteByte(s[i])
			i++
		}

	}
	if insertedSubStrs[n-1].index == len(s) {
		builder.WriteString(insertedSubStrs[n-1].subStr)
	}
	io.WriteString(w, builder.String())
	//fmt.Print(n, s, insertedSubStrs)
}

func main() {
	InsertingRows(os.Stdin, os.Stdout)
}
