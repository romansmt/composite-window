package main

import "fmt"

// Window is the Composite: it holds 1..* child Items (leaves or nested
// Windows) and satisfies the Item interface itself, so a Window can be
// a child of another Window.
type Window struct {
	Name     string
	children []Item
}

func NewWindow(name string) *Window {
	return &Window{Name: name}
}

func (w *Window) Add(item Item) {
	w.children = append(w.children, item)
}

func (w *Window) Price() float64 {
	total := 0.0
	for _, child := range w.children {
		total += child.Price()
	}
	return total
}

func (w *Window) Display(indent string) {
	fmt.Printf("%s+ %s (Window): %.2f\n", indent, w.Name, w.Price())
	for _, child := range w.children {
		child.Display(indent + "  ")
	}
}
