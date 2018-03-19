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

type BasicComponent struct {
	Natural int
	Stretched int
	Shrinked int
}

func (b *BasicComponent) NaturalSize() int {
	return b.Natural
}

func (b *BasicComponent) Stretch() int {
	return b.Stretched
}

func (b *BasicComponent) Shrink() int {
	return b.Shrinked
}

type Composition struct {
	compositor Compositor
	components []Component

	componentCount int
	lineWidth      int
	lineBreaks     []int
	lineCount      int
}

func NewComposition(c Compositor, comps []Component) Composition {
	return Composition{
		compositor: c,
		components: comps,
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

	fmt.Println("Breaks", breakCount)
}

type TeXCompositor struct {

}

func NewTeXCompositor() TeXCompositor{
	return TeXCompositor{}
}

func (c *TeXCompositor) Compose(natural, stretch, shrink, breaks []int, compCount, lineWidth int) int {
	var total int
	for i := range natural {
		total += natural[i] * stretch[i]
	}
	return total
}

type SimpleCompositor struct {

}

func NewSimpleCompositor() SimpleCompositor{
	return SimpleCompositor{}
}

func (c *SimpleCompositor) Compose(natural, stretch, shrink, breaks []int, compCount, lineWidth int) int {
	var total int
	for i := range natural {
		total += natural[i]
	}
	return total
}

func main() {

	c1 := BasicComponent{
		Natural:3,
		Stretched: 2,
		Shrinked: 1,
	}

	c2 := BasicComponent{
		Natural:3,
		Stretched: 2,
		Shrinked: 1,
	}

	c3 := BasicComponent{
		Natural:3,
		Stretched: 2,
		Shrinked: 1,
	}
	comps := []Component{&c1, &c2, &c3}
	s := NewSimpleCompositor()
	t:= NewTeXCompositor()

	simple := NewComposition(&s, comps)
	tex := NewComposition(&t, comps)

	simple.Repair()
	// 9

	tex.Repair()
	// 18
}
