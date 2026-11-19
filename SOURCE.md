# Source Code

Full Go source of the [composite-window](README.md) project.

## go.mod

```go
module composite-window

go 1.21
```

## item.go

```go
package main

// Item is the Component in the Composite pattern: both leaves (Glass,
// Frame) and composites (Window) implement it, so callers treat single
// parts and assemblies uniformly.
type Item interface {
	Price() float64
	Display(indent string)
}
```

## leaf.go

```go
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
```

## window.go

```go
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
```

## main.go

```go
package main

import "fmt"

func main() {
	subWindow := NewWindow("Window 2")
	subWindow.Add(NewGlass("Glass 2", 40.0))
	subWindow.Add(NewFrame("Frame 2", 15.0))

	mainWindow := NewWindow("Window 1")
	mainWindow.Add(NewGlass("Glass 1", 90.0))
	mainWindow.Add(NewFrame("Frame 1", 35.0))
	mainWindow.Add(subWindow)

	mainWindow.Display("")
	fmt.Printf("\nTotal price: %.2f\n", mainWindow.Price())
}
```

## item_test.go

```go
package main

import "testing"

func TestLeafPrice(t *testing.T) {
	g := NewGlass("Glass", 40.0)
	if got := g.Price(); got != 40.0 {
		t.Errorf("Glass.Price() = %.2f, want 40.00", got)
	}
}

func TestWindowPriceSumsChildren(t *testing.T) {
	w := NewWindow("Window")
	w.Add(NewGlass("Glass", 40.0))
	w.Add(NewFrame("Frame", 15.0))

	if got, want := w.Price(), 55.0; got != want {
		t.Errorf("Window.Price() = %.2f, want %.2f", got, want)
	}
}

func TestNestedWindowPrice(t *testing.T) {
	sub := NewWindow("Window 2")
	sub.Add(NewGlass("Glass 2", 40.0))
	sub.Add(NewFrame("Frame 2", 15.0))

	main := NewWindow("Window 1")
	main.Add(NewGlass("Glass 1", 90.0))
	main.Add(NewFrame("Frame 1", 35.0))
	main.Add(sub)

	want := 40.0 + 15.0 + 90.0 + 35.0
	if got := main.Price(); got != want {
		t.Errorf("nested Window.Price() = %.2f, want %.2f", got, want)
	}
}
```
