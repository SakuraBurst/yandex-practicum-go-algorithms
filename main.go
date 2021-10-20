package main

import (
	"bytes"
	"fmt"
	"log"
	"os"

	bracketsequence "github.com/SakuraBurst/yandex-practicum-go-algorithms/Sprint2/BracketSequence"
	"golang.org/x/net/html"
)

func main() {
	bracketsequence.BracketSequence()
	// buf := bytes.NewBuffer([]byte{})
	// buf.WriteString(`<ul class="tabs-menu tabs-menu_theme_normal tabs-menu_layout_vert tabs-menu_size_m tabs-menu_role_problems inline-block i-bem tabs-menu_js_inited" data-bem="{&quot;tabs-menu&quot;:{}}" role="menu"><li class="tabs-menu__tab tabs-menu__tab_first_yes"><a class="link" href="/contest/22779/problems/A/"><span class="tabs-menu__tab-content-text">A. Мониторинг</span></a></li><li class="tabs-menu__tab tabs-menu__tab_active_yes"><a class="link" href="/contest/22779/problems/B/"><span class="tabs-menu__tab-content-text">B. Список дел</span></a></li><li class="tabs-menu__tab"><a class="link" href="/contest/22779/problems/C/"><span class="tabs-menu__tab-content-text">C. Нелюбимое дело</span></a></li><li class="tabs-menu__tab"><a class="link" href="/contest/22779/problems/D/"><span class="tabs-menu__tab-content-text">D. Заботливая мама</span></a></li><li class="tabs-menu__tab"><a class="link" href="/contest/22779/problems/E/"><span class="tabs-menu__tab-content-text">E. Всё наоборот</span></a></li><li class="tabs-menu__tab"><a class="link" href="/contest/22779/problems/F/"><span class="tabs-menu__tab-content-text">F. Стек - Max</span></a></li><li class="tabs-menu__tab"><a class="link" href="/contest/22779/problems/G/"><span class="tabs-menu__tab-content-text">G. Стек - MaxEffective</span></a></li><li class="tabs-menu__tab"><a class="link" href="/contest/22779/problems/H/"><span class="tabs-menu__tab-content-text">H. Скобочная последовательность</span></a></li><li class="tabs-menu__tab"><a class="link" href="/contest/22779/problems/I/"><span class="tabs-menu__tab-content-text">I. Ограниченная очередь</span></a></li><li class="tabs-menu__tab"><a class="link" href="/contest/22779/problems/J/"><span class="tabs-menu__tab-content-text">J. Списочная очередь</span></a></li><li class="tabs-menu__tab"><a class="link" href="/contest/22779/problems/K/"><span class="tabs-menu__tab-content-text">K. Рекурсивные числа Фибоначчи</span></a></li><li class="tabs-menu__tab"><a class="link" href="/contest/22779/problems/L/"><span class="tabs-menu__tab-content-text">L. Фибоначчи по модулю</span></a></li></ul>`)
	// mainNode, err := html.Parse(buf)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// CreateAllDirectoriesAndFilesFromHtmlList(mainNode, "Sprint2")
	// createFilesForNextSprint("Sprint2")
}

func createFilesForNextSprint(path string) {
	dirs, err := os.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}
	for _, v := range dirs {
		name := v.Name()
		f, err := os.Create(path + "/" + name + "/README.md")
		if err != nil {
			log.Fatal(err)
		}
		g, err := os.Create(path + "/" + name + "/" + string(bytes.ToLower([]byte(name[:1]))) + string([]byte(name[1:])) + ".go")

		if err != nil {
			log.Fatal(err)
		}
		err = f.Close()
		if err != nil {
			log.Fatal(err)
		}
		err = g.Close()
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
