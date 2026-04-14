package main

import (
	"flag"
	"fmt"
	"strconv"
	"time"
)

var format = flag.String("f", "2006-01-02 15:04:05", "Output format")
var relative = flag.Bool("r", false, "Relative time")
var unix = flag.Bool("u", false, "Unix timestamp")

func main() {
	flag.Parse()
	args := flag.Args()

	if len(args) == 0 {
		printCurrent()
		return
	}

	input := args[0]
	if *unix || isNumeric(input) {
		ts, _ := strconv.ParseInt(input, 10, 64)
		if ts > 0 {
			t := time.Unix(ts, 0)
			fmt.Println(t.Format(*format))
			return
		}
	}

	t, err := time.Parse("2006-01-02 15:04:05", input)
	if err != nil {
		t, err = time.Parse("2006-01-02", input)
	}
	if err != nil {
		fmt.Printf("Cannot parse: %s\n", input)
		return
	}

	if *relative {
		fmt.Println(formatRelative(t))
	} else {
		fmt.Println(t.Format(*format))
	}
}

func printCurrent() {
	now := time.Now()
	if *relative {
		fmt.Println(formatRelative(now))
	} else if *unix {
		fmt.Println(now.Unix())
	} else {
		fmt.Println(now.Format(*format))
	}
}

func isNumeric(s string) bool {
	for _, c := range s {
		if c < '0' || c > '9' {
			return false
		}
	}
	return len(s) > 0
}

func formatRelative(t time.Time) string {
	d := time.Since(t)
	if d < 0 {
		return "in " + formatDuration(-d)
	}
	return formatDuration(d) + " ago"
}

func formatDuration(d time.Duration) string {
	if d < time.Minute {
		return fmt.Sprintf("%ds", int(d.Seconds()))
	}
	if d < time.Hour {
		return fmt.Sprintf("%dm", int(d.Minutes()))
	}
	if d < 24*time.Hour {
		return fmt.Sprintf("%dh", int(d.Hours()))
	}
	return fmt.Sprintf("%dd", int(d.Hours()/24))
}
