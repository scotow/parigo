package parigo

import (
	"errors"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"time"
)

var (
	months = [...]string{
		"janvier",
		"février",
		"mars",
		"avril",
		"mai",
		"juin",
		"juillet",
		"aout",
		"septembre",
		"octobre",
		"novembre",
		"décembre",
	}
	locale *time.Location
)

var (
	ErrCannotParseDay = errors.New("error while parsing menu day string")
)

func init() {
	locale, _ = time.LoadLocation("Europe/Paris")
}

func NewDay(s *goquery.Selection) (day *Day, err error) {
	meals := make([]*Meal, 0, 5)

	s.ChildrenFiltered("div.content").First().Children().Each(func(_ int, s *goquery.Selection) {
		meal, ed := NewMeal(s)
		if ed != nil {
			err = ed
			return
		}

		if meal == nil {
			return
		}

		meals = append(meals, meal)
	})

	if err != nil {
		return
	}

	title := s.ChildrenFiltered("h3").Text()
	time, err := parseDate(title)
	if err != nil {
		return
	}

	day = &Day{title, time, meals}
	return
}

func parseDate(s string) (t time.Time, err error) {
	var day, year int
	var frMonth string

	_, err = fmt.Sscanf(s, "Menu du %s %d %s %d", new(string), &day, &frMonth, &year)
	if err != nil {
		return
	}

	var month int
	for i, m := range months {
		if frMonth == m {
			month = i + 1
			break
		}
	}

	if month == 0 {
		t, err = time.Time{}, ErrCannotParseDay
		return
	}

	t = time.Date(year, time.Month(month), day, 0, 0, 0, 0, locale)
	return
}

type Day struct {
	Title string    `json:"title"`
	Time  time.Time `json:"time"`
	Meals []*Meal   `json:"meals"`
}
