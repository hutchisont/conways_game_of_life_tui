package tui

import (
	"strconv"
	"time"

	"github.com/hutchisont/conways_game_of_life_tui/internal/board"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type TUI struct {
	app             *tview.Application
	pages           *tview.Pages
	menu            *tview.Modal
	config          *tview.Form
	boardConfig     *tview.Table
	boardView       *tview.TextView
	newRows         int
	newCols         int
	curRows         int
	curCols         int
	activeGameBoard [][]string
	customGameBoard [][]string
	tickTime        time.Duration
	useCustomBoard  bool
}

func New(rows, cols int, tickTime time.Duration) *TUI {
	t := TUI{
		app:             tview.NewApplication(),
		pages:           tview.NewPages(),
		curRows:         rows,
		curCols:         cols,
		newRows:         rows,
		newCols:         cols,
		tickTime:        tickTime,
		useCustomBoard:  false,
		activeGameBoard: board.NewRPentomino(rows, cols),
	}

	t.initMenu()
	t.initConfig()
	t.initBoardConfig()

	t.boardView = tview.NewTextView().
		SetSize(t.curRows, t.curCols).
		SetText(board.AsString(t.activeGameBoard))

	t.pages.AddPage("menu", t.menu, true, true)
	t.pages.AddPage("config", t.config, true, false)
	t.pages.AddPage("board", t.boardView, true, false)
	t.pages.AddPage("boardConfig", t.boardConfig, true, false)

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
					board.Update(t.activeGameBoard)
					t.app.QueueUpdateDraw(func() {
						t.boardView.SetText(board.AsString(t.activeGameBoard))
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

func (t *TUI) initBoardConfig() {
	// TODO: probably more correct/efficient to SetContent for the table to
	// a custom implementation rather than using the default like this.
	// It could then be a direct display of a gameBoard I think.
	t.boardConfig = tview.NewTable().
		SetBorders(true)

	for r := 0; r < t.curRows; r++ {
		for c := 0; c < t.curCols; c++ {
			color := tcell.ColorWhite
			t.boardConfig.SetCell(r, c, tview.NewTableCell(" ").SetTextColor(color))
		}
	}

	t.boardConfig.Select(0, 0).
		SetSelectable(true, true).
		SetDoneFunc(func(key tcell.Key) {
			if key == tcell.KeyTab {
				newBoard := board.NewBlank(t.curRows, t.curCols)
				for r := 0; r < t.curRows; r++ {
					for c := 0; c < t.curCols; c++ {
						if t.boardConfig.GetCell(r, c).Text == "O" {
							newBoard[r][c] = "O"
						}
					}
				}
				t.customGameBoard = newBoard
				t.pages.SwitchToPage("config")
			}
		}).
		SetSelectedFunc(func(row, column int) {
			toggleCellText(t.boardConfig.GetCell(row, column))
		})

}

func toggleCellText(cell *tview.TableCell) {
	if cell.Text == " " {
		cell.SetText("O")
	} else {
		cell.SetText(" ")
	}
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
	t.config.AddInputField("Rows", strconv.Itoa(t.curRows), 10,
		func(textToCheck string, lastChar rune) bool {
			_, err := strconv.Atoi(textToCheck)
			return err == nil
		},
		func(newRows string) {
			r, err := strconv.Atoi(newRows)
			if err == nil {
				t.newRows = r
			}
		})
	t.config.AddInputField("Columns", strconv.Itoa(t.curCols), 10,
		func(textToCheck string, lastChar rune) bool {
			_, err := strconv.Atoi(textToCheck)
			return err == nil
		},
		func(newCols string) {
			c, err := strconv.Atoi(newCols)
			if err == nil {
				t.newCols = c
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
	t.config.AddCheckbox("Use Custom Board", false, func(checked bool) {
		t.useCustomBoard = checked
	})
	t.config.AddButton("Custom Board", func() {
		t.pages.SwitchToPage("boardConfig")
	})
	t.config.AddButton("Done", func() {
		changedRowCol := t.updateActiveRowsCols()

		if changedRowCol {
			// have to redo initBoardConfig to account for new rows/cols values
			t.initBoardConfig()
			t.pages.RemovePage("boardConfig")
			t.pages.AddPage("boardConfig", t.boardConfig, true, false)
		}

		if t.useCustomBoard {
			t.activeGameBoard = t.customGameBoard
		} else {
			t.activeGameBoard = board.NewRPentomino(t.curRows, t.curCols)
		}

		t.boardView.SetSize(t.curRows, t.curCols).SetText(board.AsString(t.activeGameBoard))
		t.pages.SwitchToPage("board")
	})
	t.config.AddButton("Back", func() {
		t.pages.SwitchToPage("menu")
	})
}

func (t *TUI) updateActiveRowsCols() bool {
	changed := false
	if t.newRows != t.curRows {
		changed = true
		t.curRows = t.newRows
	}
	if t.newCols != t.curCols {
		changed = true
		t.curCols = t.newCols
	}

	return changed
}
