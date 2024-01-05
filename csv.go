package main

import (
	"os"
	"path/filepath"
	"text/template"
)

func GenerateCsv(events []Event, csvDir, now string) error {
	pages := []struct {
		title    string
		fileName string
		filter   func(Event) bool
	}{
		{"All", "all", func(e Event) bool { return true }},

		{"Sundays", "sunday", func(e Event) bool { return e.IsSunday() }},
		{"Tuesdays", "tuesday", func(e Event) bool { return e.IsTuesday() }},
		{"Wednesdays", "wednesday", func(e Event) bool { return e.IsWednesday() }},

		{"IOM", "iom", func(e Event) bool { return e.IsIOM() }},
		{"RM", "rm", func(e Event) bool { return e.IsRM() }},
		{"10R", "10r", func(e Event) bool { return e.Is10R() }},
		{"A Class", "aclass", func(e Event) bool { return e.IsAClass() }},
		{"Laser", "laser", func(e Event) bool { return e.IsLaser() }},

		{"IOM on Sundays", "iom_sunday", func(e Event) bool { return e.IsIOM() && e.IsSunday() }},
		{"RM on Sundays", "rm_sunday", func(e Event) bool { return e.IsRM() && e.IsSunday() }},
		{"10R on Sundays", "10r_sunday", func(e Event) bool { return e.Is10R() && e.IsSunday() }},
		{"A Class on Sundays", "aclass_sunday", func(e Event) bool { return e.IsAClass() && e.IsSunday() }},
	}

	for _, page := range pages {
		if err := GenerateOneCsv(events, page.title, filepath.Join(csvDir, page.fileName), now, page.filter); err != nil {
			return err
		}
	}
	return nil
}

func GenerateOneCsv(events []Event, title, csvFileName, now string, filter func(Event) bool) error {
	lines := make([]CsvLine, 0, len(events))
	for _, event := range events {
		if filter(event) {
			lines = append(lines, CsvLine{event.Date.Format("Mon,Jan 2"), event.Event, event.Boat, event.Pro, event.FirstRace})
		}
	}

	csvFile, err := os.Create(csvFileName + ".csv")
	if err != nil {
		return err
	}
	defer csvFile.Close()

	page := CsvPage{lines, now}
	return template.Must(template.New("Csv Template").Parse(csvTempl)).Execute(csvFile, page)
}

type CsvPage struct {
	Lines []CsvLine
	Date  string
}

type CsvLine struct {
	Date      string
	Name      string
	Class     string
	Pro       string
	FirstRace string
}

var csvTempl = `{{range $i, $line := .Lines}}{{$line.Date}},{{$line.Class}},{{$line.Name}},{{$line.Pro}},{{$line.FirstRace}}
{{end}}`
