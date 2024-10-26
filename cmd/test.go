package main

import (
	"console"
	"fmt"
	"os"
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

	cancel := c.RenderBlock(10,
		func() string {
			s := ""
			s = fmt.Sprintf("%sstart of block\n", s)
			s = fmt.Sprintf("%smaking a block\n", s)
			s = fmt.Sprintf("%smaking a block\n", s)
			s = fmt.Sprintf("%smaking a block\n", s)
			s = fmt.Sprintf("%smaking a block\n", s)
			s = fmt.Sprintf("%smaking a block %d\n", s, x)
			s = fmt.Sprintf("%sPress Enter to stop\n", s)

			x++
			return s
		},
	)

	fmt.Scanln()
	cancel()
}
