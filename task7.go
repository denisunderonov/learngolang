package main

import "fmt"

type Shape interface {
	Area() int
}

type Rectangle struct {
	Width  int
	Height int
}

type Square struct {
	Side int
}

func (r Rectangle) Area() int {
	return r.Width * r.Height
}

func (s Square) Area() int {
	return s.Side * s.Side
}

func main() {
	
}