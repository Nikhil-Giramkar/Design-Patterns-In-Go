package main

type NormalBuilder struct {
	windowType string
	doorType   string
	floors     int
}

func getNormalBuilder() *NormalBuilder {
	return &NormalBuilder{}
}

func (n *NormalBuilder) setWindowType() {
	n.windowType = "Glass Window"
}

func (n *NormalBuilder) setDoorType() {
	n.doorType = "Wooden Door"
}

func (n *NormalBuilder) setFloors() {
	n.floors = 4
}

func (n *NormalBuilder) getHouse() House {
	return House{
		windowType: n.windowType,
		doorType:   n.doorType,
		floors:     n.floors,
	}
}
