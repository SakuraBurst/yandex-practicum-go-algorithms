package crib

import (
	"bufio"
	"errors"
	"io"
	"log"
	"os"
	"strconv"
	"unicode"
)

type Node struct {
	IsBlack bool
	Tree    PrefixTree
	IsRoot  bool
}

type PrefixTree map[string]*Node

type DP [][]*Node

func Crib(r io.Reader, w io.Writer) {
	reader := bufio.NewReader(r)
	s, _ := reader.ReadString('\n')
	s = s[:len(s)-1]
	nString, _ := reader.ReadString('\n')
	n, _ := strconv.Atoi(nString[:len(nString)-1])
	prefixesRoot := &Node{
		IsBlack: false,
		Tree:    PrefixTree{},
		IsRoot:  true,
	}
	for i := 0; i < n; i++ {
		p, _ := reader.ReadString('\n')
		if unicode.IsSpace(rune(p[len(p)-1])) {
			addPrefix(prefixesRoot, p[:len(p)-1])
		} else {
			addPrefix(prefixesRoot, p)
		}

	}
	if prefixesRoot.Tree[string(s[0])] == nil {
		io.WriteString(w, "NO")
		return
	}
	dp := make(DP, len(s))
	dp[0] = append(dp[0], prefixesRoot.Tree[string(s[0])])
	for i := 1; i < len(s); i++ {
		c := string(s[i])
		var err error
		var isSomeBlack bool
		dp[i], isSomeBlack, err = findNodes(dp[i-1], c)

		if isSomeBlack && prefixesRoot.Tree[c] != nil {
			dp[i] = append(dp[i], prefixesRoot.Tree[c])
		}
		if len(dp[i]) == 0 {
			log.Println(err)
			break
		}
	}
	//log.Println(dp)
	if someIsBlack(dp[len(s)-1]) {
		io.WriteString(w, "YES")
	} else {
		io.WriteString(w, "NO")
	}
}

func findNodes(n []*Node, char string) ([]*Node, bool, error) {
	var res []*Node
	isSomeBlack := false
	for _, v := range n {
		if v.IsBlack {
			isSomeBlack = true
		}
		if v.Tree[char] != nil {
			res = append(res, v.Tree[char])
		}
	}
	if len(res) == 0 {
		return res, isSomeBlack, errors.New("there is no such nodes")
	}
	return res, isSomeBlack, nil
}

func someIsBlack(n []*Node) bool {
	for _, v := range n {
		if v.IsBlack {
			return true
		}
	}
	return false
}

func addPrefix(root *Node, s string) {
	for _, r := range s {
		_, ok := root.Tree[string(r)]
		if !ok {
			root.Tree[string(r)] = &Node{
				IsBlack: false,
				Tree:    PrefixTree{},
			}
		}
		root = root.Tree[string(r)]
	}
	root.IsBlack = true
}

func main() {
	Crib(os.Stdin, os.Stdout)
}
