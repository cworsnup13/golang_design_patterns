package main

type MazeFactory interface {
	MakeMaze() *Maze
	MakeWall() *Wall
	MakeRoom(n int) *Room
	MakeDoor(r1, r2 *Room) *Door
}

type SimpleMazeFactory struct {

}

func (m SimpleMazeFactory) MakeMaze() *Maze {
	var rooms = make(map[int]*Room, 0)
	return &Maze{
		rooms: rooms,
	}
}

func (m SimpleMazeFactory) MakeWall() *Wall {
	return &Wall{}
}

func (m SimpleMazeFactory) MakeRoom(n int) *Room {
	// Default to 4 walled room
	var walls = make([]MapSite, 4)
	for i := range walls {
		walls[i] = &Wall{}
	}
	return &Room{
		roomNumber: n,
		sides:      walls,
	}
}
func (m SimpleMazeFactory) MakeDoor(r1, r2 *Room) *Door {
	return &Door{
		room1: r1,
		room2: r2,
	}
}
