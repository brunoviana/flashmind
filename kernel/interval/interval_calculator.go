package interval

import "brunoviana/flashmind/kernel"

type GradeValue int64

type IntervalCalculatorStrategy func(card *kernel.Card, grade GradeValue) error

type IntervalCalculator struct {
	Calculate IntervalCalculatorStrategy
}
