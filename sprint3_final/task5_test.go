package sprint3_final

import (
	"context"
	"testing"
	"time"
)

func TestFanIn(t *testing.T) {
	t.Run("default positive", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.TODO())
		defer cancel()
		numChan := 5
		channels := make([]<-chan int, numChan)
		for i := 0; i < numChan; i++ {
			ch := make(chan int)
			go func() {
				defer close(ch)
				for i := 0; i < 100; i++ {
					ch <- i
				}
			}()
			channels[i] = ch
		}
		numCount := make([]int, 100)
		for num := range FanIn(ctx, channels...) {
			numCount[num]++
		}
		for _, i := range numCount {
			if i != numChan {
				t.Errorf("not all values collected, expected count of %v to equal %v, got %v", i, numChan, numCount[i])
			}
		}
	})
	t.Run("with timeout", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.TODO(), time.Nanosecond*50)
		defer cancel()
		numChan := 5
		channels := make([]<-chan int, numChan)
		for i := 0; i < numChan; i++ {
			ch := make(chan int)
			go func() {
				defer close(ch)
				for i := 0; i < 100; i++ {
					ch <- i
				}
			}()
			channels[i] = ch
		}
		numCount := make([]int, 100)
		for num := range FanIn(ctx, channels...) {
			numCount[num]++
		}
		for _, i := range numCount {
			if i != numChan {
				return
			}
		}
		t.Errorf("context do not canceled")
	})
}
