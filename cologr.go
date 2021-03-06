package main

import "fmt"
import "regexp"
import "io"
import "os"
import "github.com/hpcloud/tail"
import "github.com/fatih/color"

type CologrLevel struct {
	regexp *regexp.Regexp
	color  func(a ...interface{})
}

func main() {
	args := os.Args[1:]

	if len(args) < 1 {
		fmt.Println("must enter file to log")
		return
	}

	filename := args[0]

	startPos := &tail.SeekInfo{Offset: 0, Whence: io.SeekEnd}
	t, err := tail.TailFile(filename, tail.Config{
		Location: startPos,
		Follow:   true,
		ReOpen:   true,
		Poll:     true,
		Logger:   tail.DiscardingLogger,
	})

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
	fmt.Println(err)
}

func getDefaultRegexMatchers() []CologrLevel {
	trace := regexp.MustCompile(`\b(?i)TRACE\b`)
	debug := regexp.MustCompile(`\b(?i)DEBUG\b`)
	info := regexp.MustCompile(`\b(?i)INFO\b`)
	warn := regexp.MustCompile(`\b(?i)WARN\b`)
	error := regexp.MustCompile(`\b(?i)ERROR\b`)
	fatal := regexp.MustCompile(`\b(?i)FATAL\b`)

	return []CologrLevel{
		CologrLevel{
			regexp: debug,
			color:  color.New(color.FgCyan).PrintlnFunc(),
		},
		CologrLevel{
			regexp: info,
			color:  color.New(color.FgGreen).PrintlnFunc(),
		},
		CologrLevel{
			regexp: warn,
			color:  color.New(color.FgMagenta).PrintlnFunc(),
		},
		CologrLevel{
			regexp: error,
			color:  color.New(color.FgRed).PrintlnFunc(),
		},
		CologrLevel{
			regexp: fatal,
			color:  color.New(color.FgHiRed).PrintlnFunc(),
		},
		CologrLevel{
			regexp: trace,
			color:  color.New(color.FgYellow).PrintlnFunc(),
		},
	}
}
