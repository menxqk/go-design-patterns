package structural

import (
	"testing"
)

func TestCompositeSwimmerA_Swim(t *testing.T) {
	localSwim := Swim
	swimmer := CompositeSwimmerA{
		MySwim: &localSwim,
	}
	swimmer.MyAthlete.Train()
	(*swimmer.MySwim)()
}

func TestCompositeSwimmerB_Swim(t *testing.T) {
	swimmer := CompositeSwimmerB{
		Athlete{},
		&SwimmerImplementor{},
	}
	swimmer.Train()
	swimmer.Swim()
}

func TestAnimal_Swim(t *testing.T) {
	fish := Shark{
		Swim: Swim,
	}
	fish.Eat()
	fish.Swim()
}
