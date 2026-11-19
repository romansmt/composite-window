package main

import "fmt"

func main() {
	subWindow := NewWindow("Kippfenster")
	subWindow.Add(NewGlass("Scheibe klein", 40.0))
	subWindow.Add(NewFrame("Rahmen klein", 15.0))

	mainWindow := NewWindow("Hauptfenster")
	mainWindow.Add(NewGlass("Scheibe groß", 90.0))
	mainWindow.Add(NewFrame("Rahmen groß", 35.0))
	mainWindow.Add(subWindow)

	mainWindow.Display("")
	fmt.Printf("\nGesamtpreis: %.2f\n", mainWindow.Price())
}
