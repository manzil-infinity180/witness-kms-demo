package main

import (
	"github.com/common-nighthawk/go-figure"
)

func Hello() string {
	return "Hello Witness!"
}

func main() {
	myFigure := figure.NewFigure("Hello Witness!", "", true)
	myFigure.Print()
}
