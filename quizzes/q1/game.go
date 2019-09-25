package q1

// IGame : じゃんけんの結果表
type IGame interface {
	IsDraw() bool
}

// Game : constructor for IGame
func Game(r, p, s int) IGame {
	return &game{r, p, s}
}

type game struct {
	rock     int
	paper    int
	scissors int
}

func (g *game) IsDraw() bool {
	return (g.rock == g.paper) && g.paper > g.scissors ||
		g.rock == g.scissors && g.scissors > g.paper ||
		g.scissors == g.paper && g.scissors > g.rock
}
