package main

import (
	"github.com/bassaer/go-lib"
)

func main() {
	m := mylib.NewMessage("Hello")
	m.Send()
}
