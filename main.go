package main

import (
	"flag"
)

type (
	HttpRequest struct {
		Method string
		Path   string
	}
)

func main() {
	flag.Parse()
}
