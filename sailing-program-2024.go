package main

import (
	"flag"
	"log"
)

func main() {
	var outDirName string
	flag.StringVar(&outDirName, "o", "out", "name of the output directory")
	flag.Parse()

	events, err := Create()
	if err != nil {
		log.Fatalln(err)
	}

	if err := GenerateHtml(events, outDirName); err != nil {
		log.Fatal(err)
	}
}
