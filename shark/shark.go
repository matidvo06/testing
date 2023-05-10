package shark

import (
	"functional/prey"
)

type Shark interface {
	Hunt(prey prey.Prey) (error, float64)
	Configure(position [2]float64, speed float64)
}
