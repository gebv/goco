package main

import (
	"encoding/json"
	"strings"
)

const MODEL_DTO = "dto"

type CodeModelDTO struct {
	Comment           string
	Name              string
	Fields            Fields
	FactoryCustomCode []string
	Implements        []string

	Transformer CodeTransformer
}

func (m *CodeModelDTO) InitFromRawConfig(in []byte) error {

	if err := json.Unmarshal(in, m); err != nil {
		return err
	}

	return nil
}

func (m CodeModelDTO) Type() string {
	return MODEL_DTO
}

func (m CodeModelDTO) Generate() []byte {
	c := new(Code)

	// Factory model
	c.Printf("// %s %s\n", m.Name, m.Comment)
	c.Printf("func New%[1]s() *%[1]s{\n", m.Name)
	c.Printf("model := new(%s)\n", m.Name)
	if len(m.FactoryCustomCode) > 0 {
		c.Println("// Custom factory code")
		c.Printf("%s\n", strings.Join(m.FactoryCustomCode, "\n"))
	}
	c.Printf("return model\n}\n")

	// Struct model
	c.Printf("type %s struct {\n", m.Name)

	if len(m.Implements) > 0 {
		c.Printf("%s\n", strings.Join(m.Implements, "\n"))
	}

	c.buf.Write(m.Fields.InStruct())
	c.Printf("\n}\n")

	// Transformer
	c.buf.Write(m.Transformer.To.Generate(m.Name))
	c.Printf("\n")
	c.buf.Write(m.Transformer.From.Generate(m.Name))
	c.Printf("\n")

	c.Printf("//\n// Helpful functions\n//\n\n\n")

	selfName := firstLower(m.Name)

	// Maps
	// TODO: move to extends?
	if existValueArrayString(m.Implements, "ModelAbstract") {
		c.Printf("func (%s %s) Maps() map[string]interface{} {\n", selfName, m.Name)
		c.Printf("maps := %s.ModelAbstract.Maps()\n", selfName)
		c.buf.Write(m.Fields.InMap(m.Name))
		c.Printf("return maps\n}\n\n")
	}

	// Fields
	c.Printf("// Fields extract of fields from map\n")
	c.Printf("func (%s %s) Fields(fields ...string) ([]string, []interface{}) {\n", selfName, m.Name)
	c.Printf("return ExtractFieldsFromMap(%s.Maps(), fields...)\n}\n\n", selfName)

	//
	c.Printf("// FromJson data as []byte or io.Reader\n")
	c.Printf("func (%s *%s) FromJson(data interface{}) error {\n", selfName, m.Name)
	c.Printf("return FromJson(%s, data)\n}\n\n", selfName)

	return c.Format()
}

func existValueArrayString(arr []string, value string) bool {
	for _, _value := range arr {
		if _value == value {
			return true
		}
	}

	return false
}
