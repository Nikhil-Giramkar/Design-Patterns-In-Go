package main

import "fmt"

type Book struct {
	subscribersList []Subscriber
	name            string
	inStock         bool
}

func newBook(name string) *Book {
	return &Book{
		name: name,
	}
}

func (b *Book) register(s Subscriber) {
	b.subscribersList = append(b.subscribersList, s)
}

func (b *Book) notifyAll() {
	for _, subscriber := range b.subscribersList {
		subscriber.update(b.name)
	}
}

func removeFromSlice(subsciberList []Subscriber, subToRemove Subscriber) []Subscriber {
	subscibersCount := len(subsciberList)

	for i, subscriber := range subsciberList {
		if subToRemove.getId() == subscriber.getId() {
			//Swap current with last
			subsciberList[subscibersCount-1], subsciberList[i] = subsciberList[i], subsciberList[subscibersCount-1]
			//Remove last
			return subsciberList[:subscibersCount-1]
		}
	}

	return subsciberList
}

func (b *Book) deregister(subscriberToRemove Subscriber) {
	b.subscribersList = removeFromSlice(b.subscribersList, subscriberToRemove)
}

func (b *Book) updateAvailability() {
	fmt.Printf("\"%s\" is now in stock\n\n", b.name)
	b.inStock = true
	b.notifyAll()
}
