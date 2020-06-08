package main

import (
	"fmt"
	"math/rand"
	"time"
)

type SkipList struct {
	// Linked list of nodes, each of which is itself a linked list.
	//                     c
	//                     |
	//           b   ->    c                   e
	//           |         |                   |
	// a   ->    b   ->    c   ->    d   ->    e
	head *node
	// next *SkipList
}

func init() {
	rand.Seed(time.Now().Unix())
}

type SL struct {
	top node
}

func (sl *SL) Include(v int) bool {
	n := &sl.top

}

func (sl *SL) Add(v int) {
	var stk []*node
	stk = append(stk, nil)
	n := &sl.top
	lvl := 0
	for n != nil {
		n = n.find(v)
		// fmt.Printf("inserting %d, found node %+v at lvl %d\n", v, n, lvl)
		fmt.Printf("inserting %d, found node %d at lvl %d\n", v, n.value, lvl)
		stk = append(stk, n)
		n = n.down
		lvl++
	}
	fmt.Printf("inserting with stk\n")
	for i, n := range stk {
		fmt.Printf("  %d: %+v\n", i, n)
	}
	var d *node
	for i := len(stk) - 1; i >= 0; i-- {
		n := stk[i]
		if n == nil {
			fmt.Printf("inserting %d, adding a new top layer\n", v)
			// We are adding a new top layer.
			ot := sl.top
			sl.top = node{
				right: &node{
					value: v,
					down:  d,
				},
				down: &ot,
			}
			// fmt.Printf("inserting %d, sl.top %+v\n", v, sl.top)
			// fmt.Printf("inserting %d, sl.top.down %+v\n", v, sl.top.down)
			break
		}
		n.right = &node{
			value: v,
			right: n.right,
			down:  d,
		}
		fmt.Printf("inserted %d, flipping a coin at level %d...", v, i)
		if rand.Intn(2) == 0 {
			fmt.Printf("tails! breaking\n")
			break
		}
		fmt.Printf("heads! looping\n")
		sl.Print()
		d = n.right
	}
}

func (sl *SL) Print() {
	n := &sl.top
	max := 10
	lvl := 0
	for n != nil {
		fmt.Printf("%d: ", lvl)
		lvl++
		n.print()
		n = n.down
		max--
		if max < 0 {
			return
		}
	}
}

type node struct {
	value int
	right *node
	down  *node
}

func (n node) String() string {
	return fmt.Sprintf("<%d right: %+v down: %+v>", n.value, n.right, n.down)
}

func (n *node) find(v int) *node {
	for n.right != nil && n.right.value < v {
		n = n.right
	}
	return n
}

func (n *node) add(v int) *node {
	prev := n
	for prev.right != nil && prev.right.value < v {
		prev = prev.right
	}
	prev.right = &node{
		right: prev.right,
		value: v,
	}
	if rand.Intn(2) == 0 {
		return prev.right
	}
	return nil
}

func (n *node) print() {
	max := 100000
	for n != nil {
		fmt.Printf("%d ", n.value)
		n = n.right
		max--
		if max == 0 {
			break
		}
	}
	fmt.Println()
}

type ins struct {
	row   *SkipList
	after *node
}

func (s *SkipList) find(v int) *node {
	n := s.head
	for n.right != nil && n.right.value < v {
		n = n.right
	}
	if n.right == nil {
		return n
	}
	for n != nil && v >= n.value {
		n = n.right
	}
	return n
}

func (s *SkipList) Add(v int) {
	n := s.find(v)
	if n == nil {
		s.head = &node{
			value: v,
		}
	} else {
		or := n.right
		n.right = &node{
			value: v,
		}
		n.right.right = or
	}

	if s.head == nil {
		s.head = &node{
			value: v,
		}
		return
	}
	// var inss []ins
	s.head = s.head.Add(v)
}

func (n *node) Add(v int) *node {
	return n // TODO
}

func main() {
	xs := make([]int, 0)
	if xs == nil {
		fmt.Println("nil list")
	}
	fmt.Println("HI")
}
