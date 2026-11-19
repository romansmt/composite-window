# Prompt

This is the single prompt that reproduces this project exactly as it currently stands.

```
Implement the Composite design pattern in Go, based on this UML structure (as drawn on
a lecture blackboard):

- Interface `Item` with methods `Price() float64` and `Display(indent string)`
- `Window` (Composite): implements `Item`, holds `[]Item` children (1..* relationship),
  `Add(item Item)` appends a child, `Price()` recursively sums all children's prices,
  `Display(indent)` prints itself then calls `Display` on each child with a deeper indent
- `Glass` and `Frame` (Leaves): implement `Item`, have a fixed price, no children,
  `Display()` prints a single line

Project structure (Go module `composite-window`, Go 1.21+, no external dependencies):
- go.mod
- item.go — the `Item` interface
- leaf.go — `Glass` and `Frame`, each with a constructor `NewGlass(name string, price
  float64)` / `NewFrame(name string, price float64)` and a `Name` field
- window.go — `Window`, with `NewWindow(name string)`, `Add`, `Price`, `Display`
- main.go — builds an example tree: "Window 1" containing "Glass 1" (90.00), "Frame 1"
  (35.00), and a nested "Window 2" containing "Glass 2" (40.00) and "Frame 2" (15.00);
  prints the tree via `Display("")` and the total price
- item_test.go — unit tests for `Price()` on a single leaf, a one-level composite, and a
  nested composite
- .gitignore — ignore built binaries (*.exe)
- README.md — project overview, a Mermaid UML class diagram of Item/Window/Glass/Frame,
  file structure, a "How it works" section explaining the Component/Leaf/Composite roles
  and the Price()/Display() recursion, an "Example" section with the exact tree and
  console output, Go requirements, test/run commands (go vet/go build/go test -v/go run),
  and the full Go source code embedded in code blocks

Naming rules: use only names derived directly from the spec (Item, Window, Glass, Frame)
for instances — no invented flavor text. Number duplicate instances as "Window 1"/"Window 2",
"Glass 1"/"Glass 2", "Frame 1"/"Frame 2". Code must be idiomatic Go; add comments only
where the pattern role (Component/Composite/Leaf) isn't obvious from naming alone.
```
