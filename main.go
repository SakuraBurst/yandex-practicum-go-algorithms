package main

import (
	"bytes"
	"log"
	"os"

	nearestzero "github.com/SakuraBurst/yandex-practicum-go-algorithms/Sprint1FinalEx/NearestZero"
)

func main() {
	// createFilesForNextSprint("Sprint1")
	nearestzero.NearestZero()
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
