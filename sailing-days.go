package main

import "time"

func GetEvent(events []Event, date time.Time) (Event, bool) {
	for _, event := range events {
		if date == event.Date {
			return event, true
		}
	}

	return Event{}, false
}

// No sailing at SARYC on these dates.
var no_sailing_days = []Event{
	{time.Date(2024, 3, 31, 0, 0, 0, 0, time.Local), "", "Easter Sunday", "", ""},
	{time.Date(2024, 5, 12, 0, 0, 0, 0, time.Local), "", "Mothers Day", "", ""},
	{time.Date(2024, 9, 1, 0, 0, 0, 0, time.Local), "", "Fathers Day", "", ""},
	{time.Date(2024, 12, 24, 0, 0, 0, 0, time.Local), "", "Christmas Eve", "", ""},
	{time.Date(2024, 12, 25, 0, 0, 0, 0, time.Local), "", "Christmas Day", "", ""},
	// {time.Date(2024, 12, 31, 0, 0, 0, 0, time.Local), "", "New Years Eve", "", ""},
}

// Sailing at other clubs, typically ranking events.
var non_saryc_sailing_days = []Event{
	{time.Date(2024, 1, 21, 0, 0, 0, 0, time.Local), "", "Whitrod (IOM) @ ARCYRC", "", ""},
	{time.Date(2024, 2, 18, 0, 0, 0, 0, time.Local), "", "Toad Cup (RM) @ ARCYRC", "", ""},
	{time.Date(2024, 3, 3, 0, 0, 0, 0, time.Local), "", "State Title (Laser) @ MLMYC", "", ""},
	{time.Date(2024, 3, 4, 0, 0, 0, 0, time.Local), "", "State Title (Laser) @ MLMYC", "", ""},
	{time.Date(2024, 3, 5, 0, 0, 0, 0, time.Local), "", "State Title (Laser) @ MLMYC", "", ""},
	{time.Date(2024, 3, 16, 0, 0, 0, 0, time.Local), "", "State Title (IOM) @ ARCYRC", "", ""},
	{time.Date(2024, 3, 17, 0, 0, 0, 0, time.Local), "", "State Title (IOM) @ ARCYRC", "", ""},
	{time.Date(2024, 3, 20, 0, 0, 0, 0, time.Local), "", "National Title (IOM) @ ARCYRC", "", ""},
	{time.Date(2024, 3, 21, 0, 0, 0, 0, time.Local), "", "National Title (IOM) @ ARCYRC", "", ""},
	{time.Date(2024, 3, 22, 0, 0, 0, 0, time.Local), "", "National Title (IOM) @ ARCYRC", "", ""},
	{time.Date(2024, 3, 23, 0, 0, 0, 0, time.Local), "", "National Title (IOM) @ ARCYRC", "", ""},
	{time.Date(2024, 9, 8, 0, 0, 0, 0, time.Local), "", "Bourneville (RM) @ ARCYRC", "", ""},
	{time.Date(2024, 10, 12, 0, 0, 0, 0, time.Local), "", "State Title (RM) @ ARCYRC", "", ""},
	{time.Date(2024, 10, 13, 0, 0, 0, 0, time.Local), "", "State Title (RM) @ ARCYRC", "", ""},
	{time.Date(2024, 11, 17, 0, 0, 0, 0, time.Local), "", "K McPherson (IOM) @ ARCYRC", "", ""},
}

// Ranking events run by SARYC, on any day of the week.
var ranking_sailing_days = []Event{
	{time.Date(2024, 3, 18, 0, 0, 0, 0, time.Local), "A-Class", "A Class Nationals @ SARYC", "", ""},
	{time.Date(2024, 3, 19, 0, 0, 0, 0, time.Local), "A-Class", "A Class Nationals @ SARYC", "", ""},
	{time.Date(2024, 6, 29, 0, 0, 0, 0, time.Local), "10R", "10R State Title @ SARYC", "", ""},
	{time.Date(2024, 6, 30, 0, 0, 0, 0, time.Local), "10R", "10R State Title @ SARYC", "", ""},
	{time.Date(2024, 9, 15, 0, 0, 0, 0, time.Local), "A-Class", "A Class State Title @ SARYC", "", ""},
}

// Sailing at SARYC on a Sunday that overrides the normalprogram, typically because a ranking event is run nearby.
var special_sailing_days = []Event{
	{time.Date(2024, 2, 18, 0, 0, 0, 0, time.Local), "IOM RM 10R A-Class", "General Sailing All Classes", "TBD", "11am"},  // Toad Cup
	{time.Date(2024, 3, 17, 0, 0, 0, 0, time.Local), "IOM RM 10R A-Class", "General Sailing All Classes", "TBD", "11am"},  // IOM States
	{time.Date(2024, 10, 13, 0, 0, 0, 0, time.Local), "IOM RM 10R A-Class", "General Sailing All Classes", "TBD", "10am"}, // RM States
	{time.Date(2024, 11, 17, 0, 0, 0, 0, time.Local), "IOM RM 10R A-Class", "General Sailing All Classes", "TBD", "10am"}, // KM IOM
}
