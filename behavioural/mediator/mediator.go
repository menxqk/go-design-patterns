package main

import "fmt"

// The Mediator design pattern is used to avoid tight coupling between objects
// and reduce the amount of dependencies of a particular type.
type Mediator interface {
	Sum(a, b interface{}) interface{}
}

type SumMediator struct{}

func (s *SumMediator) Sum(a, b interface{}) interface{} {
	switch a := a.(type) {
	case int:
		switch b := b.(type) {
		case int:
			return a + b
		case float64:
			return float64(a) + b
		default:
			return fmt.Errorf("could not define number type of b: %v", b)
		}
	case float64:
		switch b := b.(type) {
		case int:
			return a + float64(b)
		case float64:
			return a + b
		default:
			return fmt.Errorf("could not define number type of b: %v", b)
		}
	default:
		return fmt.Errorf("could not define number type of a: %v", a)
	}
}

func main() {
	useSumMediator(4, 7)
	useSumMediator(10, 45.333)
	useSumMediator(89.324, 123.67)
}

func useSumMediator(a, b interface{}) {
	sm := SumMediator{}
	result := sm.Sum(a, b)

	err, ok := result.(error)
	if ok {
		panic(err)
	}

	fmt.Printf("%v [%T] + %v [%T] = %v [%T]\n", a, a, b, b, result, result)
}
