package main

import (
	"bufio"
	"log"
	"os"

	"github.com/yowcow/urljson/parser"
)

func main() {
	r := os.Stdin
	w := os.Stdout

	s := bufio.NewScanner(r)
	for s.Scan() {
		url := s.Text()
		u, err := parser.Parse(url)
		if err != nil {
			log.Printf("failed parsing url '%s': %s", url, err)
			continue
		}
		u.WriteJson(w)
	}
}
