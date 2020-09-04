package main

import (
	"fmt"

	"github.com/dkprog/hello-go-module/morestrings"
	"github.com/google/go-cmp/cmp"
)

func main() {
	fmt.Println(morestrings.ReverseRunes("!dlrow ,olleH"))
	fmt.Println(cmp.Diff("Hello, World!", "Hello, world!"))

}
