package main

import (
	"encoding/json"
	"fmt"
)

type Models []FactoryModel

func (m Models) Generate() []byte {
	c := new(Code)

	for _, _m := range m {
		c.buf.Write(_m.Generate())
	}

	return c.Format()
}

type FactoryModel struct {
	json.RawMessage
}

func (m *FactoryModel) Generate() []byte {
	var cg Generator

	_b, _ := m.MarshalJSON()
	typeModel := &struct {
		TypeName string `json:"Type"`
	}{}

	json.Unmarshal(_b, typeModel)

	switch typeModel.TypeName {
	case MODEL_DATABASE:
		cg = new(CodeModelDataBase)
		cg.(InitFromRawConfig).InitFromRawConfig(_b)
	case MODEL_DTO:
		cg = new(CodeModelDTO)
		cg.(InitFromRawConfig).InitFromRawConfig(_b)
	default:
		fmt.Println("error: not supported model type '" + typeModel.TypeName + "'")
		return []byte("// not supported model type '" + typeModel.TypeName + "'")
	}

	return cg.Generate()
}
