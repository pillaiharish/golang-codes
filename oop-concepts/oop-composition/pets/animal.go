package pets

import (
	"fmt"
	"time"
)

type Animal struct {
	sound   string
	lastAte time.Time
}

func (a Animal) Feed(food string) string {
	a.lastAte = time.Now()
	return fmt.Sprintf("The animal is eating %s", food)
}

func (a Animal) Sound(s string) string {
	a.sound = s
	return a.sound
}

func (a Animal) IsHungry() bool {
	// every 2 seconds my dog wants food
	// time.Sleep(3 * time.Second)
	return time.Since(a.lastAte) > 2*time.Second

}
