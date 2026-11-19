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
