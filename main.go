package main

import (
	"flag"
	"github.com/hutchisont/conways_game_of_life_tui/internal/board"
	"github.com/rivo/tview"
	"os"
	"time"
)

func main() {
	var rows, cols int
	var tickTime time.Duration
	var help bool
	flag.IntVar(&rows, "rows", 50, "number of rows")
	flag.IntVar(&cols, "cols", 50, "number of columns")
	flag.DurationVar(&tickTime, "tick", 100*time.Millisecond, "time between updates in milliseconds")
	flag.BoolVar(&help, "help", false, "display help")

	flag.Parse()

	if help {
		flag.CommandLine.SetOutput(os.Stdout)
		flag.PrintDefaults()
		return
	}

	gameBoard := board.NewDefault(rows, cols)
	app := tview.NewApplication()
	textView := tview.NewTextView().
		SetSize(rows, cols)

	ticker := time.NewTicker(tickTime)
	defer ticker.Stop()

	go func() {
		for {
			select {
			case <-ticker.C:
				board.Update(gameBoard)
				app.QueueUpdateDraw(func() {
					textView.SetText(board.AsString(gameBoard))
				})
			}
		}
	}()

	if err := app.SetRoot(textView, true).Run(); err != nil {
		panic(err)
	}
}
