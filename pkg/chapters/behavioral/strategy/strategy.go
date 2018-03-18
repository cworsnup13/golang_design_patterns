package main

import "fmt"

type Compositor interface {
	Compose(natural, stretch, shrink, breaks []int, compCount, lineWidth int) int
}

type Component interface {
	NaturalSize() int
	Stretch() int
	Shrink() int
}

type Composition struct {
	compositor Compositor
	components []Component

	componentCount int
	lineWidth      int
	lineBreaks     []int
	lineCount      int
}

func NewComposition(c Compositor) Composition {
	return Composition{
		compositor: c,
	}
}

func (c *Composition) Repair() {

	var natural = make([]int, len(c.components))
	var stretch = make([]int, len(c.components))
	var shrink = make([]int, len(c.components))

	for i, v := range c.components {
		natural[i] = v.NaturalSize()
		stretch[i] = v.Stretch()
		shrink[i] = v.Shrink()
	}

	breakCount := c.compositor.Compose(natural, stretch, shrink, c.lineBreaks, c.componentCount, c.lineWidth)
	// Lay out components according to breaks
	fmt.Println("Breaks", breakCount)
}

type TeXCompositor struct {

}

func NewTeXCompositor() TeXCompositor{
	return TeXCompositor{}
}

func (c *TeXCompositor) Compose(natural, stretch, shrink, breaks []int, compCount, lineWidth int) int {
	return 0
}

type SimpleCompositor struct {

}

func NewSimpleCompositor() SimpleCompositor{
	return SimpleCompositor{}
}

func (c *SimpleCompositor) Compose(natural, stretch, shrink, breaks []int, compCount, lineWidth int) int {
	return 2
}

func main() {
	s := NewSimpleCompositor()
	t:= NewTeXCompositor()

	simple := NewComposition(&s)
	tex := NewComposition(&t)

	simple.Repair()
	// 2

	tex.Repair()
	// 0
}
