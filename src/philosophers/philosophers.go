package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type ChopStick struct {
	sync.Mutex
}
type Philosopher struct {
	id    int
	left  *ChopStick
	right *ChopStick
}

func (p Philosopher) Eat(ch chan *Philosopher, maxNumOfMeals int, wg *sync.WaitGroup) {
	for t := 0; t < maxNumOfMeals; t++ {
		ch <- &p

		p.left.Lock()
		p.right.Lock()
		fmt.Printf("Starting eating: %d (luck)\n", p.id)

		duration := time.Duration(rand.Intn(1000))
		time.Sleep(duration * time.Millisecond)

		fmt.Printf("Finished eating: %d (%d ms.)\n", p.id, duration)
		p.right.Unlock()
		p.left.Unlock()
		wg.Done()
	}
}

func Host(ch chan *Philosopher) {
	for {
		if len(ch) == cap(ch) {
			<-ch
			<-ch
		}
	}
}

func main() {
	ChopSticks := make([]*ChopStick, 5)
	for i := 0; i < 5; i++ {
		ChopSticks[i] = new(ChopStick)
	}

	Philosophers := make([]*Philosopher, 5)
	for i := 0; i < 5; i++ {
		Philosophers[i] = &Philosopher{i + 1, ChopSticks[i], ChopSticks[(i+1)%5]}
	}

	maxNumOfMeals := 3
	maxEatingPhilosophers := 2

	var wg sync.WaitGroup
	wg.Add(maxNumOfMeals * len(Philosophers))
	ch := make(chan *Philosopher, maxEatingPhilosophers)

	go Host(ch)
	for i := 0; i < 5; i++ {
		go Philosophers[i].Eat(ch, maxNumOfMeals, &wg)
	}
	wg.Wait()
}
