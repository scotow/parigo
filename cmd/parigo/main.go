package parigo

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func parseMenu() {
	// Request the HTML page.
	res, err := http.Get("http://www.crous-lille.fr/restaurant/r-u-pariselle-lille-i/")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	doc.Find("#menu-repas ul.slides").ChildrenFiltered("li").Each(parseDay)
}

func main() {
	parseMenu()
}

func parseDay(_ int, s *goquery.Selection) {
	title := s.Find("h3")
	meals := s.ChildrenFiltered("div.content").First().ChildrenFiltered("div")

	fmt.Println(title.Text())
	meals.Each(parseMeal)
}

func parseMeal(_ int, s *goquery.Selection) {
	title := s.Find("h4")

	fmt.Println(title.Text())
}
