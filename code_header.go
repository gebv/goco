package main

import (
	"strings"
)

type CodeHeader struct {
	PkgName string
	Imports []string

	c *Code
}

func (h *CodeHeader) Generate() []byte {
	h.c = new(Code)

	h.c.Println("// Code generated. DO NOT EDIT.")
	h.c.Printf("package %v\n", h.PkgName)

	if len(h.Imports) == 0 {
		return h.c.Format()
	}

	h.c.Println("import (")
	// TODO: test
	for _, _import := range h.Imports {
		if strings.HasPrefix(_import, ". ") {
			h.c.Printf(". \"%s\"\n", _import[2:])
		} else {
			h.c.Printf("\"%s\"\n", _import)
		}
	}
	h.c.Printf(")\n")

	return h.c.Format()
}
