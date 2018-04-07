package main

import "fmt"

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

type MazeFactory interface {
	MakeMaze() *Maze
	MakeWall() *Wall
	MakeRoom(n int) *Room
	MakeDoor(r1, r2 *Room) *Door
}

type Player struct {
	current *Room
}

func (p *Player) SetCurrent(room *Room) {
	p.current = room
}

func (p *Player) GetCurrent() *Room {
	return p.current
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

//type EnchantedMazeFactory struct {
//
//}
//
//func (m *EnchantedMazeFactory) MakeMaze() *Maze {
//
//}
//func (m *EnchantedMazeFactory) MakeWall() *Wall {
//
//}
//func (m *EnchantedMazeFactory) MakeRoom(n int) *Room {
//
//}
//
//func (m *EnchantedMazeFactory) MakeDoor(r1, r2 *Room) *Door {
//
//}

type Maze struct {
	rooms map[int]*Room
}

func (m *Maze) AddRoom(r *Room) {
	m.rooms[len(m.rooms)] = r
}

func (m *Maze) RoomNo(no int) *Room {
	return m.rooms[no]
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

func main() {
	mazeGame := MazeGame{}
	mazeFactory := SimpleMazeFactory{}
	maze := mazeGame.CreateMaze(mazeFactory)
	player := Player{}
	player.SetCurrent(maze.RoomNo(0))
	fmt.Println("Starting Room Number", player.GetCurrent().roomNumber)
	var ok bool
	var numDoors = 4
	for numDoors > 0 {
		for i := 0; i < 4; i++ {
			side := player.GetCurrent().GetSide(Direction(i))
			ok = side.Enter(&player)
			if ok {
				break
			}
		}
		numDoors--
	}
	fmt.Println("Complete")

}
