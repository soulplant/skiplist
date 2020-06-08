package main

import (
	"math/rand"
	"testing"
)

func TestAdd(t *testing.T) {
	sl := &SL{}
	for i := 0; i < 10; i++ {
		// sl.Add((i + 1) * 20)
		sl.Add(1 + rand.Intn(100))
		sl.Print()
	}

	t.Errorf("hi")
}
