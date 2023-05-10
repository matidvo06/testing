package prey

type tuna struct {
	// max speed in m/s
	speed float64
}

func (t *tuna) SetSpeed(speed float64) {
	t.speed = speed
}

func (t *tuna) GetSpeed() float64 {
	return t.speed
}

func CreateTuna() Prey {
	return &tuna{}
}
