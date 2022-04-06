package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

// KanbnColumns are the board columns as represented in the kanbn json
type KanbnColumns [][]struct {
	Column      string        `json:"column"`
	Comments    []interface{} `json:"comments"`
	Description string        `json:"description"`
	ID          string        `json:"id"`
	Metadata    struct {
		Assigned string        `json:"assigned"`
		Created  string        `json:"created"`
		Completed  string        `json:"completed"`
		Progress int           `json:"progress"`
		Started  *string       `json:"started,omitempty"`
		Tags     []interface{} `json:"tags"`
		Updated  string        `json:"updated"`
	} `json:"metadata"`
	Name              string        `json:"name"`
	Progress          int           `json:"progress"`
	Relations         []interface{} `json:"relations"`
	RemainingWorkload int           `json:"remainingWorkload"`
	SubTasks          []struct {
		Completed bool   `json:"completed"`
		Text      string `json:"text"`
	} `json:"subTasks"`
	Workload int `json:"workload"`
}

// KanbnBoard represents the JSON output of `kanbn board -j`
type KanbnBoard struct {
	Headings []struct {
		Heading string `json:"heading"`
		Name    string `json:"name"`
	} `json:"headings"`
	Lanes []struct {
		Columns KanbnColumns `json:"columns"`
		Name    string       `json:"name"`
	} `json:"lanes"`
}

func main() {
	err := renderBoard(os.Stdin, os.Stdout)
	if err != nil {
		log.Fatal(err)
	}
}

// renderBoard takes care of reading the kanbn json
// and rendering it as a markdown table
func renderBoard(r io.Reader, w io.Writer) error {
	var kanbn KanbnBoard
	var rows [][]string
	board, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}
	if err = json.Unmarshal(board, &kanbn); err != nil {
		return err
	}

	// Determine how many columns exists and print the header row
	columnCount := len(kanbn.Headings)
	for _, col := range kanbn.Headings {
		fmt.Fprintf(w, "| %s ", col.Name)
	}
	fmt.Fprintln(w, "|")
	for i := 0; i < columnCount; i++ {
		fmt.Fprint(w, "| --- ")
	}
	fmt.Fprintln(w, "|")

	// read the swimlanes
	// The lanes themselves are not printed on the table
	for _, lane := range kanbn.Lanes {
		maxRows := determineMaxRows(lane.Columns)
		// initialize an array of rows
		rows = make([][]string, maxRows)
		for i := range rows {
			rows[i] = make([]string, columnCount)
		}
		// convert columns of rows into rows of columns
		for curCol, column := range lane.Columns {
			for colRow, task := range column {
                                complete := ""
                                if task.Progress == 1 || task.Metadata.Completed != "" {
                                    complete="~~"
                                }
				rows[colRow][curCol] = fmt.Sprintf("%s%s%s", complete, task.Name, complete)
			}
		}
		// print the rows
		for _, row := range rows {
			for _, col := range row {
				fmt.Fprintf(w, "| %s ", col)
			}
			fmt.Fprintln(w, "|")
		}
	}
	return nil
}

func determineMaxRows(cols KanbnColumns) int {
	var maxRows int
	for _, column := range cols {
		if len(column) > maxRows {
			maxRows = len(column)
		}
	}
	return maxRows
}
