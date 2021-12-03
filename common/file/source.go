package file

import (
	"errors"
)

var (
	// ErrWatcherStopped is returned when source watcher has been stopped
	ErrWatcherStopped = errors.New("watcher stopped")
)
// Source is the source from which config is loaded
type Source interface {
	Read() (*ChangeSet, error)
	Write(*ChangeSet) error
	Watch() (Watcher, error)
	String() string
}


// Watcher watches a source for changes
type Watcher interface {
	Next() (*ChangeSet, error)
	Stop() error
}