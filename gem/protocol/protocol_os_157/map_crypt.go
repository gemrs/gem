package protocol_os_157

import (
	"bufio"
	"io/ioutil"
	"os"
	"strconv"
)

var mapKeys = map[int][]int{}

func LoadMapKeys(dir string) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		panic(err)
	}

	for _, f := range files {
		idx, err := strconv.Atoi(f.Name()[:len(f.Name())-4])
		if err != nil {
			panic(err)
		}

		mapKeys[idx] = make([]int, 0)

		fd, err := os.Open(dir + "/" + f.Name())
		if err != nil {
			panic(err)
		}

		scanner := bufio.NewScanner(fd)
		for scanner.Scan() {
			key, err := strconv.Atoi(scanner.Text())
			if err != nil {
				panic(err)
			}
			mapKeys[idx] = append(mapKeys[idx], key)
		}
	}
}
