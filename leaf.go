package main

import "fmt"

// Glass is a Leaf: a fixed-price part with no children.
type Glass struct {
	Name  string
	price float64
}

func NewGlass(name string, price float64) *Glass {
	return &Glass{Name: name, price: price}
}

func (g *Glass) Price() float64 {
	return g.price
}

func (g *Glass) Display(indent string) {
	fmt.Printf("%s- %s (Glass): %.2f\n", indent, g.Name, g.price)
}

// Frame is a Leaf: a fixed-price part with no children.
type Frame struct {
	Name  string
	price float64
}

func NewFrame(name string, price float64) *Frame {
	return &Frame{Name: name, price: price}
}

func (f *Frame) Price() float64 {
	return f.price
}

func (f *Frame) Display(indent string) {
	fmt.Printf("%s- %s (Frame): %.2f\n", indent, f.Name, f.price)
}
