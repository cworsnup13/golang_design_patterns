package main

type Maze struct {
	rooms map[int]*Room
}

func (m *Maze) AddRoom(r *Room) {
	m.rooms[len(m.rooms)] = r
}

func (m *Maze) GetRoom(no int) *Room {
	return m.rooms[no]
}

type MazeGame struct {
}

func (mg *MazeGame) CreateMaze(factory MazeFactory) *Maze {
	aMaze := factory.MakeMaze()
	r1 := factory.MakeRoom(1)
	r2 := factory.MakeRoom(2)
	aDoor := factory.MakeDoor(r1, r2)

	aMaze.AddRoom(r1)
	aMaze.AddRoom(r2)

	r1.SetSide(North, factory.MakeWall())
	r1.SetSide(East, aDoor)
	r1.SetSide(South, factory.MakeWall())
	r1.SetSide(West, factory.MakeWall())

	r2.SetSide(North, factory.MakeWall())
	r2.SetSide(East, factory.MakeWall())
	r2.SetSide(South, factory.MakeWall())
	r2.SetSide(West, aDoor)

	return aMaze
}
