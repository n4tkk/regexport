package cmd

import (
	"log"
	"os"
)

func check(err error) {
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
