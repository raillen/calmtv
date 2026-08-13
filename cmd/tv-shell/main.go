package main

import (
	"fmt"
	"os"

	"github.com/gotk3/gotk3/gtk"

	"github.com/raillen/calmtv/internal/shell"
)

func main() {
	gtk.Init(nil)
	app, err := shell.New()
	if err != nil {
		fmt.Fprintf(os.Stderr, "tv-shell: %v\n", err)
		os.Exit(1)
	}
	app.Run()
}
