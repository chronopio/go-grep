package worklist

// Entry represents a job in the worklist with a path.
type Entry struct {
	Path string
}

// WorkList is a structure that holds a channel of jobs.
type WorkList struct {
	jobs chan Entry
}

// Add adds a new entry to the worklist.
// It sends the entry to the jobs channel.
func (wl *WorkList) Add(entry Entry) {
	wl.jobs <- entry
}

// Next retrieves the next entry from the worklist.
// It receives the next entry from the jobs channel.
func (wl *WorkList) Next() Entry {
	return <-wl.jobs
}

// New creates a new WorkList with a buffered channel of the given size.
// It returns a new WorkList instance.
func New(bufferSize int) WorkList {
	return WorkList{
		jobs: make(chan Entry, bufferSize),
	}
}

// NewJob creates a new job with the given path.
// It returns a new Entry instance.
func NewJob(path string) Entry {
	return Entry{
		Path: path,
	}
}

// Finalize adds a number of empty jobs to the worklist.
// It is used to signal the workers that no more jobs will be added.
func (w *WorkList) Finalize(numWorkers int) {
	for i := 0; i < numWorkers; i++ {
		w.Add(Entry{""})
	}
}
