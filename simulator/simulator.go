package simulator

import (
	"fmt"
	"math/big"
)

type CatchSimulator interface {
	Catch(distance float64, speed float64, catchSpeed float64) (bool, float64)
	GetLinearDistance(position [2]float64) float64
}

type catchSimulator struct {
	// max time to catch the prey in seconds
	maxTimeToCatch float64
}

func (r *catchSimulator) Catch(distance float64, speed float64, catchSpeed float64) (bool, float64) {
	fmt.Printf("Distance: (%.2f) | Prey Speed: %.2f m/s\n", distance, catchSpeed)
	timeToCatch := distance / (speed - catchSpeed)
	fmt.Printf("Time to catch: %.2f\n seconds", timeToCatch)
	return timeToCatch > 0 && timeToCatch <= r.maxTimeToCatch, timeToCatch
}

func (r *catchSimulator) GetLinearDistance(position [2]float64) float64 {
	x := big.NewFloat(position[0])
	y := big.NewFloat(position[1])
	z := x.Add(x.Mul(x, x), y.Mul(y, y))
	res, _ := z.Sqrt(z).Float64()
	fmt.Printf("Distance: %.2f meters\n", res)
	return res
}

func NewCatchSimulator(maxTimeToCatch float64) CatchSimulator {
	return &catchSimulator{
		maxTimeToCatch: maxTimeToCatch,
	}
}
