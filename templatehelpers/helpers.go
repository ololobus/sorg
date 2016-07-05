package templatehelpers

import (
	"fmt"
	"html/template"
	"time"
)

// FuncMap is a set of helper functions to make available in templates for the
// project.
var FuncMap template.FuncMap

func init() {
	FuncMap = template.FuncMap{
		"FormatTime": formatTime,
		"InKM":       inKM,
		"MonthName":  monthName,
		"Pace":       pace,
		"Round":      round,
	}
}

func formatTime(t *time.Time) string {
	return t.Format("January 2, 2006")
}

func inKM(m float64) float64 {
	return m / 1000.0
}

func monthName(t *time.Time) string {
	return t.Format("January")
}

// pace calculates the pace of a run in time per kilometer. This comes out as a
// "clock" time like 4:52 which translates to "4 minutes and 52 seconds" per
// kilometer.
func pace(distance float64, duration time.Duration) string {
	speed := float64(duration.Seconds()) / inKM(distance)
	min := int64(speed / 60.0)
	sec := int64(speed) % 60
	return fmt.Sprintf("%v:%02d", min, sec)
}

func round(f float64) string {
	return fmt.Sprintf("%.1f", f)
}
