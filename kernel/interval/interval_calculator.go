package interval

import "brunoviana/flashmind/kernel"

type GradeValue int64

type IntervalCalculatorStrategy interface {
	calculate(*kernel.Card, GradeValue) error
}

type IntervalCalculator struct {
	CalculatorStrategy IntervalCalculatorStrategy
}

func (intervalCalculator *IntervalCalculator) CalculateInterval(card *kernel.Card, grade GradeValue) error {
	return intervalCalculator.CalculatorStrategy.calculate(card, grade)
}
