package main

import "testing"

func TestLeafPrice(t *testing.T) {
	g := NewGlass("Scheibe", 40.0)
	if got := g.Price(); got != 40.0 {
		t.Errorf("Glass.Price() = %.2f, want 40.00", got)
	}
}

func TestWindowPriceSumsChildren(t *testing.T) {
	w := NewWindow("Fenster")
	w.Add(NewGlass("Scheibe", 40.0))
	w.Add(NewFrame("Rahmen", 15.0))

	if got, want := w.Price(), 55.0; got != want {
		t.Errorf("Window.Price() = %.2f, want %.2f", got, want)
	}
}

func TestNestedWindowPrice(t *testing.T) {
	sub := NewWindow("Kippfenster")
	sub.Add(NewGlass("Scheibe klein", 40.0))
	sub.Add(NewFrame("Rahmen klein", 15.0))

	main := NewWindow("Hauptfenster")
	main.Add(NewGlass("Scheibe groß", 90.0))
	main.Add(NewFrame("Rahmen groß", 35.0))
	main.Add(sub)

	want := 40.0 + 15.0 + 90.0 + 35.0
	if got := main.Price(); got != want {
		t.Errorf("nested Window.Price() = %.2f, want %.2f", got, want)
	}
}
