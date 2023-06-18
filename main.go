package main

import (
	"fmt"
	"jacobjedwards/templater/templater"
)

func main() {
	data := make(map[string]string)
	data["test"] = "success"
	arr := templater.Compile("test test tesiajdga {{ test }} asfj", data)
	fmt.Println(arr)
}
