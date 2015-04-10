package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/yosssi/gohtml"
)

func parseFile(file string) {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
		return
	}

	formatted := gohtml.Format(string(data))

	err = ioutil.WriteFile(file, []byte(formatted), 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func parseStdin() {
	data, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
		return
	}

	formatted := gohtml.Format(string(data))

	fmt.Println(formatted)
}

func main() {
	flag.Parse()
	args := flag.Args()

	if len(args) > 0 {
		for _, file := range args {
			parseFile(file)
		}
	} else {
		parseStdin()
	}
}
