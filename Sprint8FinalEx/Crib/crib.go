/**
id в контесте 66950000 (элитный id 😎)
-- ПРИНЦИП РАБОТЫ --
Привет! Попотел над задачей, сидел 2 дня. В итоге как-то озарение пришло:) Решил через бор и динамическое программирование (после успешного решения почитал свой канал,
увидел, что многие так и решают). Собственно суть в том, что на каждом шаге проверки строки, мы засовываем в массив dp то какие подстроки привели нас к текущему символу.
Вычисляем мы это проверкой, ведут ли какие-то Ноды из dp[i - 1] к текущему символу. Также, для самого первого символа и, если в dp[i - 1] есть ноды, которые являются концами подстрок,
то мы ищем символ в рутовой Ноде Бора. Ответ мы узнаем в dp[len(s) -1] - если какой-то из символов подстрок,
который пришел к концу строки, является концом подстроки, то тогда ответ "Yes", если нет, то "No"


-- ДОКАЗАТЕЛЬСТВО КОРРЕКТНОСТИ --
Я без понятия, как описать это нормально, поэтому опишу весь путь для примера со всеми шагами, чтобы суть хоть немного былая ясна)
Предположим у нас есть строка s := abcdg и подстроки
1 a
2 ab
3 abc
4 abcd
5 g

На первом шаге у нас в dp[0] у нас хранится 'a' Нода, которая ведет в 'b' Ноду как продолжение подстрок 2,3,4 и является концом подстроки 1.
Так как s[0] является концом подстроки 1, мы ищем 'b' в рутовой Ноде, но там ничего нет, поэтому мы кладем только Ноду 'b' в dp[1]

На втором шаге у нас в dp[1] у нас хранится 'b' Нода, которая ведет в 'c' Ноду как продолжение подстрок 3,4 и является концом подстроки 2.
Так как s[1] является концом подстроки 2, мы ищем 'b' в рутовой Ноде, но там ничего нет, поэтому мы кладем только Ноду 'c' в dp[2]

...

На пятом шаге у нас в dp[4] у нас хранится 'd' Нода, которая никуда не ведет и является концом подстроки 4.
Так как s[4] является концом подстроки 4, мы ищем 'g' в рутовой Ноде, и находим! Кладем в dp[5] 'g' Ноду.

'g' Нода является концом подстроки 'g', поэтому строка корректна, выводим "Yes"


-- ВРЕМЕННАЯ СЛОЖНОСТЬ --
Как говорилось в теории сложность формирования бора это O(l), где l — суммарная длина слов во множестве.

Высчитывание результата основанное на динамическом программировании происходит за O(s), где s - длинна строки. В боре мы на каждом шаге ищем только 1 символ, что происходит за O(1).
Максимально к последней итерации в dp[i - 1] может храниться до s Нод
Итого получаем O(l + (s^2))


-- ПРОСТРАНСТВЕННАЯ СЛОЖНОСТЬ --
В силу тех же причин, что и во временной сложности, максимальная пространственная сложность равна, мне кажется, O(l^2)
*/

package crib

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"unicode"
)

type Node struct {
	IsBlack bool
	Tree    PrefixTree
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
		var isSomeBlack bool
		dp[i], isSomeBlack = findNodes(dp[i-1], c)

		if isSomeBlack && prefixesRoot.Tree[c] != nil {
			dp[i] = append(dp[i], prefixesRoot.Tree[c])
		}

		if len(dp[i]) == 0 {
			break
		}
	}

	if someIsBlack(dp[len(s)-1]) {
		io.WriteString(w, "YES")
	} else {
		io.WriteString(w, "NO")
	}
}

func findNodes(n []*Node, char string) ([]*Node, bool) {
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
		return res, isSomeBlack
	}
	return res, isSomeBlack
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
