package main

type ProductsIterator struct {
	index    int
	products []*Product
}

func (p *ProductsIterator) hasNext() bool {
	return p.index < len(p.products)
}

func (p *ProductsIterator) getNext() *Product {
	if p.hasNext() {
		product := p.products[p.index]
		p.index++
		return product
	}
	return nil
}
