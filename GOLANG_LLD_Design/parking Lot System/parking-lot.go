package main

type ParkingLot struct {
	paymentProcessor
	ticketingSystem
	entryPoints int
	exitPoints  int
}

type Level struct {
	name         string
	parkingSpots map[string]int
}

type ParkingSpot struct {
}

type paymentProcessor interface {
	ProcessPayment(ticket Ticket, amount uint) error
}

type ticketingSystem struct {
	tickets map[string]Ticket
}

type Ticket struct{}
