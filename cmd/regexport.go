package cmd

import (
	"encoding/csv"
	"log"
	"os"
	"regexp"
)

func check(err error) {
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
}

func Regex(pattern, input, output string) {
	data, err := os.ReadFile(input)
	check(err)

	file, err := os.Create(output)
	check(err)
	w := csv.NewWriter(file)
	defer w.Flush()

	r := regexp.MustCompile(pattern)
	ms := r.FindAllSubmatch(data, -1)

	for _, m := range ms {
		var row []string
		for _, g := range m {
			row = append(row, string(g))
		}
		err := w.Write(row)
		check(err)
	}
}
