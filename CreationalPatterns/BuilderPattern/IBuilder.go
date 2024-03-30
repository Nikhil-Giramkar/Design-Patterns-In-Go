package main

type IBuilder interface {
	setWindowType()
	setDoorType()
	setFloors()
	getHouse() House
}

func getBuilder(houseType string) IBuilder {
	if houseType == "Normal" {
		return getNormalBuilder()
	} else if houseType == "Igloo" {
		return getIglooBuilder()
	}
	return nil
}
