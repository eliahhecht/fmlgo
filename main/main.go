package main
import (
	"net/http"
	"html/template"
)

func showOwners(w http.ResponseWriter, r *http.Request) {
	players := buildPlayers(allCards)
	t, err := template.ParseFiles("templates/rosters.html")
	if err != nil {
		panic(err)
	}
	viewData := formatForRosters(players)
	t.Execute(w, viewData)
}

func formatForRosters(players []*Player) rosterData {
	data := rosterData{}
	data.Names = make([]PlayerName, len(players))
	for playerIdx, player := range players {
		data.Names[playerIdx] = player.Name
		for cardIdx, card := range player.Cards {
			row := data.Rows[cardIdx]
			data.Rows[cardIdx] = append(row, card)
		}
	}
	return data
}

type rosterData struct {
	Names []PlayerName
	Rows [][]*Card
}

var allCards *CardCollection

func main() {
	allCards = loadAllCards()

	http.HandleFunc("/", showOwners)
	http.ListenAndServe(":8080", nil)
}