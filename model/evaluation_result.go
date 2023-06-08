package model

type EvaluationResult struct {
	MidStatus   string
	FinalStatus string
	MidScore    float64 `json:"-"`
	FinalScore  float64 `json:"-"`
}
