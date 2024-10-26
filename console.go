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

	renderer *renderer
}

func MakeConsole(r io.Reader, w io.Writer) Console {
	c := Console{reader: r, writer: w}
	c.renderer = &renderer{}
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
				io.WriteString(c.writer,
					c.renderer.render(genMsg()))
			}
		}
	}()

	return cancel
}

type renderer struct {
	lastLineCount byte
}

func (r *renderer) render(msg string) string {

	b := ansi.Cursor(ansi.Up, r.lastLineCount) +
		ansi.Cursor(ansi.Left, 200) +
		ansi.ClearDown() +
		msg

	r.lastLineCount = byte(strings.Count(msg, "\n"))

	return b
}
