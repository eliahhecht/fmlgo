package main

import (
	"encoding/json"
	"io/ioutil"
)

func readJSONFile(path string, contentsHolder interface{}) {
	var fileContents = readFile(path)
	err := json.Unmarshal(fileContents, contentsHolder)
	if err != nil {
		panic(err)
	}
}

func readFile(path string) []byte {
	fileContents, err := ioutil.ReadFile(path)
	if err == nil {
		return fileContents
	}
	panic(err)
}
