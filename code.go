package main

import (
	"bytes"
	"fmt"
	"go/format"
)

type Code struct {
	buf bytes.Buffer
}

func (c *Code) Println(str ...interface{}) {
	fmt.Fprintln(&c.buf, str...)
}

func (c *Code) Printf(format string, args ...interface{}) {
	fmt.Fprintf(&c.buf, format, args...)
}

func (c *Code) Format() []byte {
	src, err := format.Source(c.buf.Bytes())
	if err != nil {
		fmt.Println("internal error: " + err.Error())

		return c.buf.Bytes()
	}

	return src
}
