package internal

import (
	"fmt"

	"github.com/fatih/color"
)

type DiffType int

const (
	Added DiffType = iota
	Removed
	Changed
)

type DiffResult struct {
	Key      string
	OldValue string
	NewValue string
	Type     DiffType
}

func DiffEnvs(from, to map[string]string) []DiffResult {
	var results []DiffResult

	for k, v := range from {
		if newVal, ok := to[k]; ok {
			if newVal != v {
				results = append(results, DiffResult{Key: k, OldValue: v, NewValue: newVal, Type: Changed})
			}
		} else {
			results = append(results, DiffResult{Key: k, OldValue: v, Type: Removed})
		}
	}

	for k, v := range to {
		if _, ok := from[k]; !ok {
			results = append(results, DiffResult{Key: k, NewValue: v, Type: Added})
		}
	}

	return results
}

func PrintDiff(results []DiffResult) {
	green := color.New(color.FgGreen).SprintFunc()
	red := color.New(color.FgRed).SprintFunc()
	yellow := color.New(color.FgYellow).SprintFunc()

	for _, r := range results {
		switch r.Type {
		case Added:
			fmt.Printf("%s %s added (%s)\n", green("+"), r.Key, r.NewValue)
		case Removed:
			fmt.Printf("%s %s removed (was %s)\n", red("-"), r.Key, r.OldValue)
		case Changed:
			fmt.Printf("%s %s changed from %s to %s\n", yellow("~"), r.Key, r.OldValue, r.NewValue)
		}
	}
}
