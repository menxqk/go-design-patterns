package structural

import "fmt"

// The Composite design pattern is used to simplify the use of common behavior.
// A method may be shared among different objects with the use of an interface
// or a common method signature.
func Swim() {
	fmt.Println("Swimming")
}

type Swimmer interface {
	Swim()
}

// -----------------------------------------
type Athlete struct{}

func (a *Athlete) Train() {
	fmt.Println("Training")
}

type SwimmerImplementor struct{}

func (s *SwimmerImplementor) Swim() {
	fmt.Println("Swimming!")
}

type Animal struct{}

func (a *Animal) Eat() {
	fmt.Println("Eating")
}

// -----------------------------------------
type CompositeSwimmerA struct {
	MyAthlete Athlete
	MySwim    *func()
}

type CompositeSwimmerB struct {
	Athlete
	Swimmer
}

type Shark struct {
	Animal
	Swim func()
}
