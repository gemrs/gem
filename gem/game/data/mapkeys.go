package data

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

var mapKeys = map[int][]int{}

type mapKeyEntry struct {
	Region int   `json:"region"`
	Keys   []int `json:"keys"`
}

//glua:bind
func LoadMapKeys(path string) {
	fd, err := os.Open(path)
	defer fd.Close()
	if err != nil {
		panic(err)
	}

	data, err := ioutil.ReadAll(fd)
	if err != nil {
		panic(err)
	}

	var keysList []mapKeyEntry
	err = json.Unmarshal(data, &keysList)
	if err != nil {
		panic(err)
	}

	for _, k := range keysList {
		mapKeys[k.Region] = k.Keys
	}
}

func GetMapKeys(region int) ([]int, bool) {
	k, ok := mapKeys[region]
	return k, ok
}
