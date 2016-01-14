package main

import (
	"encoding/json"
)

func init() {
	RegGenType(&FileModels{})
}

type FileModels struct {
	PkgName string
	Imports []string

	Models Models
}

func (FileModels) Type() string {
	return "file_models"
}

func (m *FileModels) InitFromRawConfig(in []byte) error {

	if err := json.Unmarshal(in, m); err != nil {
		return err
	}

	return nil
}

func (m *FileModels) Generate() []byte {
	c := new(Code)

	header := new(CodeHeader)
	header.Imports = m.Imports
	header.PkgName = m.PkgName
	c.buf.Write(header.Generate())

	c.buf.Write(m.Models.Generate())

	return c.Format()
}
