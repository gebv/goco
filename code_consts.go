package main

import (
	"encoding/json"
)

func init() {
	RegGenType(&CodeConsts{})
}

type CodeConsts struct {
	Comment string
	Consts  []string
}

func (c *CodeConsts) InitFromRawConfig(in []byte) error {

	if err := json.Unmarshal(in, c); err != nil {
		return err
	}

	return nil
}

func (c CodeConsts) Type() string {
	return "consts"
}

func (cs *CodeConsts) Generate() []byte {
	c := new(Code)

	c.Printf("// %s\n", cs.Comment)

	if len(cs.Consts) > 0 {
		for _, _c := range cs.Consts {
			c.Printf("const %s = \"%s\"\n", toUpper(_c, ""), toLower(_c, "_"))
		}
	}

	return c.Format()
}
