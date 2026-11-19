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
