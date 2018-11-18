package parigo

import "github.com/PuerkitoBio/goquery"

func NewMealPart(title string, s *goquery.Selection) (meal *MealPart, err error) {
	plates := s.Children().Map(func(_ int, s *goquery.Selection) string {
		return s.Text()
	})

	return &MealPart{
		title,
		plates,
	}, nil
}

type MealPart struct {
	Title  string   `json:"title"`
	Plates []string `json:"plates"`
}
