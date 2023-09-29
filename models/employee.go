package models

import (
	"fmt"
	"io"
	"strconv"
)

type Employee struct {
	ID            string      `json:"id"`
	Name          string      `json:"name"`
	Gender        Gender      `json:"gender"`
	Email         string      `json:"email"`
	LatestLoginAt string      `json:"latestLoginAt"`
	DependentsNum int         `json:"dependentsNum"`
	IsManager     bool        `json:"isManager"`
	Department    *Department `json:"department"`
	Company       *Company    `json:"company"`
}

func (Employee) IsNode()         {}
func (e Employee) GetID() string { return e.ID }

type Gender string

const (
	GenderMale   Gender = "Male"
	GenderFemale Gender = "Female"
)

var AllGender = []Gender{
	GenderMale,
	GenderFemale,
}

func (e Gender) IsValid() bool {
	switch e {
	case GenderMale, GenderFemale:
		return true
	}
	return false
}

func (e Gender) String() string {
	return string(e)
}

func (e *Gender) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = Gender(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Gender", str)
	}
	return nil
}

func (e Gender) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
