package main

import (
	"fmt"
	"time"
)

func Create() ([]Event, error) {
	events := make([]Event, 0)

	for dayOfYear := 1; dayOfYear <= 366; dayOfYear++ {
		date := time.Date(2024, 1, dayOfYear, 0, 0, 0, 0, time.Local)

		if event, exist := GetEvent(no_sailing_days, date); exist {
			events = append(events, event)
		} else {
			if event, exist := GetEvent(ranking_sailing_days, date); exist {
				events = append(events, event)
			} else {
				if event, exist := GetEvent(special_sailing_days, date); exist {
					events = append(events, event)
				} else {
					switch date.Weekday() {
					case time.Sunday:
						events = append(events, sunday(date))
					case time.Tuesday:
						events = append(events, tuesday(date))
					case time.Wednesday:
						events = append(events, wednesday(date))
					}
				}
			}
		}
		if event, exist := GetEvent(non_saryc_sailing_days, date); exist {
			events = append(events, event)
		}
	}

	return events, nil
}

var sundayClass Class = IOM

func sunday(date time.Time) (event Event) {
	firstRace := firstRace(date)

	if date.Month() == time.February && firstWeek(date) {
		event = Event{date, "IOM RM 10R A-Class", "Rutgers #1", "Tim Arland", firstRace}
	} else if date.Month() == time.April && firstWeek(date) {
		event = Event{date, "IOM RM 10R A-Class", "Rutgers #2", "Daryle Bampton", firstRace}
	} else if date.Month() == time.June && firstWeek(date) {
		event = Event{date, "IOM RM 10R A-Class", "Rutgers #3", "Simon How", firstRace}
	} else if date.Month() == time.August && firstWeek(date) {
		event = Event{date, "IOM RM 10R A-Class", "Rutgers #4", "John Brolese", firstRace}
	} else if date.Month() == time.October && firstWeek(date) {
		event = Event{date, "IOM RM 10R A-Class", "Rutgers #5", "John Brolese", firstRace}
	} else if date.Month() == time.December && firstWeek(date) {
		event = Event{date, "IOM RM 10R A-Class", "Rutgers #6", "Johan Bergkvist", firstRace}
	} else {
		//
		// IOM Club championships
		//
		if sundayClass == IOM && date.Month() == time.January {
			event = Event{date, "IOM", "Club Championship #1", "Ecio Marcel", firstRace}
		} else if sundayClass == IOM && date.Month() == time.April {
			event = Event{date, "IOM", "Club Championship #2", "Tim Paynter", firstRace}
		} else if sundayClass == IOM && date.Month() == time.May {
			event = Event{date, "IOM", "Club Championship #3", "Phil Scapens", firstRace}
		} else if sundayClass == IOM && date.Month() == time.August {
			event = Event{date, "IOM", "Club Championship #4", "Danny James", firstRace}
			//
			// RM Club championships
			//
		} else if sundayClass == RM && date.Month() == time.February {
			event = Event{date, "RM", "Club Championship #1", "Ian Dowsett", firstRace}
		} else if sundayClass == RM && date.Month() == time.July {
			event = Event{date, "RM", "Club Championship #2", "Alan Carli", firstRace}
		} else if sundayClass == RM && date.Month() == time.August {
			event = Event{date, "RM", "Club Championship #3", "Chris Juttner", firstRace}
		} else if sundayClass == RM && date.Month() == time.November {
			event = Event{date, "RM", "Club Championship #4", "Danny James", firstRace}
			//
			// 10R Club championships
			//
		} else if sundayClass == TenRater && date.Month() == time.February {
			event = Event{date, "10R", "Club Championship #1", "Tim Arland", firstRace}
		} else if sundayClass == TenRater && date.Month() == time.April {
			event = Event{date, "10R", "Club Championship #2", "Greg Peake", firstRace}
		} else if sundayClass == TenRater && date.Month() == time.July {
			event = Event{date, "10R", "Club Championship #3", "Steve Arthur", firstRace}
		} else if sundayClass == TenRater && date.Month() == time.November {
			event = Event{date, "10R", "Club Championship #4", "John Brolese", firstRace}
			//
			// A-Class Club championships
			//
		} else if sundayClass == AClass && date.Month() == time.February {
			event = Event{date, "A-Class", "Club Championship #1", "Alan Gold", firstRace}
		} else if sundayClass == AClass && date.Month() == time.March {
			event = Event{date, "A-Class", "Club Championship #2", "Chris Juttner", firstRace}
		} else if sundayClass == AClass && date.Month() == time.June {
			event = Event{date, "A-Class", "Club Championship #3", "Alan Carli", firstRace}
		} else if sundayClass == AClass && date.Month() == time.October {
			event = Event{date, "A-Class", "Club Championship #4", "Tim Arland", firstRace}
		} else {
			event = Event{date, sundayClass.ToString(), sundayClass.ToString() + " Scored Racing", sunday_pros.Next(), firstRace}
		}

		sundayClass.Next()
	}
	return
}

func tuesday(date time.Time) Event {
	if secondWeek((date)) {
		return Event{date, "IOM", fmt.Sprintf("Gold Cup #%d", int(date.Month())), tuesday_pros.Next(), "10am"}
	} else if thirdWeek(date) {
		return Event{date, "IOM", fmt.Sprintf("Trevor Jeffries #%d", int(date.Month())), tuesday_pros.Next(), "10am"}
	} else {
		return Event{date, "IOM", "Scored Racing", tuesday_pros.Next(), "10am"}
	}
}

func wednesday(date time.Time) Event {
	if firstWeek((date)) {
		return Event{date, "Laser", fmt.Sprintf("Laser Club Championship #%d", int(date.Month())), wednesday_pros.Next(), "10am"}
	} else {
		return Event{date, "Laser", "Scored Racing", wednesday_pros.Next(), "10am"}
	}
}

func firstRace(date time.Time) string {
	if date.Weekday() == time.Sunday && date.Month() < 4 {
		return "11am"
	} else {
		return "10am"
	}
}

func firstWeek(date time.Time) bool {
	return date.Day() < 8
}

func secondWeek(date time.Time) bool {
	return date.Day() >= 8 && date.Day() < 15
}

func thirdWeek(date time.Time) bool {
	return date.Day() >= 15 && date.Day() < 22
}

func fourthWeek(date time.Time) bool {
	return date.Day() >= 22 && date.Day() < 29
}
