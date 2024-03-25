package tui

import (
	"strconv"
	"time"

	"github.com/hutchisont/conways_game_of_life_tui/internal/board"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type TUI struct {
	app       *tview.Application
	pages     *tview.Pages
	menu      *tview.Modal
	config    *tview.Form
	boardView *tview.TextView
	rows      int
	cols      int
	gameBoard [][]string
	tickTime  time.Duration
}

func New(rows, cols int, tickTime time.Duration) *TUI {
	t := TUI{
		rows:     rows,
		cols:     cols,
		tickTime: tickTime,
	}

	t.app = tview.NewApplication()
	t.pages = tview.NewPages()
	t.initMenu()
	t.initConfig()

	t.gameBoard = board.NewRPentomino(t.rows, t.cols)

	t.boardView = tview.NewTextView().
		SetSize(t.rows, t.cols).
		SetText(board.AsString(t.gameBoard))

	t.pages.AddPage("menu", t.menu, true, true)
	t.pages.AddPage("config", t.config, true, false)
	t.pages.AddPage("board", t.boardView, true, false)

	t.app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyESC {
			t.pages.SwitchToPage("menu")
		}
		return event
	})

	return &t
}

func (t *TUI) Run() error {
	oldTick := t.tickTime
	ticker := time.NewTicker(t.tickTime)
	defer ticker.Stop()

	go func() {
		for {
			select {
			case <-ticker.C:
				if name, item := t.pages.GetFrontPage(); item != nil && name == "board" {
					board.Update(t.gameBoard)
					t.app.QueueUpdateDraw(func() {
						t.boardView.SetText(board.AsString(t.gameBoard))
					})
				}
				if t.tickTime > 0 && t.tickTime != oldTick {
					oldTick = t.tickTime
					ticker.Reset(t.tickTime)
				}
			}
		}
	}()

	if err := t.app.SetRoot(t.pages, true).Run(); err != nil {
		return err
	}

	return nil
}

func (t *TUI) initMenu() {
	t.menu = tview.NewModal().
		SetText("Conway's Game of Life").
		AddButtons([]string{"Play", "Config", "Quit"}).
		SetDoneFunc(func(buttonIndex int, buttonLabel string) {
			switch buttonLabel {
			case "Play":
				t.pages.SwitchToPage("board")
			case "Config":
				t.pages.SwitchToPage("config")
			case "Quit":
				t.app.Stop()
			}
		})
}

func (t *TUI) initConfig() {
	t.config = tview.NewForm()
	t.config.AddInputField("Rows", strconv.Itoa(t.rows), 10,
		func(textToCheck string, lastChar rune) bool {
			_, err := strconv.Atoi(textToCheck)
			return err == nil
		},
		func(newRows string) {
			r, err := strconv.Atoi(newRows)
			if err == nil {
				t.rows = r
			}
		})
	t.config.AddInputField("Columns", strconv.Itoa(t.cols), 10,
		func(textToCheck string, lastChar rune) bool {
			_, err := strconv.Atoi(textToCheck)
			return err == nil
		},
		func(newCols string) {
			c, err := strconv.Atoi(newCols)
			if err == nil {
				t.cols = c
			}
		})
	t.config.AddInputField("Tick Time (ms)",
		strconv.FormatInt(t.tickTime.Milliseconds(), 10), 10,
		func(textToCheck string, lastChar rune) bool {
			_, err := strconv.Atoi(textToCheck)
			return err == nil
		},
		func(newTick string) {
			tick, err := strconv.Atoi(newTick)
			if err == nil {
				t.tickTime = time.Duration(tick) * time.Millisecond
			}
		})
	t.config.AddButton("Done", func() {
		t.gameBoard = board.NewRPentomino(t.rows, t.cols)
		if t.boardView != nil {
			t.boardView.SetSize(t.rows, t.cols).SetText(board.AsString(t.gameBoard))
		}
		t.pages.SwitchToPage("board")
	})
	t.config.AddButton("Back", func() {
		t.pages.SwitchToPage("menu")
	})
}
