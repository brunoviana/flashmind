package interval

import "errors"

func calculateEaseFactor(currentEaseFactor float64, maxGrade int, grade int) (float64, error) {
	if grade == 0 {
		return 0, errors.New("grade cannot be zero, a non-zero value is required")
	}

	if maxGrade > 6 {
		return 0, errors.New("max grade cannot greater than 6")
	}

	if maxGrade < 3 {
		return 0, errors.New("max grade cannot less than 3")
	}

	if grade > maxGrade {
		return 0, errors.New("grade cannot greater than max grade")
	}

	maxGrade--
	grade--

	return currentEaseFactor + (0.1 - float64(maxGrade-grade)*(0.08+float64(maxGrade-grade)*0.02)), nil
}
