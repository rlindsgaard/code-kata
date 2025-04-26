package main

import (
	"errors"
	"fmt"
)

type State int

const (
	Locked State = iota
	Unlocked
	Start
	Stop
)

type Lock struct {
	code          string
	digitsEntered string
	state         State
	Prev          *Lock
}

func (s *State) String() string {
	switch *s {
	case Locked:
		return "Locked"
	case Unlocked:
		return "Unlocked"
	case Start:
		return "Start"
	case Stop:
		return "Stop"
	default:
		return "Unknown"
	}
}

func (l *Lock) String() string {
	return fmt.Sprintf("Lock{code: %s, digitsEntered: %s, state: %s}", l.code, l.digitsEntered, &l.state)
}

func NewLock(code string) (*Lock, error) {
	return &Lock{
		code:          code,
		digitsEntered: "",
		state:         Start,
		Prev:          nil,
	}, nil
}

func (l *Lock) Lock() (*Lock, error) {
	switch l.state {
	case Locked:
		return l, fmt.Errorf("lock is already locked")
	case Unlocked, Start:
		return &Lock{
			code:          l.code,
			digitsEntered: l.digitsEntered,
			state:         Locked,
			Prev:          l,
		}, nil
	default:
		return l, fmt.Errorf("invalid state")
	}
}

func (l *Lock) EnterDigit(digit string) (*Lock, error) {
	fmt.Printf("Entered digit: %s\n", digit)
	switch l.state {
	case Locked:
		digitsEntered := l.digitsEntered + digit
		state := l.state
		if digitsEntered == l.code {
			state = Unlocked
		} else if len(digitsEntered) >= 4 {
			digitsEntered = ""
		}
		return &Lock{
			code:          l.code,
			digitsEntered: digitsEntered,
			state:         state,
			Prev:          l,
		}, nil
	case Unlocked:
		return &Lock{
			code:          l.code,
			digitsEntered: "",
			state:         Locked,
			Prev:          l,
		}, fmt.Errorf("lock is unlocked")
	default:
		return l, fmt.Errorf("invalid state")
	}
}

func (l *Lock) Stop() (*Lock, error) {
	var err error = nil
	if l.state == Stop {
		err = errors.New("already stopped")
	}
	return &Lock{
		code:          l.code,
		digitsEntered: l.digitsEntered,
		state:         Stop,
		Prev:          l,
	}, err
}

func main() {
	lock, err := NewLock("1234")
	if err != nil {
		fmt.Println("Could not create lock", err)
	}
	lock, err = lock.Lock()
	if err != nil {
		fmt.Println("Could not lock", err)
	}
	lock, err = lock.EnterDigit("1")
	if err != nil {
		fmt.Println("Error from lock", err)
	}
	fmt.Printf("Current state: %s\n", lock)
	lock, err = lock.EnterDigit("2")
	if err != nil {
		fmt.Println("Error from lock", err)
	}
	fmt.Printf("Current state: %s\n", lock)
	lock, err = lock.EnterDigit("3")
	if err != nil {
		fmt.Println("Error from lock", err)
	}
	fmt.Printf("Current state: %s\n", lock)
	lock, err = lock.EnterDigit("4")
	if err != nil {
		fmt.Println("Error from lock", err)
	}
	fmt.Printf("Current state: %s\n", lock)
	lock, err = lock.EnterDigit("5")
	if err != nil {
		fmt.Println("Error from lock", err)
	}
	fmt.Printf("Current state: %s\n", lock)

	lock, err = lock.EnterDigit("4")
	if err != nil {
		fmt.Println("Error from lock", err)
	}
	fmt.Printf("Current state: %s\n", lock)
	lock, err = lock.EnterDigit("3")
	if err != nil {
		fmt.Println("Error from lock", err)
	}
	fmt.Printf("Current state: %s\n", lock)
	lock, err = lock.EnterDigit("2")
	if err != nil {
		fmt.Println("Error from lock", err)
	}
	fmt.Printf("Current state: %s\n", lock)
	lock, err = lock.EnterDigit("1")
	if err != nil {
		fmt.Println("Error from lock", err)
	}
	fmt.Printf("Current state: %s\n", lock)

	lock, err = lock.Stop()
	if err != nil {
		fmt.Println("Error from lock", err)
	}
	fmt.Printf("Current state: %s\n", lock)
	lock, err = lock.Stop()
	if err != nil {
		fmt.Println("Error from lock", err)
	}
	fmt.Printf("Current state: %s\n", lock)
	lock, err = lock.Lock()
	if err != nil {
		fmt.Println("Error from lock", err)
	}
	fmt.Printf("Current state: %s\n", lock)
	lock, err = lock.EnterDigit("1")
	if err != nil {
		fmt.Println("Error from lock", err)
	}
	fmt.Printf("Current state: %s\n", lock)
}
