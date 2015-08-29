package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadSetReturnsCorrectCards(t *testing.T) {
	set := loadSet("ORI")
	assert.Contains(t, set.Cards, Card("Yavimaya Coast"))
}
