package main

import (
	"context"
	"log"
)

func main() {
	ctx := context.Background()
	tools := NewOTELTools(ctx, "testtools")
	calc := NewCalculator(tools)
	defer tools.Cleanup()
	log.Println("the answer is", calc.add(1, 1))
	log.Println("the answer is", calc.sub(10, 1))
	log.Println("the answer is", calc.multiply(1, 10))
	log.Println("the answer is", calc.divide(100, 10))
}

type Calculator struct {
	calculatorImpl
	*OTELTools
}

func NewCalculator(tools *OTELTools) *Calculator {

	calculator := calculatorSpanner{
		OTELTools: tools,
		next:      &calculator{},
	}
	return &Calculator{
		calculatorImpl: &calculator,
	}
}

var _ calculatorImpl = &calculator{}

type calculator struct {
}

func (*calculator) add(x int64, y int64) (z int64) {
	//add business logic goes here
	return x + y
}

func (*calculator) sub(x int64, y int64) (z int64) {
	//sub business logic goes here
	return x - y
}

func (*calculator) multiply(x int64, y int64) (z int64) {
	//multiply business logic goes here
	return x * y
}

func (*calculator) divide(x int64, y int64) (z int64) {
	//divide business logic goes here
	return x / y
}
