package main

import (
	"strings"
)

type CodeTransformer struct {
	To   TransofmerToArray
	From TransofmerFromArray
}

type TransofmerToArray []TransformConfig

func (t TransofmerToArray) Generate(modelName string) []byte {
	c := new(Code)

	c.Printf("func (model %s) TransformTo(out interface{}) error {\n", modelName)
	c.Printf("switch out.(type) {\n")
	if len(t) > 0 {
		for _, _case := range t {
			c.Printf("case *%s:\n", _case.Name)
			c.Printf("dto := out.(*%s)\n", _case.Name)
			// case fields
			for fieldTo, fieldFtom := range _case.Map {
				c.Printf("dto.%s = model.%s\n", fieldTo, fieldFtom)
			}

			if len(_case.Custom) > 0 {
				c.Printf("%s\n", strings.Join(_case.Custom, "\n"))
			}
		}
	}
	c.Println("default:\n\tglog.Errorf(\"Not supported type %v\",out);\n\treturn ErrNotSupported\n}")
	c.Printf("return nil\n")
	c.Printf("}\n")

	return c.Format()
}

type TransofmerFromArray []TransformConfig

func (t TransofmerFromArray) Generate(modelName string) []byte {
	c := new(Code)

	c.Printf("func (model *%s) TransformFrom(in interface{}) error {\n", modelName)
	c.Printf("switch in.(type) {\n")
	if len(t) > 0 {
		for _, _case := range t {
			c.Printf("case *%s:\n", _case.Name)
			c.Printf("dto := in.(*%s)\n", _case.Name)
			// case fields
			for fieldTo, fieldFtom := range _case.Map {
				c.Printf("model.%s = dto.%s\n", fieldTo, fieldFtom)
			}

			if len(_case.Custom) > 0 {
				c.Printf("%s\n", strings.Join(_case.Custom, "\n"))
			}
		}
	}
	c.Println("default:\n\tglog.Errorf(\"Not supported type %v\",in);\n\treturn ErrNotSupported\n}\nreturn nil\n")
	c.Printf("}\n")

	return c.Format()
}

type TransformConfig struct {
	Name   string
	Map    map[string]string
	Custom []string
}
