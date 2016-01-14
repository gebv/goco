package main

import (
	"encoding/json"
	"errors"
	"strings"
)

type Fields []CodeField

func (f Fields) PrimaryKey() (*CodeField, error) {
	for _, field := range f {
		if field.IsPrimaryKey {
			return &field, nil
		}
	}

	return nil, errors.New("not found")
}

func (f Fields) InStruct() []byte {
	c := new(Code)

	for _, field := range f {
		c.buf.Write(field.InStruct())
	}

	return c.buf.Bytes()
}

func (f Fields) InMap(modelName string) []byte {
	c := new(Code)

	for _, field := range f {
		c.buf.Write(field.InMap(modelName))
	}

	return c.buf.Bytes()
}

type CodeField struct {
	Comment  string
	Name     string // Field name
	DBName   string //
	TypeName string `json:"Type"`
	Tags     map[string]string

	IsPrimaryKey bool

	c *Code
}

func (f CodeField) InStruct() []byte {
	c := new(Code)
	c.Printf("// %s\t%s\n", f.Name, f.Comment)
	c.Printf("%s\t%s", f.Name, f.TypeName)
	if len(f.Tags) > 0 {
		c.Printf("\t`")

		if _, existJson := f.Tags["json"]; !existJson {
			c.Printf("json:\"%s\" ", toLower(f.Name, "_"))
		}

		for tagKey, tagValue := range f.Tags {
			c.Printf("%s:\"%s\" ", tagKey, tagValue)
		}

		c.Printf("`")
	}
	c.Printf("\n")

	return c.buf.Bytes()
}

// InMap maps["field_name"] = &selfName.FieldName
func (f CodeField) InMap(modelName string) []byte {
	//
	selfName := firstLower(modelName)
	dbName := strings.TrimSpace(f.DBName)
	if len(dbName) == 0 {
		dbName = toLower(f.Name, "_")
	}

	c := new(Code)
	c.Printf("// %s\t%s\n", f.Name, f.Comment)
	c.Printf("maps[\"%s\"] = %s.%s", dbName, selfName, f.Name)
	c.Printf("\n")

	return c.buf.Bytes()
}

func (f *CodeField) InitFromRawConfig(in []byte) error {

	if err := json.Unmarshal(in, f); err != nil {
		return err
	}

	return nil
}

func (f CodeField) Type() string {
	return "field"
}

func (f CodeField) Generate() []byte {
	c := new(Code)
	c.Println("// " + f.Type())
	return c.Format()
}
