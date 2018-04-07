package main

type MazeBuilder interface {
	BuildMaze()
	BuildRoom(roomNumber int)
	BuildDoor(room1, room2 int)
	GetMaze() *Maze
}

type StandardMazeBuilder struct {
	currentMaze *Maze
}

func (b *StandardMazeBuilder) BuildMaze() {
	b.currentMaze = NewMaze()
}

func (b *StandardMazeBuilder) BuildRoom(roomNumber int) {
	if b.currentMaze.GetRoom(roomNumber) == nil {
		room := NewRoom(roomNumber)
		b.currentMaze.AddRoom(room)

		for i := 0; i < 4; i++ {
			room.SetSide(Direction(i), &Wall{})
		}
	}
}

func (b *StandardMazeBuilder) BuildDoor(room1, room2 int) {
	r1 := b.currentMaze.GetRoom(room1)
	r2 := b.currentMaze.GetRoom(room2)
	door := Door{room1: r1, room2: r2}

	r1.SetSide(b.CommonWall(r1, r2), &door)
	r2.SetSide(b.CommonWall(r2, r1), &door)
}

func (b *StandardMazeBuilder) GetMaze() *Maze {
	return b.currentMaze
}

func (b *StandardMazeBuilder) CommonWall(room1, room2 *Room) Direction {
	// This will take more thought to implement. This could be a fixed grid strategy. Ultimately it's the builder's
	// responsibility to organize the layout.

	// Every room could be connected.
	if room1.roomNumber-room2.roomNumber > 0 {
		return West
	}
	return East
}
