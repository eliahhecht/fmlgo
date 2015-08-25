package main

import (
    "encoding/json"
    "io/ioutil"
    "fmt"
)

type Card struct {
    Name string
}

type Set struct {
    Code  string
    Cards []Card
}

type Player struct {
    Name string
    Cards []Card
}

func main() {
    var sets map[string]Set
    readJsonFile("./AllSets.json", &sets)
    ori := sets["ORI"]

    players := buildPlayers()

    checkSanity(players, []Set{ori})
}

func buildPlayers() []Player {
    var playersMap map[string][]string
    readJsonFile("./owners.json", &playersMap)

    players := make([]Player, 0)
    for k, v := range playersMap {
        player := Player{Name: k, Cards: buildCards(v)}
        players = append(players, player)
    }
    return players
}

func buildCards(cardNames []string) []Card {
    var cards = make([]Card, 0)
    for _, cardName := range cardNames {
        cards = append(cards, Card{Name: cardName})
    }
    return cards
}

func checkSanity(players []Player, legalSets []Set) {
    for _, player := range players {
        for _, card := range player.Cards {
            if !anySetContainsCard(legalSets, card) {
                panic(fmt.Sprintf("%s's card %s not found in any legal set", player.Name, card))
            }
        }
    }
}

func anySetContainsCard(sets []Set, card Card) bool {
    for _, set := range sets {
        if set.containsCard(card) {
            return true
        }
    }
    return false
}

func (set Set) containsCard(targetCard Card) bool {
    for _, card := range set.Cards {
        if card == targetCard {
            return true
        }
    }
    return false
}

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
