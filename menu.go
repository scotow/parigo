package parigo

import (
	"errors"
	"github.com/PuerkitoBio/goquery"
	"log"
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

func NewMenu(selection *goquery.Selection) (*Menu, error) {
	
}

type Menu struct {
	Days []*Day
}