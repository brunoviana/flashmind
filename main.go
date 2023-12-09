package main

import (
	"brunoviana/flashmind/kernel"
	"brunoviana/flashmind/kernel/interval"
	"fmt"
)

func main() {
	card := kernel.Card{Repetition: 0, Interval: 1, EaseFactor: 2.5}

	intervalCalculator := interval.IntervalCalculator{
		Calculate: interval.FourGradesIntervalCalculator,
	}

	intervalCalculator.Calculate(&card, interval.FourGradesRecalled)

	fmt.Printf(
		"Repetition: %d, Interval: %d days, Ease Factor: %.2f\n",
		card.Repetition,
		card.Interval,
		card.EaseFactor,
	)

	intervalCalculator.Calculate(&card, interval.FourGradesRecalled)

	fmt.Printf(
		"Repetition: %d, Interval: %d days, Ease Factor: %.2f\n",
		card.Repetition,
		card.Interval,
		card.EaseFactor,
	)
}
