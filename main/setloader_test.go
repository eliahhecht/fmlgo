package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var set = loadSet("ORI")

func TestLoadSetReturnsCorrectCards(t *testing.T) {
	assert.Contains(t, set.Cards, Card("Yavimaya Coast"))
}

func TestLoadSetDoesNotIncludeBasicLands(t *testing.T) {
	assert.NotContains(t, set.Cards, Card("Swamp"))
}
