package main

import (
	"fmt"
)

type StrengthAssesment interface {
	Access() string
}

type SuperStrongType struct{}

func (s *SuperStrongType) Access() string {
	return "You are so strong"
}

type MiddleStrongType struct{}

func (m *MiddleStrongType) Access() string {
	return "You are normal strong not too much but not weak"
}

type WeakInStrength struct{}

func (w *WeakInStrength) Access() string {
	return "Omg You are so weak!"
}

type Showcase struct {
	StrengthAssesment StrengthAssesment
}

func (a *Showcase) SetAccessment(strategy StrengthAssesment) {
	a.StrengthAssesment = strategy
}

func (a *Showcase) GetAssessment() string {
	return a.StrengthAssesment.Access()
}

func main() {

	context := &Showcase{StrengthAssesment: &SuperStrongType{}}

	fmt.Println("Strength:", context.GetAssessment())

	context.SetAccessment(&MiddleStrongType{})

	fmt.Println("Strength:", context.GetAssessment())

	context.SetAccessment(&WeakInStrength{})

	fmt.Println("Strength:", context.GetAssessment())
}
