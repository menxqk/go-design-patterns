package example1

import "fmt"

// The Command design pattern is used to standardize actions and alllow their execution
// to be invoked from somewhere else.
type Command interface {
	Execute()
}

func CreateCommand(s string) Command {
	fmt.Println("Creating command")
	return &ConsoleOutput{message: s}
}

type ConsoleOutput struct {
	message string
}

func (c *ConsoleOutput) Execute() {
	fmt.Println(c.message)
}

type CommandQueue struct {
	queue []Command
}

func (q *CommandQueue) AddCommand(c Command) {
	q.queue = append(q.queue, c)
	if len(q.queue) == 3 {
		for _, command := range q.queue {
			command.Execute()
		}
		q.queue = make([]Command, 0)
	}
}

func (q *CommandQueue) Len() int {
	return len(q.queue)
}
