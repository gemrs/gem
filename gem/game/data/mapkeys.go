package data

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

var mapKeys = map[int][]uint32{}

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
		mapKeys[k.Region] = make([]uint32, len(k.Keys))
		for i, _ := range k.Keys {
			mapKeys[k.Region][i] = uint32(k.Keys[i])
		}
	}

	logger.Notice("Loaded [%v] map keys", len(mapKeys))
}

func GetMapKeys(region int) ([]uint32, bool) {
	k, ok := mapKeys[region]
	return k, ok
}
