package course3week4

import (
	"fmt"
	"sync"
	//"time"
)

var wg sync.WaitGroup
var doneCh chan int

type ChopStick struct {
	mutex sync.Mutex
}

type Philospher struct {
	philospherNum  int
	leftChopStick  *ChopStick
	rightChopStick *ChopStick
}

func (p *Philospher) Eat(simpleOrdering bool) {
	for i := 0; i < 5; i++ {
		if simpleOrdering {
			p.leftChopStick.mutex.Lock()
			p.rightChopStick.mutex.Lock()
			fmt.Printf("Starting to eat %v \n", p.philospherNum)
			p.rightChopStick.mutex.Unlock()
			p.leftChopStick.mutex.Unlock()
			fmt.Printf("Finishing eating %v \n", p.philospherNum)
		} else {
			p.rightChopStick.mutex.Lock()
			p.leftChopStick.mutex.Lock()
			fmt.Printf("Starting to eat %v \n", p.philospherNum)
			p.leftChopStick.mutex.Unlock()
			p.rightChopStick.mutex.Unlock()
			fmt.Printf("Finishing eating %v \n", p.philospherNum)
		}
	}
	wg.Done()
}
func Question_1() {
	chopSticks := make([]*ChopStick, 5)
	for i := 0; i < 5; i++ {
		chopSticks[i] = new(ChopStick)
	}

	philosphers := make([]*Philospher, 5)
	for i := 0; i < 5; i++ {
		philosphers[i] = &Philospher{
			philospherNum:  i + 1,
			leftChopStick:  chopSticks[i],
			rightChopStick: chopSticks[(i+1)%5],
		}
	}

	doneCh = make(chan int)
	go Host(philosphers)

	<-doneCh
}

func Host(philosphers []*Philospher) {
	i := 0
	simpleOrdering := true
	for {
		wg.Add(1)
		go philosphers[i].Eat(simpleOrdering)
		if i+3 > 4 {
			wg.Wait()
			break
		}
		wg.Add(1)
		go philosphers[i+3].Eat(simpleOrdering)
		simpleOrdering = !simpleOrdering
		i++
		wg.Wait()
	}
	doneCh <- 3
}
