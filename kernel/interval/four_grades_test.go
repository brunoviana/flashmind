package interval

import (
	"brunoviana/flashmind/kernel"
	"testing"
)

func TestFourGradesCalculatorShouldResetIntervalAndRepetitionIfCardWasForgotten(t *testing.T) {
	card := kernel.Card{Repetition: 5, Interval: 50, EaseFactor: 2.5}

	intervalCalculator := FourGradesIntervalCalculator{}
	intervalCalculator.calculate(&card, FourGradesForgot)

	if card.Interval > 1 {
		t.Errorf("expected to have card interval set to 1. Given value was: %d", card.Interval)
	}

	if card.Repetition > 0 {
		t.Errorf("expected to have card repetition set to 0. Given value was: %d", card.Repetition)
	}
}

func TestFourGradesCalculatorShouldNotSetEaseFactorLessThanOnePointThree(t *testing.T) {
	card := kernel.Card{Repetition: 5, Interval: 50, EaseFactor: 1.3}

	intervalCalculator := FourGradesIntervalCalculator{}
	intervalCalculator.calculate(&card, FourGradesForgot)

	if card.EaseFactor < 1.3 {
		t.Errorf("expected to have ease factor not less than 1.3. Given value was: %.2f", card.EaseFactor)
	}
}

func TestFourGradesCalculatorShouldSetIntervalEqualsTwoIfCardHasThreeRepetitionsAndOneInterval(t *testing.T) {
	card := kernel.Card{Repetition: 3, Interval: 1, EaseFactor: 1.3}

	intervalCalculator := FourGradesIntervalCalculator{}
	intervalCalculator.calculate(&card, FourGradesRecalled)

	if card.Interval != 2 {
		t.Errorf("expected to have interval equals 2. Given value was: %d", card.Interval)
	}
}

func TestFourGradesCalculatorShouldUseCurrentCardEaseFactorToCalculateNewInterval(t *testing.T) {
	card := kernel.Card{Repetition: 0, Interval: 1, EaseFactor: 2.5}

	intervalCalculator := FourGradesIntervalCalculator{}

	intervalCalculator.calculate(&card, FourGradesRecalled)
	if float64(card.Interval) != 2 {
		t.Errorf("expected new interval is equals 2. Given value was: %d", card.Interval)
	}

	intervalCalculator.calculate(&card, FourGradesRecalled)
	if float64(card.Interval) != 5 {
		t.Errorf("expected new interval is equals 5. Given value was: %d", card.Interval)
	}
}

func TestFourGradesCalculatorShouldDecreaseEaseFactorIfCardWasForgotten(t *testing.T) {
	card := kernel.Card{Repetition: 5, Interval: 50, EaseFactor: 2.5}

	oldEaseFactor := card.EaseFactor

	intervalCalculator := FourGradesIntervalCalculator{}
	intervalCalculator.calculate(&card, FourGradesForgot)

	if card.EaseFactor >= oldEaseFactor {
		t.Errorf("expected to have ease factor decreased. Given value was: %.2f", card.EaseFactor)
	}
}

func TestFourGradesCalculatorShouldDecreaseEaseFactorIfCardWasHardToRecall(t *testing.T) {
	card := kernel.Card{Repetition: 5, Interval: 50, EaseFactor: 2.5}

	oldEaseFactor := card.EaseFactor

	intervalCalculator := FourGradesIntervalCalculator{}
	intervalCalculator.calculate(&card, FourGradesRecalledHard)

	if card.EaseFactor >= oldEaseFactor {
		t.Errorf("expected to have ease factor decreased. Given value was: %.2f", card.EaseFactor)
	}
}

func TestFourGradesCalculatorShouldNotChangeEaseFactorIfCardWasRecalled(t *testing.T) {
	card := kernel.Card{Repetition: 5, Interval: 50, EaseFactor: 2.5}

	oldEaseFactor := card.EaseFactor

	intervalCalculator := FourGradesIntervalCalculator{}
	intervalCalculator.calculate(&card, FourGradesRecalled)

	if card.EaseFactor != oldEaseFactor {
		t.Errorf("expected ease factor not change. Given value was: %.2f", card.EaseFactor)
	}
}

func TestFourGradesCalculatorShouldNotChangeEaseFactorIfCardWasRecalledEasily(t *testing.T) {
	card := kernel.Card{Repetition: 5, Interval: 50, EaseFactor: 2.5}

	oldEaseFactor := card.EaseFactor

	intervalCalculator := FourGradesIntervalCalculator{}
	intervalCalculator.calculate(&card, FourGradesRecalledEasy)

	if card.EaseFactor <= oldEaseFactor {
		t.Errorf("expected to have ease factor increased. Given value was: %.2f", card.EaseFactor)
	}
}

func TestFourGradesCalculatorShouldHandleErrorIfEaseCalculationReturnsError(t *testing.T) {
	card := kernel.Card{Repetition: 5, Interval: 50, EaseFactor: 2.5}

	var zeroGrade GradeValue = 0

	intervalCalculator := FourGradesIntervalCalculator{}
	err := intervalCalculator.calculate(&card, zeroGrade)

	if err == nil {
		t.Errorf("expected to have error handled, but nil was returned.")
	}

	if err.Error() != "An error occured to calculate ease factor" {
		t.Errorf("expected message does not match. Got: %s", err.Error())
	}
}

func TestFourGradesCalculatorShouldNotChangeTheCardIfEaseFactorCalculatorReturnsError(t *testing.T) {
	card := kernel.Card{Repetition: 5, Interval: 50, EaseFactor: 2.5}

	var zeroGrade GradeValue = 0

	intervalCalculator := FourGradesIntervalCalculator{}
	intervalCalculator.calculate(&card, zeroGrade)

	if card.Repetition != 5 {
		t.Errorf("expected to have card repetition unchange. Got: %d", card.Repetition)
	}

	if card.Interval != 50 {
		t.Errorf("expected to have card interval unchange. Got: %d", card.Interval)
	}

	if card.EaseFactor != 2.5 {
		t.Errorf("expected to have card ease factor unchange. Got: %.2f", card.EaseFactor)
	}

}

// handle error
