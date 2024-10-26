package main

import (
	"fmt"
	"os"

	"github.com/nadedan/console"
)

func main() {

	c := console.MakeConsole(os.Stdin, os.Stdout)

	c.Printf("Printing some lines\n")
	c.Printf("Printing some lines\n")
	c.Printf("Printing some lines\n")
	c.Printf("Printing some lines\n")
	c.Printf("Printing some lines\n")
	c.Printf("Printing some lines\n")
	c.Printf("Printing some lines\n")
	c.Printf("Printing last line\n")

	x := 0

	cancel := c.RenderBlock(100,
		func() string {
			s := ""
			s = fmt.Sprintf("%s--------------\n", s)
			s = fmt.Sprintf("%s  making a block\n", s)
			s = fmt.Sprintf("%s  making a block\n", s)
			s = fmt.Sprintf("%s  making a block\n", s)
			s = fmt.Sprintf("%s  making a block\n", s)
			s = fmt.Sprintf("%s  making a block %d\n", s, x)
			s = fmt.Sprintf("%s  Press Enter to stop", s)

			x++
			return s
		},
	)

	fmt.Scanln()
	cancel()
}
