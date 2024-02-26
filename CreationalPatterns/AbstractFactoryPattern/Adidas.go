package main

//Adidas derives ISportsFactory
type Adidas struct {
}

//Derives IShoesFactory
type AdidasShoe struct {
	Shoe
}

//Derives IShirtFactory
type AdidasShirt struct {
	Shirt
}

func (n *Adidas) makeShoes() IShoeFactory {
	return &AdidasShoe{
		Shoe: Shoe{
			lace: "Green",
			sole: "Yellow",
		},
	}
}

func (n *Adidas) makeShirts() IShirtFactory {
	return &AdidasShirt{
		Shirt: Shirt{
			size: 8,
			logo: "Adidas",
		},
	}
}
