package main

type CodeFuncInit struct {
	Rows []string

	c *Code
}

func (h *CodeFuncInit) Generate() []byte {
	h.c = new(Code)

	if len(h.Rows) == 0 {
		return h.c.Format()
	}

	h.c.Println("func init() {")
	for _, row := range h.Rows {
		h.c.Println(row)
	}
	h.c.Println("}")

	return h.c.Format()
}
