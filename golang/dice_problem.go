package main

func DiceProbability(sides int) float64 {
	if sides <= 0 {
		return 0
	}
	return 1.0 / float64(sides)
}
