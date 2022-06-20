package main

import (
	"fmt"
	"time"
)

type Handler interface {
	Do(s string) string
}

type HandlerFunc func(string) string

func (d HandlerFunc) Do(s string) string {
	return d(s)
}

func toUpper(str string) HandlerFunc {
	fmt.Println("before to upper")
	return func(str string) string {
		defer fmt.Println("after to upper")
		return fmt.Sprintf("___________%s", str)
	}
}

func beforeAfterLogging(f HandlerFunc) HandlerFunc {
	fmt.Println("before prefix underscore")
	return func(str string) string {
		defer fmt.Println("after prefix underscore")
		return f(str)
	}
}

func withDelay(delay time.Duration, f HandlerFunc) HandlerFunc {
	fmt.Printf("sleeping for %s\n", delay)
	time.Sleep(delay)
	return func(str string) string {
		defer fmt.Println("after sleeping")
		return f(str)
	}
}

func main() {
	pf := beforeAfterLogging(withDelay(time.Second, toUpper("hello")))
	fmt.Println(pf("thisone"))
}

func HandleFunc(str string, handler HandlerFunc) string {
	return handler.Do(str)
}

func Handle(str string, handlers ...HandlerFunc) string {
	toTotal := str
	for _, handler := range handlers {
		toTotal = handler(toTotal)
	}
	return toTotal
}
