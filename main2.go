package main

import (
	"fmt"
)

type StrengthAssessment interface {
	Evaluate() string
}

type SuperStrongType struct{}

func (s *SuperStrongType) Evaluate() string {
	return "You've got superhero strength!"
}

type MiddleStrongType struct{}

func (m *MiddleStrongType) Evaluate() string {
	return "You're pretty strong, like an everyday hero!"
}

type WeakInStrength struct{}

func (w *WeakInStrength) Evaluate() string {
	return "Uh-oh, you might need some more spinach!"
}

type StrengthShowcase struct {
	StrengthAssessment StrengthAssessment
}

func (s *StrengthShowcase) SetAssessment(strategy StrengthAssessment) {
	s.StrengthAssessment = strategy
}

func (s *StrengthShowcase) GetEvaluation() string {
	return s.StrengthAssessment.Evaluate()
}

func main() {

	showcase := &StrengthShowcase{StrengthAssessment: &SuperStrongType{}}

	fmt.Println("Strength Evaluation:", showcase.GetEvaluation())

	showcase.SetAssessment(&MiddleStrongType{})

	fmt.Println("Strength Evaluation:", showcase.GetEvaluation())

	showcase.SetAssessment(&WeakInStrength{})

	fmt.Println("Strength Evaluation:", showcase.GetEvaluation())
}
