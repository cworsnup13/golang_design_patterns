package main

import (
	"fmt"
)

type Direction int

const (
	North Direction = iota
	South
	East
	West
)


type MapSite interface {
	Enter(p *Player) bool
}

type Room struct {
	sides      []MapSite
	roomNumber int
}

func (r *Room) Enter(p *Player) bool {
	p.SetCurrent(r)
	fmt.Println("You've entered room", r.roomNumber)
	return true
}

func (r *Room) GetSide(direction Direction) MapSite {
	return r.sides[direction]
}

func (r *Room) SetSide(direction Direction, ms MapSite) {
	r.sides[direction] = ms
}

type Wall struct {
}

func (w *Wall) Enter(p *Player) bool {
	fmt.Println("You've hit a wall!")
	return false
}

type Door struct {
	room1, room2 *Room
	isOpen       bool
}

func (d *Door) Enter(p *Player) bool {
	fmt.Println("You've found a door, entering.")
	d.isOpen = true
	if p.GetCurrent() == d.room1 {
		return d.room2.Enter(p)
	} else {
		return d.room1.Enter(p)
	}

}
