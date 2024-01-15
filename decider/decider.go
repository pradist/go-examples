package decider

type Decider struct {
	RandImpl func(x int) int
}

func (d *Decider) Decide(randRange int) int {
	randVal := d.RandImpl(randRange)
	return randVal
}

func (d *Decider) DecideWithRange(min, max int) int {
	randVal := d.RandImpl(max - min + 1)
	return randVal + min
}
