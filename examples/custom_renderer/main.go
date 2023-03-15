package main

import (
	"fmt"

	"github.com/charmbracelet/glamour"
)

func main() {
	in := `# Custom Renderer

Word-wrapping will occur when lines exceed the limit of 40 characters.

[Hyperlink support.](https://www.youtube.com/watch?v=dQw4w9WgXcQ)
`

	r, _ := glamour.NewTermRenderer(
		glamour.WithStandardStyle("dark"),
		glamour.WithWordWrap(40),
		glamour.WithHyperlinks(true),
	)

	out, _ := r.Render(in)
	fmt.Print(out)
}
