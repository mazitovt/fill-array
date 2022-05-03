package main

import (
	"fmt"
	fillarray "github.com/mazitovt/fill-array"
	"github.com/mazitovt/fill-array/generators"
	"github.com/mazitovt/fill-array/matrixstore"
	"github.com/mazitovt/fill-array/producer"
	"log"
	"os"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}

func run() error {
	w, _ := matrixstore.NewMatrixStore[int64](5, 5)

	c, next := generators.NewInt64Randomizer(1, 25)

	r := producer.NewProducer(true, c, next)

	if err := fillarray.FullFill[int64](w, r); err != nil {
		return err
	}

	fmt.Println(w)

	return nil
}
