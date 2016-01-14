package main

import (
	"bytes"
	"testing"
)

func TestSimple(t *testing.T) {
	codeGen := new(CodeSimple)
	if err := codeGen.InitFromRawConfig([]byte(`{"Comment":"comment test"}`)); err != nil {
		t.Error("error init from config: ", err.Error())
	}

	if codeGen.Comment != "comment test" {
		t.Error("not expected comment: ", codeGen.Comment)
	}

	code := codeGen.Generate()

	if !bytes.Equal(code, []byte("// comment test\n")) {
		t.Error("not expected code: ", string(code))
	}
}
