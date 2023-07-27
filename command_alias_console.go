package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

// Fyne GUI
func main() {
	run()
}

// Fyne GUI
func run() {
	a := app.New()
	w := a.NewWindow("Command Alias Console")
	w.Resize(fyne.NewSize(800, 600))

	tableData, err := createTableData()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	table, err := NewTable(tableData)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	w.SetContent(table)
	w.ShowAndRun()
}

type Alias struct {
	alias    string
	command  string
	isActive bool
}

func NewTable(data [][]string) (fyne.Widget, error) {
	table := widget.NewTable(
		func() (int, int) {
			return len(data), len(data[0])
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("Read profile file///////////////////////////.")
		},
		func(i widget.TableCellID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(data[i.Row][i.Col])
		},
	)
	return table, nil
}

func createTableData() ([][]string, error) {
	aliases, err := getAliaces()
	if err != nil {
		return nil, fmt.Errorf("failed to get aliases: %w", err)
	}
	data := make([][]string, 0, len(aliases))
	for _, alias := range aliases {
		data = append(data, []string{alias.alias, alias.command})
	}
	return data, nil
}

func getAliaces() ([]Alias, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return nil, fmt.Errorf("failed to get user home directory: %w", err)
	}
	profile := home + "/.zshrc"
	f, err := os.Open(profile)
	if err != nil {
		return nil, fmt.Errorf("failed to open profile file: %w", err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var aliases []Alias
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
		return nil, fmt.Errorf("failed to scan profile file: %w", err)
	}

	return aliases, nil
}
