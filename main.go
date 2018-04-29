package main

import (
	"fmt"

	"honnef.co/go/js/dom"
)

func main() {
	body := dom.GetWindow().Document().GetElementsByTagName("body")[0]
	body.SetInnerHTML("Hello, HTML!")
	fmt.Println("Hello, console!")
}
