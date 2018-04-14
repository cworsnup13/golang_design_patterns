package main

import "fmt"

type Shape interface {
	BoundingBox(bottomLeft, topRight *Point)
	CreateManipulator() *Manipulator
}

type TextView interface {
	GetOrigin(x, y *Coord)
	GetExtent(width, height *Coord)
	IsEmpty() bool
}

type Point struct {
	x, y float64
}

type Coord struct {
	val float64
}

type Manipulator struct {
}

type SimpleTextView struct {
	originX, originY *Coord
	height, width    *Coord
	empty            bool
}

func (s *SimpleTextView) GetOrigin(x, y *Coord) {
	*x = *s.originX
	*y = *s.originY
}

func (s *SimpleTextView) GetExtent(width, height *Coord) {
	*width = *s.width
	*height = *s.height
}

func (s *SimpleTextView) IsEmpty() bool {
	return s.empty
}

type TextShape struct {
	text TextView // adaptee is a private variable for delegation
}

func (s *TextShape) BoundingBox(bottomLeft, topRight *Point) {
	var bottom, left, width, height Coord
	s.text.GetOrigin(&bottom, &left)
	s.text.GetExtent(&width, &height)
	*bottomLeft = Point{x: left.val, y: bottom.val}
	*topRight = Point{x: left.val + width.val, y: bottom.val + height.val}
}

func (s *TextShape) CreateManipulator() *Manipulator {
	return &Manipulator{}
}

func (s *TextShape) IsEmpty() bool {
	return s.text.IsEmpty()
}

func main() {
	fmt.Println("Adapting TextView to be compatible as a Shape")
	simple := SimpleTextView{
		originX: &Coord{1},
		originY: &Coord{1},
		height:  &Coord{5},
		width:   &Coord{10},
	}
	textShape := TextShape{&simple}

	var bottomLeft = &Point{}
	var topRight = &Point{}
	textShape.BoundingBox(bottomLeft, topRight)
	fmt.Println(bottomLeft, topRight)
}
