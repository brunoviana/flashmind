package old

import (
	"fmt"
)

type Card struct {
	id         int
	repetition int     // repetiton number, how many times a studend made a successful recall
	interval   int     // how many days the card will be shown again
	easeFactor float64 // easiness factor
}

type Grade int64

const (
	FORGET Grade = iota
	RECALLED_HARD
	RECALLED
	RECALLED_EASY
)

/*
*
  - Values o "howDifficultTheCardWasToRemember" (grade):
  - 0: "Total blackout", complete failure to recall the information
  - 1: Incorrect response, but upon seeing the correct answer it felt familiar
  - 2: Incorrect response, but upon seeing the correct answer it seemed easy to remember
  - 3: Correct response, but required significant effort to recall
  - 4: Correct response, after some hesitation
  - 5: Correct response with perfect recall
*/
func calculateGraduatedInterval(card *Card, howDifficultTheCardWasToRemember int) {
	if howDifficultTheCardWasToRemember < 3 {
		card.repetition = 0
		card.interval = 1
	} else {
		card.repetition++

		if card.repetition == 1 {
			card.interval = 1
		} else if card.repetition == 2 {
			card.interval = 6
		} else {
			card.interval = int(float64(card.interval) * card.easeFactor)
		}
	}

	// easeBefore := card.easeFactor
	card.easeFactor = card.easeFactor + (0.1 - float64(5-howDifficultTheCardWasToRemember)*(0.08+float64(5-howDifficultTheCardWasToRemember)*0.02))

	// fmt.Printf("Ease before: %.2f | Ease after: %.2f\n", easeBefore, card.easeFactor)

	if card.easeFactor < 1.3 {
		card.easeFactor = 1.3
	}
}

func calculateGraduatedIntervalBkp(card *Card, howDifficultTheCardWasToRemember int) {
	if howDifficultTheCardWasToRemember < 1 {
		card.repetition = 0
		card.interval = 1
	} else {
		card.repetition++

		if card.repetition == 1 {
			card.interval = 1
		} else if card.repetition == 2 {
			card.interval = 6
		} else {
			card.interval = int(float64(card.interval) * card.easeFactor)
		}
	}

	// easeBefore := card.easeFactor
	card.easeFactor = card.easeFactor + (0.1 - float64(3-howDifficultTheCardWasToRemember)*(0.08+float64(3-howDifficultTheCardWasToRemember)*0.02))

	// fmt.Printf("Ease before: %.2f | Ease after: %.2f\n", easeBefore, card.easeFactor)

	if card.easeFactor < 1.3 {
		card.easeFactor = 1.3
	}
}

func calculateGraduatedIntervalNew(card *Card, howDifficultTheCardWasToRemember int) {
	if howDifficultTheCardWasToRemember < 1 {
		card.repetition = 0
		card.interval = 1
	} else {
		card.repetition++

		if card.repetition == 1 {
			card.interval = 1
		} else if card.repetition == 3 && card.interval == 1 {
			// it means that one failed a lot and the easy factor is too low to change the interval
			card.interval = 2
		} else {
			card.interval = int(float64(card.interval) * card.easeFactor)
		}
	}

	// easeBefore := card.easeFactor
	card.easeFactor = card.easeFactor + (0.1 - float64(3-howDifficultTheCardWasToRemember)*(0.08+float64(3-howDifficultTheCardWasToRemember)*0.02))

	// fmt.Printf("Ease before: %.2f | Ease after: %.2f\n", easeBefore, card.easeFactor)

	if card.easeFactor < 1.3 {
		card.easeFactor = 1.3
	}
}

func calculateAndPrintGraduatedInterval(card *Card, grade int) {
	calculateGraduatedIntervalNew(card, grade)
	onlyPrint(card, grade)
}

func onlyPrint(card *Card, grade int) {
	fmt.Printf(
		"Grade: %d, Repetition: %d, Interval: %d days, Ease Factor: %.2f\n",
		grade,
		card.repetition,
		card.interval,
		card.easeFactor,
	)
}

func main() {

	card1 := Card{id: 1, repetition: 0, interval: 0, easeFactor: 2.5}
	card2 := Card{id: 1, repetition: 0, interval: 0, easeFactor: 2.5}
	card3 := Card{id: 1, repetition: 0, interval: 0, easeFactor: 2.5}
	card4 := Card{id: 1, repetition: 0, interval: 0, easeFactor: 2.5}
	card5 := Card{id: 1, repetition: 0, interval: 0, easeFactor: 2.5}
	card6 := Card{id: 1, repetition: 0, interval: 0, easeFactor: 2.5}

	fmt.Println("First Case: A lot hard to recall then recall")
	// First Case: Recalled then only forgets
	calculateAndPrintGraduatedInterval(&card1, int(FORGET))
	calculateAndPrintGraduatedInterval(&card1, int(FORGET))
	calculateAndPrintGraduatedInterval(&card1, int(FORGET))
	calculateAndPrintGraduatedInterval(&card1, int(FORGET))
	calculateAndPrintGraduatedInterval(&card1, int(FORGET))

	fmt.Println("\n----------------------\n")

	fmt.Println("Second Case: A lot hard to recall then recall")
	// Second Case: A lot hard to recall then recall
	calculateAndPrintGraduatedInterval(&card2, int(RECALLED_HARD))
	calculateAndPrintGraduatedInterval(&card2, int(RECALLED_HARD))
	calculateAndPrintGraduatedInterval(&card2, int(RECALLED_HARD))
	calculateAndPrintGraduatedInterval(&card2, int(RECALLED_HARD))
	calculateAndPrintGraduatedInterval(&card2, int(RECALLED))

	fmt.Println("\n----------------------\n")

	// Third Case: Always recall
	fmt.Println("Third Case: Always recall")
	calculateAndPrintGraduatedInterval(&card3, int(RECALLED))
	calculateAndPrintGraduatedInterval(&card3, int(RECALLED))
	calculateAndPrintGraduatedInterval(&card3, int(RECALLED))
	calculateAndPrintGraduatedInterval(&card3, int(RECALLED))
	calculateAndPrintGraduatedInterval(&card3, int(RECALLED))

	fmt.Println("\n----------------------\n")

	fmt.Println("Forth Case: A lot hard to recall then many recalls")
	// First Case: Recalled then only forgets
	calculateAndPrintGraduatedInterval(&card4, int(FORGET))
	calculateAndPrintGraduatedInterval(&card4, int(FORGET))
	calculateAndPrintGraduatedInterval(&card4, int(FORGET))
	calculateAndPrintGraduatedInterval(&card4, int(RECALLED))
	calculateAndPrintGraduatedInterval(&card4, int(RECALLED))
	calculateAndPrintGraduatedInterval(&card4, int(RECALLED))
	calculateAndPrintGraduatedInterval(&card4, int(RECALLED))
	calculateAndPrintGraduatedInterval(&card4, int(RECALLED))
	calculateAndPrintGraduatedInterval(&card4, int(RECALLED))

	fmt.Println("\n----------------------\n")

	fmt.Println("Fifth Case: recalls and easy recalls")
	// First Case: Recalled then only forgets
	calculateAndPrintGraduatedInterval(&card5, int(RECALLED))
	calculateAndPrintGraduatedInterval(&card5, int(RECALLED))
	calculateAndPrintGraduatedInterval(&card5, int(RECALLED))
	calculateAndPrintGraduatedInterval(&card5, int(RECALLED_EASY))
	calculateAndPrintGraduatedInterval(&card5, int(RECALLED_EASY))
	calculateAndPrintGraduatedInterval(&card5, int(RECALLED_EASY))

	fmt.Println("\n----------------------\n")

	fmt.Println("Sixth Case: recalls and easy recalls")
	// First Case: Recalled then only forgets
	calculateAndPrintGraduatedInterval(&card6, int(RECALLED))
	calculateAndPrintGraduatedInterval(&card6, int(FORGET))
	calculateAndPrintGraduatedInterval(&card6, int(RECALLED))
	calculateAndPrintGraduatedInterval(&card6, int(RECALLED))
	calculateAndPrintGraduatedInterval(&card6, int(RECALLED))
	calculateAndPrintGraduatedInterval(&card6, int(FORGET))
	calculateAndPrintGraduatedInterval(&card6, int(RECALLED))
	calculateAndPrintGraduatedInterval(&card6, int(RECALLED))
	calculateAndPrintGraduatedInterval(&card6, int(RECALLED))
	calculateAndPrintGraduatedInterval(&card6, int(RECALLED))
	calculateAndPrintGraduatedInterval(&card6, int(RECALLED))
	calculateAndPrintGraduatedInterval(&card6, int(RECALLED))
	calculateAndPrintGraduatedInterval(&card6, int(RECALLED))
	calculateAndPrintGraduatedInterval(&card6, int(RECALLED))
	calculateAndPrintGraduatedInterval(&card6, int(RECALLED))
	calculateAndPrintGraduatedInterval(&card6, int(RECALLED))

	fmt.Println("\n----------------------\n")

	// calculateAndPrintGraduatedInterval(&card, 5)
	// calculateAndPrintGraduatedInterval(&card, 5)
	// calculateAndPrintGraduatedInterval(&card, 5)
	// calculateAndPrintGraduatedInterval(&card, 4)
	// calculateAndPrintGraduatedInterval(&card, 4)
	// calculateAndPrintGraduatedInterval(&card, 4)
	// calculateAndPrintGraduatedInterval(&card, 3)
	// calculateAndPrintGraduatedInterval(&card, 3)
	// calculateAndPrintGraduatedInterval(&card, 3)

	// calculateAndPrintGraduatedInterval(&card, 1)
	// calculateAndPrintGraduatedInterval(&card, 2)
	// calculateAndPrintGraduatedInterval(&card, 2)
	// calculateAndPrintGraduatedInterval(&card, 2)
	// calculateAndPrintGraduatedInterval(&card, 2)
	// calculateAndPrintGraduatedInterval(&card, 0)

	// fmt.Println(" -- Relearning phase --")

	// // relearning
	// card.interval = 1
	// card.repetition = 1
	// onlyPrint(&card, 6)

	// card.interval = 7
	// card.repetition = 1
	// onlyPrint(&card, 6)

	// // minimum interval
	// card.interval = 5
	// card.repetition = 1
	// onlyPrint(&card, 6)

	// fmt.Println(" -- Relearning phase --")

	// calculateAndPrintGraduatedInterval(&card, 2)
	// calculateAndPrintGraduatedInterval(&card, 2)
	// calculateAndPrintGraduatedInterval(&card, 3)
	// calculateAndPrintGraduatedInterval(&card, 3)
	// calculateAndPrintGraduatedInterval(&card, 3)
	// calculateAndPrintGraduatedInterval(&card, 3)
	// calculateAndPrintGraduatedInterval(&card, 3)
	// calculateAndPrintGraduatedInterval(&card, 3)
	// calculateAndPrintGraduatedInterval(&card, 3)
	// calculateAndPrintGraduatedInterval(&card, 2)
	// calculateAndPrintGraduatedInterval(&card, 1)

}
