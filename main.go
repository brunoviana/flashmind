package main

import (
	"brunoviana/flashmind/kernel"
	"brunoviana/flashmind/kernel/interval"
	"fmt"
)

func main() {
	card := kernel.Card{Repetition: 0, Interval: 1, EaseFactor: 2.5}

	fourGradesCalculator := &interval.FourGradesIntervalCalculator{}

	intervalCalculator := interval.IntervalCalculator{
		CalculatorStrategy: fourGradesCalculator,
	}

	intervalCalculator.CalculateInterval(&card, interval.FourGradesRecalled)

	fmt.Printf(
		"Repetition: %d, Interval: %d days, Ease Factor: %.2f\n",
		card.Repetition,
		card.Interval,
		card.EaseFactor,
	)

	intervalCalculator.CalculateInterval(&card, interval.FourGradesRecalled)

	fmt.Printf(
		"Repetition: %d, Interval: %d days, Ease Factor: %.2f\n",
		card.Repetition,
		card.Interval,
		card.EaseFactor,
	)
}
