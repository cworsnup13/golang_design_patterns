package main

import (
	"bytes"
	"encoding/gob"
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
	Clone() MapSite
}

type Room struct {
	Sides      []MapSite
	roomNumber int
}

func initMapSites() {
	gob.Register(&Wall{})
	gob.Register(&Room{})
	gob.Register(&Door{})
}

func NewRoom() *Room {
	var walls = make([]MapSite, 4)
	for i := range walls {
		walls[i] = &Wall{}
	}
	return &Room{
		roomNumber: 0,
		Sides:      walls,
	}
}

func (r *Room) Enter(p *Player) bool {
	p.SetCurrent(r)
	fmt.Println("You've entered room", r.roomNumber)
	return true
}

func (r *Room) Clone() MapSite {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	dec := gob.NewDecoder(&buf)
	err := enc.Encode(*r) // ignoring error in this toy example
	if err != nil {
		fmt.Println(err)
	}

	newRoom := Room{}
	err = dec.Decode(&newRoom)
	if err != nil {
		fmt.Println(err)
	}
	return &newRoom
}

func (r *Room) GetSide(direction Direction) MapSite {
	return r.Sides[direction]
}

func (r *Room) SetSide(direction Direction, ms MapSite) {
	r.Sides[direction] = ms
}

type Wall struct {
}

func (w *Wall) Enter(p *Player) bool {
	fmt.Println("You've hit a wall!")
	return false
}

func (w *Wall) Clone() MapSite {
	newWall := Wall{}
	newWall = *w
	return &newWall
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

func (d *Door) Clone() MapSite {
	newDoor := Door{}
	newDoor = *d
	return &newDoor
}
