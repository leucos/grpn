package main

import (
	"strings"

	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
)

type executor interface {
	BulkEval(string) error
	Stack() *Stack
	PreviousStack() *Stack
}

func runUI(e executor) error {
	// Initialize application
	app := tview.NewApplication()

	// Create input field
	input := tview.NewInputField()

	// Create deg/rad label
	lmode := tview.NewTextView().SetText("RAD")

	// Create error label
	lerror := tview.NewTextView().SetText("")

	// Stack
	tstack := tview.NewTextView()
	tdebug := tview.NewTextView()

	// Create Grid containing the application's widgets
	appGrid := tview.NewGrid().
		SetColumns(0, 3, 3).
		SetRows(0, 1).
		SetBorders(true).
		// AddItem(tstack, 0, 0, 1, 3, 0, 0, true).
		AddItem(tstack, 0, 0, 1, 1, 0, 0, true).
		AddItem(tdebug, 0, 1, 1, 2, 0, 0, true).
		AddItem(input, 1, 0, 1, 1, 0, 0, true).
		AddItem(lmode, 1, 1, 1, 1, 0, 0, false).
		AddItem(lerror, 1, 2, 1, 1, 0, 0, false)

	var history []string
	histindex := -1

	// Capture user input
	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		// Anything handled here will be executed on the main thread
		switch event.Key() {
		// case tcell.KeyRune:
		// 	val := event.Rune()
		// 	if val == '+' {
		// 		args := input.GetText() + "+"
		// 		tstack.SetText(tstack.GetText(false) + args)
		// 		input.SetText("")
		// 		return nil
		// 	}
		case tcell.KeyUp:
			if len(history) == 0 {
				return nil
			}
			if histindex < 1 || histindex > len(history) {
				histindex = len(history)
			}
			histindex--
			input.SetText(history[histindex])
			return nil
		case tcell.KeyDown:
			if len(history) == 0 {
				return nil
			}
			if histindex >= len(history)-1 {
				histindex = -1
			}
			histindex++
			input.SetText(history[histindex])
			return nil
		case tcell.KeyEnter:
			// Clear error
			lerror.SetText("")
			name := input.GetText()
			if strings.TrimSpace(name) == "" {
				name = "dup"
			}
			history = append(history, name)
			err := e.BulkEval(name)
			if err != nil {
				lerror.SetText("ERR")
			}
			tstack.SetText(e.Stack().String())
			tdebug.SetText(e.PreviousStack().String())
			input.SetText("")
			histindex = len(history)
			return nil
		case tcell.KeyEsc:
			// Exit the application
			app.Stop()
			return nil
		}
		return event
	})

	// Set the grid as the application root and focus the input field
	app.SetRoot(appGrid, true).SetFocus(input)

	// Run the application
	err := app.Run()

	return err
}
