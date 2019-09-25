package q1

// IGameTable :
type IGameTable interface {
	Count() int
}

// GameTable : constructor for IGameTable
func GameTable(n int) IGameTable {
	return &table{n}
}

type table struct {
	numberOfPlayers int
}

// Count : count number of games which are not draw
func (t *table) Count() int {
	cnt := 0
	for _, game := range t.generate() {
		if !game.IsDraw() {
			cnt++
		}
	}
	return cnt
}

// generate : generate games
func (t *table) generate() []IGame {
	var games []IGame
	for i := 0; i <= t.numberOfPlayers; i++ {
		for j := 0; j <= t.numberOfPlayers-i; j++ {
			games = append(games, &game{i, j, t.numberOfPlayers - (i + j)})
		}
	}
	return games
}
