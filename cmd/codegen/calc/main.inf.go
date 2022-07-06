package main

type calculatorImpl interface {
	add(x int64, y int64) (z int64)
	sub(x int64, y int64) (z int64)
	multiply(x int64, y int64) (z int64)
	divide(x int64, y int64) (z int64)
}
