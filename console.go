package console

import (
	"context"
	"fmt"
	"io"
	"strings"
	"time"

	"github.com/nadedan/console/ansi"
)

type Console struct {
	reader io.Reader
	writer io.Writer
}

func MakeConsole(r io.Reader, w io.Writer) Console {
	c := Console{reader: r, writer: w}
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

		r := newRenderer(c.writer)
		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				r.render(genMsg())
			}
		}
	}()

	return cancel
}

type renderer struct {
	writer        io.Writer
	lastLineCount byte
}

func newRenderer(w io.Writer) *renderer {
	return &renderer{writer: w}
}

func (r *renderer) render(msg string) {

	io.WriteString(r.writer, ""+
		ansi.Cursor(ansi.Up, r.lastLineCount)+
		ansi.Cursor(ansi.Left, 200)+
		ansi.ClearDown(),
	)

	io.WriteString(r.writer, msg)

	r.lastLineCount = byte(strings.Count(msg, "\n"))
}
