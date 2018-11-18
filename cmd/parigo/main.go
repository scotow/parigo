package main

import (
	"flag"
	"fmt"
	"github.com/olekukonko/tablewriter"
	"github.com/pkg/errors"
	"github.com/scotow/parigo"
	"os"
	"strings"
	"time"
)

var (
	allFlag = flag.Bool("a", false, "display the whole current week")
)

var (
	errTooManyMeal = errors.New("can only display menu if there is only one meal")
)

func main() {
	flag.Parse()

	if *allFlag {
		menu, err := parigo.Current()
		if err != nil {
			fmt.Println(err)
			return
		}

		for _, day := range menu.Days {
			if err := displayDay(day); err != nil {
				fmt.Println(err)
				return
			}
		}
	} else {
		today, err := nextMenuDay()
		if err != nil {
			fmt.Println(err)
			return
		}

		if today == nil {
			fmt.Println("No service to display.")
			return
		}

		if err := displayDay(today); err != nil {
			fmt.Println(err)
			return
		}
	}
}

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

func displayDay(day *parigo.Day) error {
	if len(day.Meals) != 1 {
		return errTooManyMeal
	}

	day.Title = strings.Replace(day.Title, "Menu du ", "", -1)
	meal := day.Meals[0]

	table := tablewriter.NewWriter(os.Stdout)
	table.SetRowLine(true)
	table.SetAutoWrapText(false)
	table.SetAlignment(tablewriter.ALIGN_CENTER)

	table.SetHeader([]string{day.Title})
	table.SetHeaderColor(tablewriter.Colors{tablewriter.Bold})

	for _, part := range meal.Parts {
		plates := make([]string, len(part.Plates))
		for i, plate := range part.Plates {
			plates[i] = strings.Title(plate)
		}
		table.Append([]string{strings.Join(plates, "\n")})
	}

	table.Render()

	return nil
}
