package main

import "fmt"

// Conway's Game of Life Rules
// 1. Any live cell with fewer than two live neighbours dies, as if by underpopulation.
// 2. Any live cell with two or three live neighbours lives on to the next generation.
// 3. Any live cell with more than three live neighbours dies, as if by overpopulation.
// 4. Any dead cell with exactly 3 live neighbors becomes a live cell, as if by reproduction.

type World struct{}

/*
Generate a new world for your cells to live in.
This world should have a height and width and data structure for
your cells to exist in.
*/
func NewWorld(width, height int) *World {
	return &World{}
}

/*
Iterate through your world and create random values for your cells
*/
func (w *World) RandomizeWorld() {

}

/*
Count the number of alive neighbors that surround a cell.
return count of alive neighbors
*/
func (w *World) Neighbors(x, y int) int {
	return 0
}

/*
For a cell x, y determine if given the number of neighbors
the current cell is alive or dead
*/
func (w *World) Alive(x, y int) bool {
	return false
}

/*
For the existing world, generate a new world
by iterating through the existing world to determine alive or dead cells
and save that into a new world
returns new world of next generation of alive or dead cells
*/
func (w *World) NextGeneration() *World {
	return nil
}

/*
Iterate through the world and display the world into the terminal
*/
func (w *World) Display() {

}

func main() {
	world := NewWorld(3, 3)
	fmt.Println(world)
}
