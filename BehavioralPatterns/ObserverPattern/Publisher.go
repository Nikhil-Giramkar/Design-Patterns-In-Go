package main
type Publisher interface {
	register(subscriber Subscriber)
	deregister(subscriber Subscriber)
	notifyAll()
}