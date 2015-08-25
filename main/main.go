package main

import (
    "fmt"
    "io/ioutil"
    "encoding/json"
)

type Card struct {
    Name string
}

type Set struct {
    Code string
    Cards []Card
}

func main() {
    file, err := ioutil.ReadFile("./AllSets.json")
    if err != nil {
        panic(err)
    }

    var sets map[string]Set
    json.Unmarshal(file, &sets)

    fmt.Printf("%v", sets["ORI"])
}