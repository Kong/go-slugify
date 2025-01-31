package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/kong/go-slugify"
	"github.com/mattn/go-isatty"
)

func main() {
	version := flag.Bool("V", false, "Output version info")
	flag.Parse()
	if *version {
		v := slugify.Version()
		fmt.Printf("slugify %s\n", v)
		os.Exit(0)
	}

	textSlice := flag.Args()
	stdin := []byte{}
	if !isatty.IsTerminal(os.Stdin.Fd()) {
		stdin, _ = ioutil.ReadAll(os.Stdin)
	}
	if len(stdin) > 0 {
		textSlice = append(textSlice, string(stdin))
	}

	if len(textSlice) == 0 {
		fmt.Println("Usage: slugify STRING")
		os.Exit(1)
	}

	s := strings.Join(textSlice, " ")
	slugifier := slugify.NewSlugifier()
	ret := slugifier.Slugify(s)
	fmt.Println(ret)
}
