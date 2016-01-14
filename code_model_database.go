package main

import (
	"encoding/json"
)

func init() {
	RegGenType(&CodeModelDataBase{})
}

const MODEL_DATABASE = "database"

type CodeModelDataBase struct {
	CodeModelDTO

	TableName string
}

func (m *CodeModelDataBase) InitFromRawConfig(in []byte) error {

	if err := json.Unmarshal(in, m); err != nil {
		return err
	}

	return nil
}

func (m CodeModelDataBase) Type() string {
	return MODEL_DATABASE
}

func (m CodeModelDataBase) Generate() []byte {
	c := new(Code)
	c.buf.Write(m.CodeModelDTO.Generate())

	// Render table name
	c.Printf("func (%s) TableName() string{\n", m.Name)
	c.Printf("return \"%s\"\n}\n", m.TableName)

	// Render primary key
	if field, err := m.Fields.PrimaryKey(); err == nil {
		c.Printf("// PrimaryName primary field name\n")
		c.Printf("func (%s) PrimaryName() string{\n", m.Name)
		name := toLower(field.Name, "_")
		if len(field.DBName) > 0 {
			name = field.DBName
		}
		c.Printf("return \"%s\"\n}\n", name)

		selfName := firstLower(m.Name)
		c.Printf("// PrimaryValue primary value\n")
		c.Printf("func (%s %s) PrimaryValue() %s{\n", selfName, m.Name, field.TypeName)
		c.Printf("return %s.%s\n}\n", selfName, field.Name)
	}
	c.Println("// model")

	return c.Format()
}
