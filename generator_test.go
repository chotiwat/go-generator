package generator

import (
	"fmt"
	"testing"
)

func fib(yield YieldFunc) {
	previous, current := 0, 1
	yield(0)
	yield(1)
	for {
		previous, current = current, previous+current
		yield(current)
	}
}

func TestGenerator(t *testing.T) {
	g := MakeGenerator(func(yield YieldFunc) {
		for i := 0; i < 10; i++ {
			yield(i * i)
		}
	})
	for {
		if val, ok := g.Next(); ok {
			fmt.Println(val)
		} else {
			break
		}
	}
}

func TestStop(t *testing.T) {
	g := MakeGenerator(fib)
	for i := 0; i < 10; i++ {
		fmt.Println(g.Next())
	}
	g.Stop()
}

func TestNextAfterClose(t *testing.T) {
	g := MakeGenerator(func(yield YieldFunc) {
		for i := 0; i < 10; i++ {
			yield(i * i)
		}
	})
	for i := 0; i < 11; i++ {
		fmt.Println(g.Next())
	}
}
