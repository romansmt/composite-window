package main

// Item is the Component in the Composite pattern: both leaves (Glass,
// Frame) and composites (Window) implement it, so callers treat single
// parts and assemblies uniformly.
type Item interface {
	Price() float64
	Display(indent string)
}
