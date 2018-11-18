package main

import (
	"encoding/json"
	"fmt"
	"github.com/scotow/parigo"
	"log"
)

func main() {
	menu, err := parigo.Current()
	if err != nil {
		log.Panic(err)
	}

	jsonData, err := json.MarshalIndent(menu, "", "\t")
	if err != nil {
		log.Panic(err)
	}

	fmt.Println(string(jsonData))
}
