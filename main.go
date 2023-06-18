package main

import (
	"fmt"
	"jacobjedwards/templater/templater"
)

func main() {
    arr := templater.Compile("test test tesiajdga {{ test }} asfj")
    fmt.Println(arr)
}
