package kube

import (
	"fmt"
	"github.com/liushuochen/gotable"
	"github.com/liushuochen/gotable/table"
)

type Img struct {
	name         string `json:"name"`
	resourceType string `json:"type"`
	namespace    string `json:"namespace"`
	image        string `json:"image"`
}

var title []string = []string{
	"NAMESPACE",
	"TYPE",
	"NAME",
	"IMAGE",
}

func GenTable(mapList []map[string]string) *table.Table {
	t, err := gotable.Create(title...)
	if err != nil {
		fmt.Printf("create table error: %s", err.Error())
		return nil
	}
	t.AddRows(mapList)
	return t
}
