package main

type ProductsCollection struct {
	products []*Product
}

func (p *ProductsCollection) createIterator() Iterator {
	return &ProductsIterator{
		products: p.products,
	}
}
