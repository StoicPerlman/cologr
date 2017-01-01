package main

import "fmt"
import "regexp"
import "github.com/hpcloud/tail"
import "github.com/fatih/color"

type CologrLevel struct {
	regexp *regexp.Regexp
	color func(a ...interface{})
}

func main() {
    t, err := tail.TailFile("./test.log", tail.Config{Follow: true, Poll: true})

    cologrLevels := getDefaultRegexMatchers()
    printer := color.New(color.FgWhite).PrintlnFunc()

	for line := range t.Lines {

		for _, cologrLevel := range cologrLevels {
			if cologrLevel.regexp.MatchString(line.Text) {
				printer = cologrLevel.color
				break
			}
		}

		printer(line.Text)
	}
	fmt.Println(err);
}

func getDefaultRegexMatchers() []CologrLevel {
	trace := regexp.MustCompile("(?i)TRACE")
	debug := regexp.MustCompile("(?i)DEBUG")
	info := regexp.MustCompile("(?i)INFO")
	warn := regexp.MustCompile("(?i)WARN")
	error := regexp.MustCompile("(?i)ERROR")
	fatal := regexp.MustCompile("(?i)FATAL")

	return []CologrLevel {
		CologrLevel {
			regexp: debug,
			color: color.New(color.FgCyan).PrintlnFunc(),
		},
		CologrLevel {
			regexp: info,
			color: color.New(color.FgGreen).PrintlnFunc(),
		},
		CologrLevel {
			regexp: warn,
			color: color.New(color.FgHiRed).PrintlnFunc(),
		},
		CologrLevel {
			regexp: error,
			color: color.New(color.FgRed).PrintlnFunc(),
		},
		CologrLevel {
			regexp: fatal,
			color: color.New(color.FgMagenta).PrintlnFunc(),
		},
		CologrLevel {
			regexp: trace,
			color: color.New(color.FgYellow).PrintlnFunc(),
		},
	}
}