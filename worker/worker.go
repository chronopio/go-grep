package worker

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Result struct {
	Line    string
	LineNum int
	Path    string
}

type Results struct {
	Inner []Result
}

// NewResult creates a new Result instance.
// It takes a line where the search query was found, the line number, and the path of the file.
// It returns a new Result instance.
func NewResult(line string, lineNum int, path string) Result {
	return Result{
		Line:    line,
		LineNum: lineNum,
		Path:    path,
	}
}

// FindInFile opens the file at the given path and searches for the query string in the file.
// It returns a pointer to a Results struct containing all the lines where the query was found.
// If the file cannot be opened, it prints an error message and returns nil.
// If the query is not found in the file, it returns nil.
func FindInFile(path, query string) *Results {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Error opening file: ", err)
		return nil
	}

	results := Results{make([]Result, 0)}

	scanner := bufio.NewScanner(file)
	lineNum := 1
	for scanner.Scan() {
		if strings.Contains(scanner.Text(), query) {
			results.Inner = append(results.Inner, NewResult(scanner.Text(), lineNum, path))
		}
		lineNum++
	}

	if len(results.Inner) == 0 {
		return nil
	} else {
		return &results
	}
}
