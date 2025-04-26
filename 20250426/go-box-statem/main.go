package main

import (
	"errors"
	"fmt"
)

type BoxState int

const (
	Opened = iota
	Closed
)

type Box struct {
	state BoxState
}

func (b BoxState) String() string {
	switch b {
	case Opened:
		return "Opened"
	case Closed:
		return "Closed"
	default:
		return "Unknown"
	}
}

func NewBox() *Box {
	return &Box{
		state: Closed,
	}
}

func (b *Box) Open() (BoxState, error) {
	if b.state == Opened {
		return b.state, errors.New("Box is already opened")
	}
	if b.state == Closed {
		b.state = Opened
	}
	return b.state, nil
}

func (b *Box) Close() (BoxState, error) {
	if b.state == Closed {
		return b.state, errors.New("Box is already closed")
	}
	if b.state == Opened {
		b.state = Closed
	}
	return b.state, nil
}

func (b *Box) GetState() BoxState {
	return b.state
}

func main() {
	box := NewBox()
	_, err := box.Open()
	if err != nil {
		println(err.Error())
	}
	fmt.Printf("Box state after opening: %s\n", box.GetState())
	_, err = box.Close()
	if err != nil {
		println(err.Error())
	}
	fmt.Printf("Box state after closing: %s\n", box.GetState())
}
