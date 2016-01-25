package main

import (
	"flag"
	"fmt"
	"path/filepath"
)

func main() {
	d := filepath.Join("/home", "milin/work")
	fmt.Println(d)
}

func init() {

	line := flag.String("line", "\r\n", "line")
	if *line == "\r\n" {
		fmt.Println("crlf")
	} else {
		fmt.Println("Hello")
	}

}
