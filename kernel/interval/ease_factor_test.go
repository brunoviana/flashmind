package interval

import (
	"math"
	"strings"
	"testing"
)

func TestEaseFactorMustNotChangeIfTheGradeIsTheMaximumGradeMinusOne(t *testing.T) {
	factors := []struct {
		param1 float64
		param2 int
		param3 int
	}{
		{2.5, 3, 2},
		{2.5, 4, 3},
		{2.5, 6, 5},
	}

	for _, f := range factors {
		factor, _ := calculateEaseFactor(f.param1, f.param2, f.param3)

		if factor != 2.5 {
			t.Errorf("The expected ease factor is 2.5. Given value was: %.2f", factor)
		}
	}
}

func TestEaseFactorCalculatorMustIncreaseGradeEqualsMaxGrade(t *testing.T) {
	parameters := []struct {
		currentEaseFactor float64
		maxGrade          int
		grade             int
	}{
		{2.5, 3, 3},
		{2.5, 4, 4},
		{2.5, 6, 6},
	}

	for _, p := range parameters {
		newEaseFactor, _ := calculateEaseFactor(p.currentEaseFactor, p.maxGrade, p.grade)

		diff := math.Floor((newEaseFactor-p.currentEaseFactor)*100) / 100

		if diff != 0.1 {
			t.Errorf("Expected ease factor grow by 0.1. Got: %2.f", diff)
		}
	}
}

func TestEaseFactorCalculatorMustDecreaseIfTheGradeIsLessThanTheMaximumGradeMinusOne(t *testing.T) {
	parameters := []struct {
		currentEaseFactor float64
		maxGrade          int
		grade             int
	}{
		{2.5, 3, 0},
		{2.5, 3, 1},
		{2.5, 4, 0},
		{2.5, 4, 1},
		{2.5, 4, 2},
		{2.5, 6, 0},
		{2.5, 6, 1},
		{2.5, 6, 2},
		{2.5, 6, 3},
	}

	for _, p := range parameters {
		factor, _ := calculateEaseFactor(p.currentEaseFactor, p.maxGrade, p.grade)

		if factor > p.currentEaseFactor {
			t.Errorf(
				"The expected new ease factor be less than current one (%.2f). Given value was: %.2f",
				p.currentEaseFactor,
				factor,
			)
		}
	}
}

func TestEaseFactorCalculatorCannotReceiveGradeZero(t *testing.T) {
	_, err := calculateEaseFactor(2.5, 3, 0)

	if err == nil {
		t.Errorf("Expected an error but got none")
	} else if !strings.Contains(err.Error(), "grade cannot be zero, a non-zero value is required") {
		t.Errorf("Error message does not contain expected string. Got: %s", err.Error())
	}
}

func TestEaseFactorCalculatorCannotReceiveMaxGradeGreaterThenSix(t *testing.T) {
	_, err := calculateEaseFactor(2.5, 7, 1)

	if err == nil {
		t.Errorf("Expected an error but got none")
	} else if !strings.Contains(err.Error(), "max grade cannot greater than 6") {
		t.Errorf("Error message does not contain expected string. Got: %s", err.Error())
	}
}

func TestEaseFactorCalculatorCannotReceiveMaxGradeLessThanThree(t *testing.T) {
	parameters := []struct {
		currentEaseFactor float64
		maxGrade          int
		grade             int
	}{
		{2.5, 0, 1},
		{2.5, 1, 1},
		{2.5, 2, 1},
	}

	for _, p := range parameters {
		_, err := calculateEaseFactor(p.currentEaseFactor, p.maxGrade, p.grade)

		if err == nil {
			t.Errorf("Expected an error but got none")
		} else if !strings.Contains(err.Error(), "max grade cannot less than 3") {
			t.Errorf("Error message does not contain expected string. Got: %s", err.Error())
		}
	}

}

func TestEaseFactorCalculatorCannotReceiveGradeGreaterThanMaxGrade(t *testing.T) {
	_, err := calculateEaseFactor(2.5, 4, 5)

	if err == nil {
		t.Errorf("Expected an error but got none")
	} else if !strings.Contains(err.Error(), "grade cannot greater than max grade") {
		t.Errorf("Error message does not contain expected string. Got: %s", err.Error())
	}
}
