package main

type Subscriber interface {
	update(string)
	getId() string
}