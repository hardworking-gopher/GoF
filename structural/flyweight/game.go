package main

type game struct {
	terrorists        []*Player
	counterTerrorists []*Player
}

func newGame() *game {
	return &game{
		terrorists:        make([]*Player, 0, 1),
		counterTerrorists: make([]*Player, 0, 1),
	}
}

func (c *game) addTerrorist(dressType string) {
	player := newPlayer("T", dressType)
	c.terrorists = append(c.terrorists, player)
	return
}

func (c *game) addCounterTerrorist(dressType string) {
	player := newPlayer("CT", dressType)
	c.counterTerrorists = append(c.counterTerrorists, player)
	return
}
