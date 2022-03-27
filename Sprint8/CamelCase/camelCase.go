package camelCase

import (
	"bufio"
	"errors"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
	"unicode"
)

type Node struct {
	IsBlack bool
	Words   []string
	Tree    PrefixTree
}

type PrefixTree map[string]*Node

type ClassName string

func (n ClassName) Abbreviation() string {
	builder := strings.Builder{}
	for _, i := range n {
		if unicode.IsUpper(i) {
			builder.WriteRune(i)
		}
	}
	return builder.String()
}

func CamelCase(r io.Reader, w io.Writer) {
	prefixTreeRoot := Node{
		IsBlack: false,
		Tree:    PrefixTree{},
	}
	scanner := bufio.NewScanner(r)
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	for i := 0; i < n; i++ {
		scanner.Scan()
		addString(&prefixTreeRoot, ClassName(scanner.Text()))
	}
	scanner.Scan()
	m, _ := strconv.Atoi(scanner.Text())
	writer := bufio.NewWriter(w)
	for i := 0; i < m; i++ {
		scanner.Scan()
		r, err := findNode(&prefixTreeRoot, scanner.Text())
		sort.Strings(r)
		if err == nil && len(r) > 0 {
			for _, s := range r {
				writer.WriteString(s)
				writer.WriteByte('\n')
			}
		}
	}
	writer.Flush()
}

func addString(root *Node, s ClassName) {
	currentNode := root
	for _, i := range s.Abbreviation() {
		char := string(i)
		if currentNode.Tree == nil {
			currentNode.Tree = PrefixTree{}
		}
		_, ok := currentNode.Tree[char]
		if !ok {
			currentNode.Tree[char] = &Node{
				IsBlack: false,
				Tree:    nil,
				Words:   nil,
			}
		}
		currentNode = currentNode.Tree[char]
	}
	currentNode.IsBlack = true
	currentNode.Words = append(currentNode.Words, string(s))
}

func findNode(root *Node, s string) ([]string, error) {
	currentNode := root
	for _, i := range s {
		char := string(i)
		_, ok := currentNode.Tree[char]
		if ok {
			currentNode = currentNode.Tree[char]
		} else {
			return nil, errors.New("there is no such subsequence")
		}
	}
	return findAllNodes(currentNode), nil

}

func findAllNodes(root *Node) []string {
	var result []string = nil
	if root.IsBlack {
		result = root.Words
	}
	if root.Tree == nil {
		return result
	}
	for _, node := range root.Tree {
		result = append(result, findAllNodes(node)...)
	}
	return result
}

func main() {
	CamelCase(os.Stdin, os.Stdout)
}
