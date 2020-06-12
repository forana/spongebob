package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// Spongebob SpOnGeBoBs text
func Spongebob(text string) string {
	// do this the lazy way
	b := ""
	lower := true
	for _, c := range text {
		if lower {
			b += strings.ToLower(string(c))
		} else {
			b += strings.ToUpper(string(c))
		}
		lower = !lower
	}
	return b
}

func main() {
	s := ""
	flag.Parse()
	args := flag.Args()
	if len(args) < 1 {
		b, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			panic(err)
		}
		s = string(b)
	} else {
		s = args[0]
	}
	fmt.Print(Spongebob(s))
}
