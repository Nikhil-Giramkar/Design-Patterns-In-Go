package main

//Nike derives ISportsFactory
type Nike struct {
}

//Derives IShoesFactory
type NikeShoe struct {
	Shoe
}

//Derives IShirtFactory
type NikeShirt struct {
	Shirt
}

func (n *Nike) makeShoes() IShoeFactory {
	return &NikeShoe{
		Shoe: Shoe{
			lace: "White",
			sole: "Blue",
		},
	}
}

func (n *Nike) makeShirts() IShirtFactory {
	return &NikeShirt{
		Shirt: Shirt{
			size: 7,
			logo: "Nike",
		},
	}
}
