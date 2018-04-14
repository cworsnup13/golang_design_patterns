package main

type MazePrototypeFactory interface {
	MakeMaze() *Maze
	MakeWall() *Wall
	MakeRoom(n int) *Room
	MakeDoor(r1, r2 *Room) *Door
}

type SimpleMazePrototypeFactory struct {
	prototypeMaze *Maze
	prototypeWall *Wall
	prototypeRoom *Room
	prototypeDoor *Door
}

func NewMazePrototypeFactory(m *Maze, w *Wall, r *Room, d *Door) MazePrototypeFactory {
	initMapSites()
	return SimpleMazePrototypeFactory{
		prototypeMaze: m,
		prototypeWall: w,
		prototypeDoor: d,
		prototypeRoom: r,
	}
}

func (m SimpleMazePrototypeFactory) MakeMaze() *Maze {
	var rooms = make(map[int]*Room, 0)
	return &Maze{
		rooms: rooms,
	}
}

func (m SimpleMazePrototypeFactory) MakeWall() *Wall {
	return m.prototypeWall.Clone().(*Wall)
}

func (m SimpleMazePrototypeFactory) MakeRoom(n int) *Room {
	p := m.prototypeRoom.Clone().(*Room)
	p.roomNumber = n
	return p
}

func (m SimpleMazePrototypeFactory) MakeDoor(r1, r2 *Room) *Door {
	p := m.prototypeDoor.Clone().(*Door)
	p.room1 = r1
	p.room2 = r2
	return p
}
