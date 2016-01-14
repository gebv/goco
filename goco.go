package main

import (
	"flag"
	"fmt"
	"os"
	// "path/filepath"
	// "bytes"
	"io/ioutil"
)

var (
	codeType   = flag.String("type", "", "type code")
	outputFile = flag.String("out", "", "output file name")
	configFile = flag.String("config", "", "config file (json)")
)

var allowedTypes = make(map[string]Generator)

func RegGenType(g Generator) {
	allowedTypes[g.Type()] = g
}

func main() {
	flag.Parse()
	var configRaw []byte
	var err error

	if len(*configFile) == 0 {
		configRaw, err = ioutil.ReadAll(os.Stdin)
	} else {
		pwd, _ := os.Getwd()
		pathConfig := pwd + string(os.PathSeparator) + *configFile

		if _, err := os.Stat(pathConfig); err != nil {
			if os.IsNotExist(err) {
				fmt.Println("goco: file not exist: " + err.Error())
			}

			fmt.Println("goco: error config file: " + err.Error())
			os.Exit(0)
		}

		configRaw, err = ioutil.ReadFile(pathConfig)
	}

	if err != nil {
		fmt.Println("goco: error config data: " + err.Error())
		os.Exit(0)
	}

	gen, exist := allowedTypes[*codeType]

	if !exist {
		fmt.Println("goco: not supported type: " + *codeType)
		os.Exit(0)
	}

	_, ok := gen.(InitFromRawConfig)

	if !ok {
		fmt.Println("goco: not supported load config type " + *codeType)
		os.Exit(0)
	}

	if err := gen.(InitFromRawConfig).InitFromRawConfig(configRaw); err != nil {
		fmt.Println("goco: init from raw config: " + err.Error())
		os.Exit(0)
	}

	if *outputFile == "stdout" {
		os.Stdout.Write(gen.Generate())
	} else {
		if err := ioutil.WriteFile(*outputFile, gen.Generate(), 0644); err != nil {
			fmt.Println("goco: error write file: " + err.Error())
			os.Exit(0)
		}
	}
}

type Generator interface {
	Type() string
	Generate() []byte
}

type InitFromRawConfig interface {
	InitFromRawConfig(in []byte) error
}
