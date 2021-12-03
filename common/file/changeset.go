package file

import (
	"crypto/md5"
	"fmt"
	"time"
)


// ChangeSet represents a set of changes from a source
type ChangeSet struct {
	Data      []byte
	Checksum  string
	Format    string
	Source    string
	Timestamp time.Time
}

func (c *ChangeSet) Sum() string  {
	h := md5.New()
	h.Write(c.Data)
	return fmt.Sprintf("%x", h.Sum(nil))

}
