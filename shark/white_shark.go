package shark

import (
	"fmt"
	"functional/prey"
	"functional/simulator"
	"math"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

type whiteShark struct {
	// speed in m/s
	speed float64
	// position of the shark in the map of 500 * 500 meters
	position [2]float64
	// simulator
	simulator simulator.CatchSimulator
}

func (w *whiteShark) Configure(position [2]float64, speed float64) {
	w.position = position
	w.speed = speed
}

func (w *whiteShark) Hunt(prey prey.Prey) (error, float64) {
	ok, timeToCatch := w.simulator.Catch(w.simulator.GetLinearDistance(w.position), w.speed, prey.GetSpeed())
	if !ok {
		return fmt.Errorf("could not catch it"), 0
	}
	return nil, math.Floor(timeToCatch)
}

func CreateWhiteShark(simulator simulator.CatchSimulator) Shark {
	return &whiteShark{
		simulator: simulator,
	}
}
