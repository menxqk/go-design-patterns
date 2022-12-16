package main

import (
	"flag"
	"log"
	"os"

	"github.com/menxqk/go-design-patterns/behavioural/strategy/shapes"
)

var output = flag.String("output", "text", "The output to use between 'text' and 'image'")

func main() {
	flag.Parse()

	activeStrategy, err := shapes.NewPrinter(*output)
	if err != nil {
		log.Fatal(err)
	}

	switch *output {
	case shapes.TEXT_STRATEGY:
		activeStrategy.SetWriter(os.Stdout)
	case shapes.IMAGE_STRATEGY:
		w, err := os.Create("/tmp/image.jpg")
		if err != nil {
			log.Fatal("error opening image")
		}
		defer w.Close()
		activeStrategy.SetWriter(w)
	}

	err = activeStrategy.Print()
	if err != nil {
		log.Fatal(err)
	}
}
