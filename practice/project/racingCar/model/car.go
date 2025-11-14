package model

import "strings"

const moveLimit = 4

type Car struct {
	name     *CarName
	position int
}

func NewCar(name string) (*Car, error) {
	carName, err := NewCarName(name)
	if err != nil {
		return nil, err
	}

	return &Car{
		name:     carName,
		position: 0,
	}, nil
}

func (c *Car) Name() string {
	return c.name.Value()
}

func (c *Car) Position() int {
	return c.position
}

func (c *Car) PositionString() string {
	if c.position <= 0 {
		return ""
	}
	return strings.Repeat("-", c.position)
}

func (c *Car) MoveForward(randomValue int) {
	if c.canMove(randomValue) {
		c.position++
	}
}

func (c *Car) canMove(randomValue int) bool {
	return randomValue >= moveLimit
}
