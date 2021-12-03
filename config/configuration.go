package config

import "gomo/common/file"

type Configuration interface {
	// Values provide the reader.Values interface
	reader.Values
	// Init the config
	Init(opts ...Option) error
	// Options in the config
	Options() Options
	// Close Stop the config loader/watcher
	Close() error
	// Load config sources
	Load(source ...Source) error
	// Sync Force a source changeset sync
	Sync() error
	// Watch a value for changes
	Watch(path ...string) (Watcher, error)
}