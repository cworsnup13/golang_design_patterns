package main

import "fmt"

type Component interface {
	Draw()
	Resize()
}

type VisualComponent struct {
}

func NewVisualComponent() Component {
	return &VisualComponent{}
}

func (v *VisualComponent) Draw() {
	fmt.Println("A Component was drawn")
}

func (v *VisualComponent) Resize() {

}

type BorderDecorator struct {
	component Component
	width     int
}

func NewBorderDecorator(width int, comp Component) Component {
	return &BorderDecorator{
		component: comp,
		width:     width,
	}
}

func (d *BorderDecorator) Draw() {
	d.component.Draw()
	d.DrawBorder(d.width)
}
func (d *BorderDecorator) Resize() {
	d.component.Resize()
}

func (d *BorderDecorator) DrawBorder(width int) {
	fmt.Println("A Border was drawn")
}

type ScrollDecorator struct {
	component Component
	height    int
}

func NewScrollDecorator(height int, comp Component) Component {
	return &ScrollDecorator{
		component: comp,
		height:    height,
	}
}

func (d *ScrollDecorator) Draw() {
	d.component.Draw()
	d.DrawScroll(d.height)
}
func (d *ScrollDecorator) Resize() {
	d.component.Resize()
}

func (d *ScrollDecorator) DrawScroll(height int) {
	fmt.Println("A Scroll was drawn")
}

func main() {
	var v1 Component
	v1 = NewVisualComponent()

	decorators := []func(int, Component) Component{
		NewBorderDecorator,
		NewScrollDecorator,
	}

	for _, v := range decorators {
		v1 = v(5, v1)
	}

	v1.Draw()
}
