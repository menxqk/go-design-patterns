package example1

import "testing"

func TestCommandQueue(t *testing.T) {
	cq := CommandQueue{}

	cq.AddCommand(CreateCommand("First message"))
	cq.AddCommand(CreateCommand("Second message"))
	cq.AddCommand(CreateCommand("Third message"))
	cq.AddCommand(CreateCommand("Fourth message"))
	cq.AddCommand(CreateCommand("Fifth message"))

	if cq.Len() != 2 {
		t.Fatalf("expected 2 commands in queue, got %d", cq.Len())
	}

	cq.AddCommand(CreateCommand("Sixth message"))

	if cq.Len() != 0 {
		t.Fatalf("expected 0 commands in queue, got %d", cq.Len())
	}
}
