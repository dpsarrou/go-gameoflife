package main

import (
	"context"
	"flag"
	"fmt"
	"time"

	"go-gameoflife/gameoflife"
)

func main() {
	// allow user to select random seed or glider pattern seed from cli
	useGlider := flag.Bool("glider", false, "Seed the universe with a 25x25 grid having the glider pattern in the middle")
	flag.Parse()
	u := gameoflife.NewRandomUniverse(25, 25)
	if *useGlider {
		u = gameoflife.NewGliderUniverse()
	}

	// run the engine
	e := gameoflife.Engine{}
	evolutions := e.Run(u, context.Background(), 10)
	for {
		u, open := <-evolutions
		if !open {
			break
		}
		fmt.Print(u.ToString())
		// smoother the updates to 30hz
		time.Sleep(time.Second * 1/30)
	}
}
