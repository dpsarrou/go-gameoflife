package gameoflife

import (
	"context"
	"testing"
	"time"
)

func TestEngineRunsUntilContextCancellation(t *testing.T) {
	e := Engine{}
	ctx, cancel := context.WithCancel(context.Background())
	evolutions := e.Run(NewRandomUniverse(9, 9), ctx, 10)

	// cancel the generation after some little amount of time
	go func() {
		<-time.After(20 * time.Millisecond)
		cancel()
	}()

	count := 0
	for {
		_, open := <-evolutions
		if !open {
			break
		}
		count++
	}

	if count < 1 {
		t.Errorf("Expected at least some generations")
	}
}
