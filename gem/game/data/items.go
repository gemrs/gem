package data

/*
var items []ItemDefinition

type ItemDefinition struct {
	Id            int
	Name          string
	Stackable     int
	NotedTemplate int      `json:"notedTemplate"`
	NotedId       int      `json:"notedID"`
	ShopValue     int      `json:"cost"`
	GroundActions []string `json:"options"`
	Actions       []string `json:"interfaceOptions"`
	Team          int
}

//glua:bind
func LoadItems(dir string) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		panic(err)
	}

	items = make([]ItemDefinition, len(files))

	for _, f := range files {
		idx, err := strconv.Atoi(f.Name()[:len(f.Name())-5])
		if err != nil {
			panic(err)
		}

		data, err := ioutil.ReadFile(dir + "/" + f.Name())
		if err != nil {
			panic(err)
		}

		err = json.Unmarshal(data, &items[idx])
		if err != nil {
			panic(err)
		}
	}

	logger.Notice("Loaded [%v] item definitions", len(items))
}
*/
