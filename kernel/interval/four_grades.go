package interval

import (
	"brunoviana/flashmind/kernel"
	"errors"
)

const (
	FourGradesForgot       GradeValue = 1
	FourGradesRecalledHard GradeValue = 2
	FourGradesRecalled     GradeValue = 3
	FourGradesRecalledEasy GradeValue = 4

	FourGradeMaxGrade int = 4
)

func FourGradesIntervalCalculator(card *kernel.Card, grade GradeValue) error {

	newEaseFactor, err := calculateEaseFactor(card.EaseFactor, FourGradeMaxGrade, int(grade))

	if err != nil {
		return errors.New("An error occured to calculate ease factor")
	}

	if grade == FourGradesForgot {
		card.Interval = 1
		card.Repetition = 0
	} else if card.Repetition == 3 && card.Interval == 1 {
		// this means the card is stuck due to many forget, so we force new interval
		card.Interval = 2
	} else {
		card.Interval = int(float64(card.Interval) * card.EaseFactor)
	}

	if newEaseFactor < 1.3 {
		newEaseFactor = 1.3
	}

	card.EaseFactor = newEaseFactor

	return nil
}

// implement easy bonus
