package parigo

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/pkg/errors"
)

var (
	ErrInvalidNumberOfMealParts = errors.New("invalid number of meal parts")
)

func NewMeal(s *goquery.Selection) (meal *Meal, err error) {
	parts := make([]*MealPart, 0, 3)

	children := s.ChildrenFiltered("div.content-repas").First().Children().First().Children()
	childrenLength := children.Length()

	if childrenLength == 1 {
		return
	}

	if childrenLength%2 != 0 {
		err = ErrInvalidNumberOfMealParts
		return
	}

	for i := 0; i < childrenLength; i += 2 {
		children.Eq(i)
		part, ep := NewMealPart(children.Eq(i).Text(), children.Eq(i+1))
		if ep != nil {
			err = ep
			return
		}
		parts = append(parts, part)
	}

	if err != nil {
		return
	}

	title := s.ChildrenFiltered("h4").Text()
	meal = &Meal{title, parts}
	return
}

type Meal struct {
	Title string      `json:"title"`
	Parts []*MealPart `json:"parts"`
}
