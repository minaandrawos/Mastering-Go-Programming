package main

import (
	"encoding/csv"
	"log"
	"os"
)

func main() {
	records := [][]string{
		{"Jaro", "5", "ALA,IOI"},
		{"Hala", "4", "A8D,B0O"},
		{"Kay", "3", "H8J,D3N"},
	}

	file, err := os.Create("cfilew.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	w := csv.NewWriter(file)
	//w.WriteAll(records)
	w.Comma = ';'

	for _, record := range records {
		if err := w.Write(record); err != nil {
			log.Fatal(err)
		}
	}

	w.Flush()

	err = w.Error()
	if err != nil {

		log.Fatal(err)
	}
}
