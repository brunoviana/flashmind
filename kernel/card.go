package kernel

type Card struct {
	Repetition int     // repetiton number, how many times a studend made a successful recall
	Interval   int     // how many days the card will be shown again
	EaseFactor float64 // easiness factor
}
