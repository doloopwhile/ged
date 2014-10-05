package main

func main() {
	model := &Model{}
	view := NewView(model)

	view.Main()
}
