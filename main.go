package main

import (
	"flag"
	"os"
	"strconv"
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/hutchisont/conways_game_of_life_tui/internal/board"
	"github.com/rivo/tview"
)

var (
	rows, cols int
	tickTime   time.Duration
	help       bool

	gameBoard [][]string

	app       *tview.Application
	pages     *tview.Pages
	menu      *tview.Modal
	config    *tview.Form
	boardView *tview.TextView
)

func initMenu() *tview.Modal {
	return tview.NewModal().
		SetText("Conway's Game of Life").
		AddButtons([]string{"Play", "Config", "Quit"}).
		SetDoneFunc(func(buttonIndex int, buttonLabel string) {
			switch buttonLabel {
			case "Play":
				pages.SwitchToPage("board")
			case "Config":
				pages.SwitchToPage("config")
			case "Quit":
				app.Stop()
			}
		})
}

func initConfig() *tview.Form {
	form := tview.NewForm()
	form.AddInputField("Rows", "50", 10,
		func(textToCheck string, lastChar rune) bool {
			_, err := strconv.Atoi(textToCheck)
			return err == nil
		},
		func(newRows string) {
			r, err := strconv.Atoi(newRows)
			if err == nil {
				rows = r
			}
		})
	form.AddInputField("Columns", "50", 10,
		func(textToCheck string, lastChar rune) bool {
			_, err := strconv.Atoi(textToCheck)
			return err == nil
		},
		func(newCols string) {
			c, err := strconv.Atoi(newCols)
			if err == nil {
				cols = c
			}
		})
	form.AddInputField("Tick Time (ms)", "16", 10,
		func(textToCheck string, lastChar rune) bool {
			_, err := strconv.Atoi(textToCheck)
			return err == nil
		},
		func(newTick string) {
			t, err := strconv.Atoi(newTick)
			if err == nil {
				tickTime = time.Duration(t) * time.Millisecond
			}
		})
	form.AddButton("Done", func() {
		gameBoard = board.NewRPentomino(rows, cols)
		if boardView != nil {
			boardView.SetSize(rows, cols).SetText(board.AsString(gameBoard))
		}
		pages.SwitchToPage("menu")
	})
	return form
}

func main() {
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

	gameBoard = board.NewRPentomino(rows, cols)
	app = tview.NewApplication()
	pages = tview.NewPages()
	menu = initMenu()
	config = initConfig()

	boardView = tview.NewTextView().
		SetSize(rows, cols).
		SetText(board.AsString(gameBoard))

	pages.AddPage("menu", menu, true, true)
	pages.AddPage("config", config, true, false)
	pages.AddPage("board", boardView, true, false)

	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyESC {
			pages.SwitchToPage("menu")
		}
		return event
	})

	oldTick := tickTime
	ticker := time.NewTicker(tickTime)
	defer ticker.Stop()

	go func() {
		for {
			select {
			case <-ticker.C:
				if name, item := pages.GetFrontPage(); item != nil && name == "board" {
					board.Update(gameBoard)
					app.QueueUpdateDraw(func() {
						boardView.SetText(board.AsString(gameBoard))
					})
				}
				if tickTime > 0 && tickTime != oldTick {
					oldTick = tickTime
					ticker.Reset(tickTime)
				}
			}
		}
	}()

	if err := app.SetRoot(pages, true).Run(); err != nil {
		panic(err)
	}
}
