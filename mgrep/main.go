package main

import (
	"fmt"
	"mgrep/worker"
	"mgrep/worklist"
	"os"
	"path/filepath"
	"sync"

	"github.com/alexflint/go-arg"
)

var args struct {
	SearchTerm string `arg:"positional,required"`
	SearchDir  string `arg:"positional"`
}

// discoverDirs is a recursive function that traverses the directory tree starting from the given path.
// It adds all files it finds to the worklist.
// If it encounters a directory, it calls itself with the new directory path.
// If it encounters an error while reading a directory, it prints an error message and returns.
func discoverDirs(wl *worklist.WorkList, path string) {
	entries, err := os.ReadDir(path)
	if err != nil {
		fmt.Println("Error reading directory: ", err)
		return
	}

	for _, entry := range entries {
		nextPath := filepath.Join(path, entry.Name())
		if entry.IsDir() {
			discoverDirs(wl, nextPath)
		} else {
			wl.Add(worklist.NewJob(nextPath))
		}
	}
}

// The main function in this Go code snippet uses goroutines and channels to concurrently search for a
// specific term in files within a directory and display the results.
func main() {
	arg.MustParse(&args)

	var workersWg sync.WaitGroup

	wl := worklist.New(150)
	results := make(chan worker.Result, 150)
	numWorkers := 15

	workersWg.Add(1)
	go func() {
		defer workersWg.Done()
		discoverDirs(&wl, args.SearchDir)
		wl.Finalize(numWorkers)
	}()

	for i := 0; i < numWorkers; i++ {
		workersWg.Add(1)
		go func() {
			defer workersWg.Done()
			for {
				workEntry := wl.Next()
				if workEntry.Path != "" {
					result := worker.FindInFile(workEntry.Path, args.SearchTerm)
					if result != nil {
						for _, res := range result.Inner {
							results <- res
						}
					}
				} else {
					return
				}
			}
		}()
	}

	blockWorkersWg := make(chan struct{})
	go func() {
		workersWg.Wait()
		close(blockWorkersWg)
	}()

	var displayWg sync.WaitGroup

	displayWg.Add(1)
	go func() {
		for {
			select {
			case res := <-results:
				fmt.Printf("%v[%v]: %v\n", res.Path, res.LineNum, res.Line)
			case <-blockWorkersWg:
				if len(results) == 0 {
					displayWg.Done()
					return
				}
			}
		}
	}()
	displayWg.Wait()
}
