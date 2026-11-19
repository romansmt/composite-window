# composite-window

Go implementation of the Composite design pattern, modeled on the lecture
example: a `Window` is composed of `Glass` and `Frame` parts, and can itself
contain nested `Window`s.

## Structure

- `item.go` — `Item` interface (`Price() float64`, `Display(indent string)`), the Component role
- `leaf.go` — `Glass`, `Frame`: Leaf implementations with a fixed price
- `window.go` — `Window`: Composite implementation holding `[]Item` children
- `main.go` — builds an example window tree and prints its structure and total price
- `item_test.go` — tests for `Price()` on a single leaf, a composite, and a nested composite

## Requirements

- Go 1.21+

## Commands

```bash
go vet ./...
```
```bash
go build ./...
```
```bash
go test ./... -v
```
```bash
go run .
```
