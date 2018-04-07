package main

type Player struct {
	current *Room
}

func (p *Player) SetCurrent(room *Room) {
	p.current = room
}

func (p *Player) GetCurrent() *Room {
	return p.current
}
