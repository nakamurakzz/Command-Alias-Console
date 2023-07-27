package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// Fyne GUI
func main() {
	run()
}

var aliases = []Alias{}
var table *widget.Table

// Fyne GUI
func run() {
	a := app.New()
	w := a.NewWindow("Command Alias Console")
	w.Resize(fyne.NewSize(800, 600))

	loadAliases()

	NewTable()
	sizedTable := container.NewGridWrap(
		fyne.NewSize(800, 500),
		table,
	)

	addButton := addButton()

	w.SetContent(container.NewVBox(
		sizedTable,
		addButton,
	))
	w.ShowAndRun()
}

func addButton() *widget.Button {
	return widget.NewButton("Add", func() {
		addAlias := Alias{
			alias:    "alias",
			command:  "command",
			isActive: true,
		}
		if !canAddAlias(addAlias) {
			return
		}
		aliases = append(aliases, addAlias)
		table.Refresh()
	})
}

func canAddAlias(alias Alias) bool {
	for _, alias := range aliases {
		if alias.alias == "alias" {
			return false
		}
		if alias.command == "command" {
			return false
		}
	}
	return true
}

type Alias struct {
	alias    string
	command  string
	isActive bool
}

func NewTable() {
	table = widget.NewTable(
		func() (int, int) {
			return len(aliases), 2
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("Loading...")
		},
		func(i widget.TableCellID, o fyne.CanvasObject) {
			var text string

			switch i.Col {
			case 0:
				text = aliases[i.Row].alias
			case 1:
				text = aliases[i.Row].command
			}
			o.(*widget.Label).SetText(text)
		},
	)
}

func loadAliases() error {
	home, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("failed to get user home directory: %w", err)
	}
	profile := home + "/.zshrc"
	f, err := os.Open(profile)
	if err != nil {
		return fmt.Errorf("failed to open profile file: %w", err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "alias") {
			alias := strings.Split(line, "=")[0]
			command := strings.Split(line, "=")[1]
			delim := string(command[0])
			aliases = append(aliases, Alias{
				alias:    strings.Split(alias, " ")[1],
				command:  strings.Trim(command, delim),
				isActive: true,
			})
		}
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("failed to scan profile file: %w", err)
	}

	return nil
}
