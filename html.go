package main

import (
	"os"
	"path/filepath"
	"text/template"
)

func GenerateHtml(events []Event, htmlDir, now string) error {
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

	if err := GenerateTOC(htmlDir); err != nil {
		return err
	}
	for _, page := range pages {
		if err := GenerateOneHtml(events, page.title, filepath.Join(htmlDir, page.fileName), now, page.filter); err != nil {
			return err
		}
	}
	return nil
}

func GenerateOneHtml(events []Event, title, htmlFileName, now string, filter func(Event) bool) error {
	lines := make([]Line, 0, len(events))
	month := events[0].Date.Month()
	for _, event := range events {
		if filter(event) {
			lines = append(lines, Line{month != event.Date.Month(), event.Date.Format("Mon, Jan 2"), event.Event, event.Boat, event.Pro, event.FirstRace})
			month = event.Date.Month()
		}
	}

	htmlFile, err := os.Create(htmlFileName + ".html")
	if err != nil {
		return err
	}
	defer htmlFile.Close()

	page := Page{title, lines, now}
	return template.Must(template.New("Html Template").Parse(htmlTempl)).Execute(htmlFile, page)
}

type Page struct {
	Title string
	Lines []Line
	Date  string
}

type Line struct {
	NewMonth  bool
	Date      string
	Name      string
	Class     string
	Pro       string
	FirstRace string
}

var htmlTempl = `<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd">
<html xmlns="http://www.w3.org/1999/xhtml">
<head>
<meta content="text/html; charset=utf-8" http-equiv="Content-Type" />
<title>SARYC 2024 Sailing Program, {{.Title}}</title>
<style>
table {
  border-collapse: collapse;
  width: 100%;
}
th {
  text-align: left;
}
</style>
</head>
<body>
<h1 style="text-align:center">SARYC 2024 Sailing Program, {{.Title}}</h1>
<p style="text-align:center">Generated on <strong>{{.Date}}</strong></p></body>
<p>The PRO of the day is responsible for the key to the shed and for laying the course. The PRO should enlist other skippers to help.
<p>Non dual-rated RM may sail in all 10R events except ranking events.
<p><a href="NoR.pdf">Notice of Race</a>
</p>
<table>
<tr><td><b>Date</b></td><td><b>Class</b></td><td><b>Event</b></td><td><b>PRO</b></td><td><b>First Race</b></td></tr>
{{- range $i, $line := .Lines}}
{{if $line.NewMonth}}<tr><td colspan="100%"><hr></hr></td></tr>{{end}}
<tr><td>{{$line.Date}}</td><td>{{$line.Class}}</td><td>{{$line.Name}}</td><td>{{$line.Pro}}</td><td>{{$line.FirstRace}}</td></tr>
{{- end}}
</table>
</html>
`

func GenerateTOC(destDir string) error {
	f, err := os.Create(filepath.Join(destDir, "toc.html"))
	if err != nil {
		return err
	}
	defer f.Close()

	f.WriteString(`<table>
	<tr>
	 <td></td>
	 <td><a href="all.html"><strong>Full program</strong></a></td>
	</tr>
	<tr>
	 <td>By day of the week : </td>
	 <td>
	   <a href="sunday.html">Sundays</a>
	   <a href="tuesday.html">Tuesdays</a>
	   <a href="wednesday.html">Wednesdays</a>
	 </td>
	</tr>
	<tr>
	 <td>By class : </td>
	 <td>
	  <a href="iom.html">IOM</a>
	  <a href="rm.html">Marblehead</a>
	  <a href="10r.html">10R</a>
	  <a href="aclass.html">A Class</a>
	  <a href="laser.html">Laser</a>
	 </td>
	</tr>
	<tr>
	 <td>By class on Sundays: </td>
	 <td>
	  <a href="iom_sunday.html">IOM</a>
	  <a href="rm_sunday.html">Marblehead</a>
	  <a href="10r_sunday.html">10R</a>
	  <a href="aclass_sunday.html">A Class</a>
	 </td>
	</tr>
	</table>
	`)

	return nil
}
