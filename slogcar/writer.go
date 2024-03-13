package slogcar

import (
	"io"
	"strings"

	"github.com/lingdor/go-logcar/entity"
)

type carWriter struct {
	w io.Writer
}

func (c *carWriter) Write(p []byte) (n int, err error) {

	txt := string(p)
	txt = strings.ReplaceAll(txt, "\n", string([]byte{byte('\n'), entity.WrapChar}))
	return c.w.Write([]byte(txt))
}
