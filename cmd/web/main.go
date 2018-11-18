package main

import (
	"github.com/olekukonko/tablewriter"
	"github.com/scotow/parigo"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

func nextMenuDay() (*parigo.Day, error) {
	menu, err := parigo.Current()
	if err != nil {
		return nil, err
	}

	now := time.Now()
	year, month, day := now.Date()
	today := time.Date(year, month, day, 0, 0, 0, 0, now.Location())

	for _, day := range menu.Days {
		if day.Time.Equal(today) || day.Time.After(today) {
			return day, nil
		}
	}

	return nil, nil
}

func handler(w http.ResponseWriter, _ *http.Request) {
	today, err := nextMenuDay()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if today == nil {
		http.Error(w, "no service to display", http.StatusInternalServerError)
		return
	}

	if len(today.Meals) != 1 {
		http.Error(w, "can only display menu if there is only one meal", http.StatusInternalServerError)
		return
	}

	today.Title = strings.Replace(today.Title, "Menu du ", "", -1)
	meal := today.Meals[0]

	table := tablewriter.NewWriter(w)
	table.SetRowLine(true)
	table.SetAutoWrapText(false)
	table.SetAlignment(tablewriter.ALIGN_CENTER)

	table.SetHeader([]string{today.Title})

	for _, part := range meal.Parts {
		plates := make([]string, len(part.Plates))
		for i, plate := range part.Plates {
			plates[i] = strings.Title(plate)
		}
		table.Append([]string{strings.Join(plates, "\n")})
	}

	table.Render()
}

func listeningAddress() string {
	port, set := os.LookupEnv("PORT")
	if !set {
		port = "8080"
	}

	return ":" + port
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(listeningAddress(), nil))
}
