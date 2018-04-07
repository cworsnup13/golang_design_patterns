package main

import (
	"fmt"
)

func main() {
	mazeGame := MazeGame{}
	mazeFactory := SimpleMazeFactory{}
	maze := mazeGame.CreateMaze(mazeFactory)
	player := Player{}
	player.SetCurrent(maze.GetRoom(0))
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
