package main

import (
	"flag"
	"os"
	"time"

	"github.com/hutchisont/conways_game_of_life_tui/internal/tui"
)

func main() {
	var (
		rows, cols int
		tickTime   time.Duration
		help       bool
	)
	flag.IntVar(&rows, "rows", 50, "number of rows")
	flag.IntVar(&cols, "cols", 50, "number of columns")
	flag.DurationVar(&tickTime, "tick", 16*time.Millisecond, "time between updates in milliseconds")
	flag.BoolVar(&help, "help", false, "display help")

	flag.Parse()

	if help {
		flag.CommandLine.SetOutput(os.Stdout)
		flag.PrintDefaults()
		return
	}

	t := tui.New(rows, cols, tickTime)
	if err := t.Run(); err != nil {
		panic(err)
	}
}
