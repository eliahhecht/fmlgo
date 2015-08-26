package main


type Card string

type Set struct {
    Code  string
    Cards []Card
}

type Player struct {
    Name string
    Cards []Card
}