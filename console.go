package console

import (
	"bytes"
	"console/ansi"
	"context"
	"fmt"
	"io"
	"strings"
	"time"
)

type Console struct {
	reader io.Reader
	writer io.Writer

	renderer *renderer
}

func MakeConsole(r io.Reader, w io.Writer) Console {
	c := Console{reader: r, writer: w}
	c.renderer = &renderer{writer: w}
	return c
}

func (c Console) Printf(format string, a ...any) {
	fmt.Fprintf(c.writer, format, a...)
}

func (c Console) RenderBlock(rateHz uint, genMsg func() string) context.CancelFunc {
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		ticker := time.NewTicker(time.Duration((1.0 / float64(rateHz)) * float64(time.Second)))
		defer ticker.Stop()

		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				b := c.renderer.render(genMsg())
				c.writer.Write(b)
				//io.WriteString(c.writer, string(b))
			}
		}

	}()

	return cancel
}

type renderer struct {
	writer        io.Writer
	lastLineCount byte
}

func (r *renderer) render(msg string) []byte {

	a := bytes.NewBufferString(
		ansi.Cursor(ansi.Up, r.lastLineCount) +
			ansi.Cursor(ansi.Left, 200) +
			ansi.ClearDown(),
	).Bytes()

	b := bytes.NewBufferString(msg).Bytes()

	lines := strings.Split(msg, "\n")
	r.lastLineCount = byte(len(lines)) - 1

	return append(a, b...)
}
