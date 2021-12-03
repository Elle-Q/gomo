package file

import (
	"io/ioutil"
	"os"
)

type file struct {
	path string
	opts Options
}


var (
	DefaultPath = "config.yml"
)

func (f *file) Read() (*ChangeSet, error)  {
	fh, err := os.Open(f.path)
	if err != nil {
		return nil, err
	}
	defer fh.Close()

	b, err := ioutil.ReadAll(fh)
	if err != nil {
		return nil, err
	}

	info,err := fh.Stat()
	if err != nil {
		return nil, err
	}

	cs := &ChangeSet{
		Format: format(f.path, f.opts.Encoder),
		Source: f.String(),
		Timestamp: info.ModTime(),
		Data:b,
	}
	cs.Checksum = cs.Sum()

	return cs, nil
}

func (f *file) String() string {
	return "file"
}

func (f *file) Watch() (Watcher, error)  {
	if _, err := os.Stat(f.path); err != nil {
		return nil,err
	}
	return newWatcher(f)
}

func (f *file) Write(cs *ChangeSet) error {
	return nil
}

func NewSource(opts ...Option) Source {
	options := NewOptions(opts...)
	path := DefaultPath
	f, ok := options.Context.Value(filePathKey{}).(string)
	if ok {
		path = f
	}
	return &file{opts: options, path: path}
}
