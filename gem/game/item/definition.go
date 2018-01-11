package item

import (
	"encoding/json"
	"io/ioutil"
)

var definitions []ItemDefinition

//glua:bind
type ItemDefinition struct {
	Id        int
	Name      string
	Examine   string
	Noted     bool
	Notable   bool
	Stackable bool
	ParentId  int
	NotedId   int
	Members   bool
	ShopValue int
}

func LoadItemDefinitions(path string) (int, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return 0, err
	}

	err = json.Unmarshal(data, &definitions)
	if err != nil {
		return 0, err
	}

	return len(definitions), nil
}
