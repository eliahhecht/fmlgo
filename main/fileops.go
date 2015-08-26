package main

import (
    "encoding/json"
    "io/ioutil"
)

func readJsonFile(path string, contentsHolder interface{}) {
    var fileContents = readFile(path)
    json.Unmarshal(fileContents, contentsHolder)
}

func readFile(path string) []byte {
    fileContents, err := ioutil.ReadFile(path)
    if err == nil {
        return fileContents
    } else {
        panic(err)
    }
}
