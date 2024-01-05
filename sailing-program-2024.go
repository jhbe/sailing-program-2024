package main

import (
	"flag"
	"log"
	"os"
	"path/filepath"
	"time"
)

func main() {
	var outHtmlDirName string
	var outCsvDirName string
	flag.StringVar(&outHtmlDirName, "o", "out", "name of the output HTML directory")
	flag.StringVar(&outCsvDirName, "c", "csv", "name of the output CSV directory")
	flag.Parse()

	now := time.Now().Format("2006-01-02_15-04-05")

	fullCsvOutDir := filepath.Join(outCsvDirName, now)
	if err := os.MkdirAll(fullCsvOutDir, os.FileMode(os.O_RDWR)); err != nil {
		log.Fatal(err, fullCsvOutDir)
	}

	events, err := Create()
	if err != nil {
		log.Fatalln(err)
	}

	if err := GenerateHtml(events, outHtmlDirName, now); err != nil {
		log.Fatal(err)
	}
	if err := GenerateCsv(events, fullCsvOutDir, now); err != nil {
		log.Fatal(err)
	}
}
