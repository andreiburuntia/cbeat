package main

import (
	"os"

	"github.com/elastic/beats/libbeat/beat"

	"github.com/andreiburuntia/cbeat/beater"
)

func main() {
	err := beat.Run("cbeat", "", beater.New)
	if err != nil {
		os.Exit(1)
	}
}
