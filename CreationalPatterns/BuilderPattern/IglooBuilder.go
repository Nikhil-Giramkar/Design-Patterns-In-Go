package main

type IglooBuilder struct {
	windowType string
	doorType   string
	floors     int
}

func getIglooBuilder() *IglooBuilder {
	return &IglooBuilder{}
}

func (n *IglooBuilder) setWindowType() {
	n.windowType = "Ice Window"
}

func (n *IglooBuilder) setDoorType() {
	n.doorType = "Snow Door"
}

func (n *IglooBuilder) setFloors() {
	n.floors = 1
}

func (n *IglooBuilder) getHouse() House {
	return House{
		windowType: n.windowType,
		doorType:   n.doorType,
		floors:     n.floors,
	}
}
