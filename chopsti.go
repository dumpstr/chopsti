package main

import (
	"fmt"
	"sync"
	"time"
)

type ChopSti struct { sync.Mutex }

type Phil struct {
	leftCS, rightCS *ChopSti
	numb            int
}

func (p Phil) eat(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 3; i++ {
		p.leftCS.Lock()
		p.rightCS.Lock()

		fmt.Println("philosopher ", p.numb, "starting to eat...")
		time.Sleep(time.Second)
		fmt.Println("philosopher ", p.numb, "done eating: ", i+1)

		p.leftCS.Unlock()
		p.rightCS.Unlock()
	}
}

func main() {
	var wg sync.WaitGroup
	host := make(chan struct{}, 2)
	Philo := make([]*Phil, 5)
	Choppies := make([]*ChopSti, 5)

	for i := 0; i < 5; i++ {
		Choppies[i] = new(ChopSti)
	}

	for j := 0; j < 5; j++ {
		wg.Add(1)
		Philo[j] = &Phil{Choppies[j], Choppies[(j+1)%5], j + 1}
	}

	for k := 0; k < 5; k++ {
		host <- struct{}{}
		go func(n int) {
			Philo[n].eat(&wg)
			<-host
		}(k)
	}

	wg.Wait()
}
