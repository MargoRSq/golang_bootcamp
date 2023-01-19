package main

import (
	"container/heap"
	"errors"
	"fmt"
	"log"
)

type Present struct {
	Value int
	Size  int
}

type PresentHeap struct {
	Presents []*Present
}

func (p *PresentHeap) Len() (res int) {
	for range p.Presents {
		res++
	}
	return
}

func (p *PresentHeap) Less(i int, j int) (result bool) {
	if i < p.Len() || j < p.Len() {
		if p.Presents[i].Value == p.Presents[j].Value {
			result = p.Presents[j].Size <= p.Presents[i].Size
		} else {
			result = p.Presents[i].Value < p.Presents[j].Value
		}
		return
	} else {
		log.Fatalln("Index out of range!")
	}
	return
}

func (p *PresentHeap) Swap(i int, j int) {
	if i < p.Len() || j < p.Len() {
		p.Presents[i], p.Presents[j] = p.Presents[j], p.Presents[i]
	} else {
		log.Fatalln("Index out of range!")
	}
}

func (p *PresentHeap) Pop() (elem any) {
	elem = p.Presents[p.Len()-1]
	p.Presents = p.Presents[:p.Len()-1]
	return
}

func (p *PresentHeap) Push(x any) {
	p.Presents = append(p.Presents, x.(*Present))
	if !p.isSorted() {
		p.sort()
	}
}

func (p *PresentHeap) sort() {
	for i := range p.Presents {
		for j := range p.Presents {
			if !p.Less(i, j) {
				p.Swap(i, j)
			}
		}
	}
}

func (p *PresentHeap) isSorted() (result bool) {
	for i := 0; i < p.Len()-1; i++ {
		if p.Less(i, i+1) {
			return
		}
	}
	result = true
	return
}

func getNCoolestPresents(presents []*Present, n int) (coolest PresentHeap, err error) {
	if n < 0 {
		err = errors.New("error: n less than 0")
		return
	}
	if n > len(presents) {
		err = errors.New("error: n too big")
		return
	}
	coolest = PresentHeap{}
	for _, pr := range presents {
		coolest.Push(pr)
	}
	for coolest.Len() > n {
		coolest.Pop()
	}
	return
}

func (p PresentHeap) printHeap() {
	for count, pr := range p.Presents {
		fmt.Println(count, " -> value:", pr.Value, "size:", pr.Size)
	}
}

func printSlice(p []*Present) {
	for count, pr := range p {
		fmt.Println(count, " -> value:", pr.Value, "size:", pr.Size)
	}
}

func main() {
	lol := PresentHeap{Presents: []*Present{{3, 1}, {4, 5}, {5, 2}}}
	heap.Init(&lol)
	lol.Pop()
	lol.Push(&Present{5, 1})

	sl, _ := (getNCoolestPresents(lol.Presents, 2))
	printSlice(sl.Presents)
}
