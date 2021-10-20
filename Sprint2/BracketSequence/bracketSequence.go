package bracketsequence

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

func BracketSequence() {
	stdReader := bufio.NewReader(os.Stdin)
	str, _ := stdReader.ReadString('\n')
	str = strings.TrimSpace(str)
	stack := make(Stack, 0, len(str))
	for i := 0; i < len(str); i++ {
		s := string(str[i])
		if isOpenedBracket(s) {
			stack.push(s)
		} else {
			if err := stack.popItem(toOpenedBracket(s)); err != nil {
				fmt.Println("False")
				return
			}
		}
	}
	if stack.isEmpty() {
		fmt.Println("True")
		return
	} else {
		fmt.Println("False")
	}

}

func toOpenedBracket(b string) string {
	switch b {
	case "}":
		return "{"
	case ")":
		return "("
	case "]":
		return "["
	default:
		return ""
	}
}

func isOpenedBracket(b string) bool {
	switch b {
	case "{":
		return true
	case "(":
		return true
	case "[":
		return true
	default:
		return false
	}
}

type Stack []string

// func (b *Stack) pop() error {
// 	if b.isEmpty() {
// 		return errors.New("stack is empty")
// 	}
// 	*b = (*b)[:len(*b)-1]
// 	return nil
// }

func (b *Stack) popItem(i string) error {
	if b.isEmpty() {
		return errors.New("stack is empty")
	}
	deleteItem := (*b)[len(*b)-1]
	if deleteItem != i {
		return errors.New("items Is Not Equal " + deleteItem + " !== " + i)
	}
	*b = (*b)[:len(*b)-1]
	return nil
}

func (b *Stack) push(s string) {
	*b = append(*b, s)
}

func (b Stack) isEmpty() bool {
	return len(b) == 0
}
