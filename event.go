package main

import (
	"strings"
	"time"
)

type Event struct {
	Date      time.Time
	Boat      string
	Event     string
	Pro       string
	FirstRace string
}

func (e Event) IsSunday() bool {
	return e.Date.Weekday() == time.Sunday
}

func (e Event) IsTuesday() bool {
	return e.Date.Weekday() == time.Tuesday
}

func (e Event) IsWednesday() bool {
	return e.Date.Weekday() == time.Wednesday
}

func (e Event) IsIOM() bool {
	return strings.Contains(strings.ToLower(e.Boat), "iom")
}

func (e Event) IsRM() bool {
	return strings.Contains(strings.ToLower(e.Boat), "rm")
}

func (e Event) Is10R() bool {
	return strings.Contains(strings.ToLower(e.Boat), "10r")
}

func (e Event) IsAClass() bool {
	return strings.Contains(strings.ToLower(e.Boat), "a-class")
}

func (e Event) IsLaser() bool {
	return strings.Contains(strings.ToLower(e.Boat), "laser")
}
