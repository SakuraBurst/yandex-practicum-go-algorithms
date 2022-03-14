package main

import (
	"bytes"
	"fmt"
	"golang.org/x/net/html"
	"log"
	"os"
	"strings"
	"text/template"
)

type Node struct {
	val  int
	next *Node
}

func main() {
	// buf := bytes.NewBuffer([]byte{})
	// buf.WriteString(`<ul class="tabs-menu tabs-menu_theme_normal tabs-menu_layout_vert tabs-menu_size_m tabs-menu_role_problems inline-block i-bem tabs-menu_js_inited" data-bem="{&quot;tabs-menu&quot;:{}}" role="menu"><li class="tabs-menu__tab tabs-menu__tab_first_yes"><a class="link" href="/contest/22779/problems/A/"><span class="tabs-menu__tab-content-text">A. Мониторинг</span></a></li><li class="tabs-menu__tab tabs-menu__tab_active_yes"><a class="link" href="/contest/22779/problems/B/"><span class="tabs-menu__tab-content-text">B. Список дел</span></a></li><li class="tabs-menu__tab"><a class="link" href="/contest/22779/problems/C/"><span class="tabs-menu__tab-content-text">C. Нелюбимое дело</span></a></li><li class="tabs-menu__tab"><a class="link" href="/contest/22779/problems/D/"><span class="tabs-menu__tab-content-text">D. Заботливая мама</span></a></li><li class="tabs-menu__tab"><a class="link" href="/contest/22779/problems/E/"><span class="tabs-menu__tab-content-text">E. Всё наоборот</span></a></li><li class="tabs-menu__tab"><a class="link" href="/contest/22779/problems/F/"><span class="tabs-menu__tab-content-text">F. Стек - Max</span></a></li><li class="tabs-menu__tab"><a class="link" href="/contest/22779/problems/G/"><span class="tabs-menu__tab-content-text">G. Стек - MaxEffective</span></a></li><li class="tabs-menu__tab"><a class="link" href="/contest/22779/problems/H/"><span class="tabs-menu__tab-content-text">H. Скобочная последовательность</span></a></li><li class="tabs-menu__tab"><a class="link" href="/contest/22779/problems/I/"><span class="tabs-menu__tab-content-text">I. Ограниченная очередь</span></a></li><li class="tabs-menu__tab"><a class="link" href="/contest/22779/problems/J/"><span class="tabs-menu__tab-content-text">J. Списочная очередь</span></a></li><li class="tabs-menu__tab"><a class="link" href="/contest/22779/problems/K/"><span class="tabs-menu__tab-content-text">K. Рекурсивные числа Фибоначчи</span></a></li><li class="tabs-menu__tab"><a class="link" href="/contest/22779/problems/L/"><span class="tabs-menu__tab-content-text">L. Фибоначчи по модулю</span></a></li></ul>`)
	// mainNode, err := html.Parse(buf)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// CreateAllDirectoriesAndFilesFromHtmlList(mainNode, "Sprint2")
	//bigBigInt := 1 << 30
	//aBuilder := strings.Builder{}
	//bBuilder := strings.Builder{}
	//rand.Seed(time.Now().Unix())
	//log.Println("string builds are started")
	//for i := 0; i < bigBigInt; i++ {
	//	a := uint8(rand.Intn(122-65) + 65)
	//	aBuilder.WriteByte(a)
	//	bBuilder.WriteByte(a)
	//}
	//log.Println("strings build are ended with length", bigBigInt)
	//log.Println("strings compare are started")
	//a := aBuilder.String()
	//b := bBuilder.String()
	//res := a == b
	//log.Println("strings compare are ended")
	//log.Println("result ", res)
	//log.Println("custom check are started")
	//for i := 0; i < bigBigInt; i++ {
	//	if a[i] != b[i] {
	//		log.Println("custom check is ended")
	//		log.Println("result false")
	//	}
	//}
	//log.Println("custom check is ended")
	//log.Println("result true")
	createFilesForNextSprint("Sprint8")
}

func reverseLinkedList(n, previesNode *Node) *Node {
	if n.next == nil {
		n.next = previesNode
		return n
	}
	oldNext := n.next
	n.next = previesNode
	return reverseLinkedList(oldNext, n)
}

func createFilesForNextSprint(path string) {
	dirs, err := os.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}
	readMeExample, err := os.ReadFile("README-example.md")
	if err != nil {
		log.Fatal(err)
	}
	for _, v := range dirs {
		name := v.Name()
		r, err := os.Create(path + "/" + name + "/README.md")
		if err != nil {
			log.Fatal(err)
		}
		lowerName := string(bytes.ToLower([]byte(name[:1]))) + string([]byte(name[1:]))
		f, err := os.Create(path + "/" + name + "/" + lowerName + ".go")
		if err != nil {
			log.Fatal(err)
		}
		t, err := os.Create(path + "/" + name + "/" + lowerName + "_test.go")
		if err != nil {
			log.Fatal(err)
		}
		n, err := r.Write(readMeExample)
		if n == 0 || err != nil {
			log.Fatal(err)
		}
		err = r.Close()
		if err != nil {
			log.Fatal(err)
		}

		mainFileTemplate, err := template.New("file").Parse(`package {{.lowerName}}

import (
	"io"
	"bufio"
	"fmt"
)

func {{.name}}(r io.Reader, w io.Writer){
	reader := bufio.NewReader(r)
	//writer := bufio.NewWriter(w)
	fmt.Println(reader)
}`)

		if err != nil {
			log.Fatal(err)
		}

		testFileTemplate, err := template.New("Test").Parse(`package {{.lowerName}}_test

import (
	"bytes"
	"io"
	"strings"
	"testing"

	{{.lowerName}} "github.com/SakuraBurst/yandex-practicum-go-algorithms/{{.path}}/{{.name}}"
)

type Test struct {
	inputData  io.Reader
	outputData string
}

func Test{{.name}}(t *testing.T){
	tests := []Test{{"{{"}}strings.NewReader(""), ""}}
	for _, v := range tests {
		buf := bytes.NewBuffer(nil)
		{{.lowerName}}.{{.name}}(v.inputData, buf)
		if buf.String() != v.outputData {
			t.Errorf("\nexpected:\n%s\ngot:\n%s", v.outputData, buf.String())
		}
	}
}`)
		if err != nil {
			log.Fatal(err)
		}
		err = mainFileTemplate.Execute(f, map[string]string{"lowerName": lowerName, "name": name})
		if err != nil {
			log.Fatal(err)
		}
		err = testFileTemplate.Execute(t, map[string]string{"lowerName": lowerName, "name": name, "path": path})
		if err != nil {
			log.Fatal(err)
		}
		err = f.Close()
		if err != nil {
			log.Fatal(err)
		}
	}
}

func CreateAllDirectoriesAndFilesFromHtmlList(node *html.Node, dir string) {
	if node == nil {
		return
	}
	if node.Type == html.TextNode {
		os.Mkdir((dir + "/" + node.Data), os.ModePerm)
		fmt.Println(node.Data)
	}
	CreateAllDirectoriesAndFilesFromHtmlList(node.FirstChild, dir)
	CreateAllDirectoriesAndFilesFromHtmlList(node.NextSibling, dir)

}

func stringsCompare(a, b string) bool {
	return strings.Compare(a, b) == 0
}

func operatorCompare(a, b string) bool {
	return a == b
}
