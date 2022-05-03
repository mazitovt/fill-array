package main

import (
	fillarray "github.com/mazitovt/fill-array"
	"github.com/mazitovt/fill-array/generators"
	"github.com/mazitovt/fill-array/matrixstore"
	"github.com/mazitovt/fill-array/producer"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestFullFillWithStrings(t *testing.T) {

	s := `hey
its
tim
wow`
	left, next := generators.NewStringsInfGenerator(strings.Split(s, "\n"))
	r := producer.NewProducer[string](false, left, next)
	w, _ := matrixstore.NewMatrixStore[string](5, 5)

	assert.Equal(t, nil, fillarray.FullFill[string](w, r))
}

func TestFullFillWithStringsUnique(t *testing.T) {

	s := `hey
its
tim
wow`
	left, next := generators.NewStringsInfGenerator(strings.Split(s, "\n"))
	r := producer.NewProducer[string](true, left, next)
	w, _ := matrixstore.NewMatrixStore[string](5, 5)

	assert.Equal(t, fillarray.ErrWriterIsNotFull, fillarray.FullFill[string](w, r))
}

func TestFullFillWithIntUnique(t *testing.T) {

	left, next := generators.NewInt64Randomizer(1, 25)
	r := producer.NewProducer[int64](true, left, next)
	w, _ := matrixstore.NewMatrixStore[int64](5, 5)

	assert.Equal(t, nil, fillarray.FullFill[int64](w, r))
}

func TestFullFillWithIntUniqueFail(t *testing.T) {

	left, next := generators.NewInt64Randomizer(1, 10)
	r := producer.NewProducer[int64](true, left, next)
	w, _ := matrixstore.NewMatrixStore[int64](5, 5)

	assert.Equal(t, fillarray.ErrWriterIsNotFull, fillarray.FullFill[int64](w, r))
}

func TestFullFillWithInt(t *testing.T) {

	left, next := generators.NewInt64Randomizer(1, 10)
	r := producer.NewProducer[int64](false, left, next)
	w, _ := matrixstore.NewMatrixStore[int64](5, 5)

	assert.Equal(t, nil, fillarray.FullFill[int64](w, r))
}

func TestFullFillWithInt2(t *testing.T) {

	left, next := generators.NewInt64Randomizer(1, 26)
	r := producer.NewProducer[int64](false, left, next)
	w, _ := matrixstore.NewMatrixStore[int64](5, 5)

	assert.Equal(t, nil, fillarray.FullFill[int64](w, r))
}
