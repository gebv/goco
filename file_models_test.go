package main

import (
	"testing"
)

func TestFileModels(t *testing.T) {
	fileModels := new(FileModels)
	fileMOdelsConfig := `{
	"PkgName": "packagename",
	"Imports": ["tratata/tata"],
	"Models": [
		{
			"Type": "database",
			"Name": "ModelOne",
			"Implements": ["ModelAbstract"],
			"TableName": "items",
			"Fields": [
				{
					"Comment": "Primary key",
					"Name": "ItemId",
					"Type": "int64",
					"Tags": {
						"v": "required,min=1"
					},
					"IsPrimaryKey": true
				}
			],
			"Transformer": {
				"To": [
					{
						"Name": "ExampleModel",
						"Map" :{
							"ItemId": "ItemId"
						}
					}
				],
				"From": [
					{
						"Name": "ExampleModel",
						"Map" :{
							"ItemId": "ItemId"
						},
						"Custom": [
							"model.T = \"1\""
						]
					}
				]
			}
		},
		{
			"Type": "dto",
			"Name": "DTOOne",
			"Comment": "comment dto1"
		},
		{
			"Type": "dto",
			"Name": "DTOTwo",
			"Comment": "comment dto2"
		}
	]
}`
	fileModels.InitFromRawConfig([]byte(fileMOdelsConfig))

	if len(fileModels.Models) != 3 {
		t.Error("unexpected count models")
	}
}
