package parigo

import (
	"errors"
	"github.com/PuerkitoBio/goquery"
	"net/http"
)

var (
	ErrInvalidAPIResponse = errors.New("invalid response from the API")
)

func Current() (menu *Menu, err error) {
	res, err := http.Get("http://www.crous-lille.fr/restaurant/r-u-pariselle-lille-i/")
	if err != nil {
		return
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		err = ErrInvalidAPIResponse
		return
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return
	}

	menu, err = NewMenu(doc.Find("#menu-repas ul.slides"))
	return
}

func NewMenu(s *goquery.Selection) (menu *Menu, err error) {
	days := make([]*Day, 0, 5)

	s.ChildrenFiltered("li").Each(func(_ int, s *goquery.Selection) {
		day, ed := NewDay(s)
		if ed != nil {
			err = ed
			return
		}
		days = append(days, day)
	})

	if err != nil {
		return nil, err
	}

	return &Menu{days}, nil
}

type Menu struct {
	Days []*Day `json:"days"`
}
