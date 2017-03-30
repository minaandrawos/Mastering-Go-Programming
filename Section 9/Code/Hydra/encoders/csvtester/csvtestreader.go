package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	file, err := os.Open("cfile.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	r := csv.NewReader(file)
	r.Comment = '#'
	r.Comma = ';'
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			if pe, ok := err.(*csv.ParseError); ok {
				fmt.Println("bad column:", pe.Column)
				fmt.Println("bad line:", pe.Line)
				fmt.Println("Error reported", pe.Err)
				if pe.Err == csv.ErrFieldCount {
					continue
				}
			}
			log.Fatal(err)
		}
		fmt.Println("CSV Row:", record)
		/*
			i, err := strconv.Atoi(record[1])
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(i * 4)
		*/
	}

}
