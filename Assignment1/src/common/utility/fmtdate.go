package fmtdate

import (
	"strings"
	"time"
)

func replace(in string) (out string) {
	out = in
	for _, ph := range Placeholder {
		out = strings.Replace(out, ph.find, ph.subst, -1)
	}
	return
}

// Format formats a date based on Microsoft Excel (TM) conventions
func Format(format string, date time.Time) string {
	return date.Format(replace(format))
}

// Parse parses a value to a date based on Microsoft Excel (TM) formats
func Parse(format string, value string) (time.Time, error) {
	return time.Parse(replace(format), value)
}

type p struct{ find, subst string }

var Placeholder = []p{
	p{"hh", "15"},
	p{"h", "03"},
	p{"mm", "04"},
	p{"ss", "05"},
	p{"MMMM", "January"},
	p{"MMM", "Jan"},
	p{"MM", "01"},
	p{"M", "1"},
	p{"pm", "PM"},
	p{"ZZZZ", "-0700"},
	p{"ZZZ", "MST"},
	p{"ZZ", "Z07:00"},
	p{"YYYY", "2006"},
	p{"YY", "06"},
	p{"DDDD", "Monday"},
	p{"DDD", "Mon"},
	p{"DD", "02"},
	p{"D", "2"},
}

var (
	DefaultTimeFormat     = "hh:mm:ss"
	DefaultDateFormat     = "YYYY-MM-DD"
	DefaultDateTimeFormat = "YYYY-MM-DD hh:mm:ss"
)

// FormatDate formats the given date to the DefaultDateFormat
func FormatDate(date time.Time) string {
	return Format(DefaultDateFormat, date)
}

// FormatTime formats the given date to the DefaultTimeFormat
func FormatTime(date time.Time) string {
	return Format(DefaultTimeFormat, date)
}

// FormatTime formats the given date to the DefaultDateTimeFormat
func FormatDateTime(date time.Time) string {
	return Format(DefaultDateTimeFormat, date)
}

// Parse parses a date in DefaultDateFormat to a date
func ParseDate(value string) (time.Time, error) {
	return Parse(DefaultDateFormat, value)
}

// Parse parses a date in DefaultTimeFormat to a date
func ParseTime(value string) (time.Time, error) {
	return Parse(DefaultTimeFormat, value)
}

// Parse parses a date in DefaultDateTimeFormat to a date
func ParseDateTime(value string) (time.Time, error) {
	return Parse(DefaultDateTimeFormat, value)
}

func GetMonthString(month time.Month) string {

	if month == time.January {
		return "January"
	} else if month == time.February {
		return "February"
	} else if month == time.March {
		return "March"
	} else if month == time.April {
		return "April"
	} else if month == time.May {
		return "May"
	} else if month == time.June {
		return "June"
	} else if month == time.July {
		return "July"
	} else if month == time.August {
		return "August"
	} else if month == time.September {
		return "September"
	} else if month == time.October {
		return "October"
	} else if month == time.November {
		return "November"
	} else {
		return "December"
	}

}
