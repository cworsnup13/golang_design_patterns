package main

import "fmt"

type Maze struct {
	rooms map[int]*Room
}

func NewMaze() *Maze {
	var rooms = make(map[int]*Room)
	maze := Maze{rooms: rooms}
	return &maze
}

func (m *Maze) AddRoom(r *Room) {
	m.rooms[len(m.rooms)] = r
	fmt.Printf("Adding %#v\n", m.rooms)
}

func (m *Maze) GetRoom(no int) *Room {
	return m.rooms[no-1]
}

type MazeGame struct {
}

func (mg *MazeGame) CreateMaze(builder MazeBuilder) *Maze {
	builder.BuildMaze()

	builder.BuildRoom(1)
	builder.BuildRoom(2)
	fmt.Printf("*** %#v\n", builder.GetMaze().GetRoom(2))
	builder.BuildDoor(1, 2)

	return builder.GetMaze()
}
