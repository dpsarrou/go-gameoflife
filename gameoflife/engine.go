package gameoflife

import (
	"context"
)

// Engine models the game engine that can generate universe evolutions
type Engine struct{}

// Run continuously generates new universe evolutions based on the game rules.
// The results are asynchronously posted on a channel, that has a buffer length
// of channelSize. The calling code can use the Context's cancellation function
// to stop the execution.
// A simple usage:
//
//	e := Engine{}
//	e.Run(NewRandomUniverse(9,9), context.Background(), 10)
//
// See the engine_test.go file for example usage with context cancellation.
func (e *Engine) Run(u Universe, ctx context.Context, channelSize int) chan Universe {
	evolutions := make(chan Universe, channelSize)
	go func() {
		defer close(evolutions)
		for {
			select {
			case <-ctx.Done():
				return
			default:
				evolved := u.Evolve()
				evolutions <- evolved
				u = evolved
			}
		}
	}()
	return evolutions
}
