package main

import (
	"fmt"
)

type Equipment interface {
	Name() string
	Power() float64 // Watt
	NetPrice() float64
	DiscountPrice() float64
	Add(equipment Equipment)
	Remove(equipment Equipment) (removed bool)
}

type Power struct {
	watts float64
}

type Price struct {
	price    float64
	discount float64
}

type CompositeEquipment struct {
	name       string
	power      Power
	price      Price
	equipments []Equipment
}

func (c *CompositeEquipment) Name() string {
	return c.name
}

func (c *CompositeEquipment) Power() float64 {
	var total = c.power.watts
	for _, e := range c.equipments {
		total += e.Power()
	}
	return total
}

func (c *CompositeEquipment) NetPrice() float64 {
	var total = c.price.price
	for _, e := range c.equipments {
		total += e.NetPrice()
	}
	return total
}

func (c *CompositeEquipment) DiscountPrice() float64 {
	var total = c.price.discount
	for _, e := range c.equipments {
		total += e.DiscountPrice()
	}
	return total
}

func (c *CompositeEquipment) Add(equipment Equipment) {
	c.equipments = append(c.equipments, equipment)
}

func (c *CompositeEquipment) Remove(equipment Equipment) (removed bool) {
	var idx int
	for i := range c.equipments {
		if c.equipments[i] == equipment {
			removed = true
			idx = i
			break
		}
	}
	if removed {
		c.equipments = append(c.equipments[:idx], c.equipments[idx+1:]...)
	}
	return
}

type Chassis struct {
	Composite CompositeEquipment
}

func NewChassis(name string, watts, price, discount float64) Chassis {
	var f = CompositeEquipment{
		name:  name,
		power: Power{watts: watts},
		price: Price{price: price, discount: discount},
	}
	return Chassis{Composite: f}
}

func (c *Chassis) Name() string {
	return c.Composite.Name()
}

func (c *Chassis) Power() float64 {
	return c.Composite.Power()
}

func (c *Chassis) NetPrice() float64 {
	return c.Composite.NetPrice()
}

func (c *Chassis) DiscountPrice() float64 {
	return c.Composite.DiscountPrice()
}

func (c *Chassis) Add(equipment Equipment) {
	c.Composite.Add(equipment)
}

func (c *Chassis) Remove(equipment Equipment) (removed bool) {
	return c.Composite.Remove(equipment)
}

type Cabinet struct {
	Composite CompositeEquipment
}

func NewCabinet(name string, watts, price, discount float64) Cabinet {
	var f = CompositeEquipment{
		name:  name,
		power: Power{watts: watts},
		price: Price{price: price, discount: discount},
	}
	return Cabinet{Composite: f}
}

func (c *Cabinet) Name() string {
	return c.Composite.Name()
}

func (c *Cabinet) Power() float64 {
	return c.Composite.Power()
}

func (c *Cabinet) NetPrice() float64 {
	return c.Composite.NetPrice()
}

func (c *Cabinet) DiscountPrice() float64 {
	return c.Composite.DiscountPrice()
}

func (c *Cabinet) Add(equipment Equipment) {
	c.Composite.Add(equipment)
}

func (c *Cabinet) Remove(equipment Equipment) (removed bool) {
	return c.Composite.Remove(equipment)
}

type Bus struct {
	name  string
	power Power
	price Price
}

func NewBus(name string, watts, price, discount float64) Bus {
	return Bus{
		name:  name,
		power: Power{watts: watts},
		price: Price{price: price, discount: discount},
	}
}

func (b *Bus) Name() string {
	return b.name
}

func (b *Bus) Power() float64 {
	return b.power.watts
}

func (b *Bus) NetPrice() float64 {
	return b.price.price
}

func (b *Bus) DiscountPrice() float64 {
	return b.price.discount
}

func (b *Bus) Add(equipment Equipment) {
	return
}

func (b *Bus) Remove(equipment Equipment) (removed bool) {
	return false
}

func main() {
	cabinet := NewCabinet("PC Cabinet", 0, 200, 200)
	chassis := NewChassis("PC Chassis", 0, 200, 200)
	chassisUpgrade := NewCabinet("PC Chassis 2", 0, 300, 200)
	bus := NewBus("MCA Bus", 1, 50, 45)

	cabinet.Add(&chassis)

	fmt.Println("Build 1:", cabinet.NetPrice())

	// Replacement Example
	cabinet.Remove(&chassis)
	cabinet.Add(&chassisUpgrade)

	// Adding components to a lower level composite

	chassisUpgrade.Add(&bus)

	fmt.Println("Build 2:", cabinet.NetPrice())
	fmt.Println("Build 2 Discount:", cabinet.DiscountPrice())
}
