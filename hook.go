package logrusbolt

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/coreos/bbolt"
	"github.com/sirupsen/logrus"
)

const (
	format = "2006-01-02 15:04:05.000000000"
)

type BoltHook struct {
	DBLoc     string
	Bucket    string
	Formatter logrus.Formatter
	Level     logrus.Level
}

// Creates a hook for instance of logrus logger
func NewHook(b BoltHook) (*BoltHook, error) {

	return &BoltHook{
		DBLoc:     b.DBLoc,
		Bucket:    b.Bucket,
		Formatter: b.Formatter,
		Level:     b.Level,
	}, nil
}

// Formats boltdb key
func (b *BoltHook) now() string {
	return time.Now().Format(format) + "." + fmt.Sprint(rand.Uint32())
}

// Calls Fire method when event is fired
func (b *BoltHook) Fire(e *logrus.Entry) error {
	db, err := bolt.Open(b.DBLoc, 0600, nil)

	if err != nil {
		return err
	}
	defer db.Close()

	return db.Update(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists([]byte(b.Bucket))

		if err != nil {
			return err
		}

		bytes, err := b.Formatter.Format(e)

		if err != nil {
			return err
		}

		bucket.Put([]byte(b.now()), bytes)

		return nil
	})
}

// Returns the available logging levels in logrus
func (b *BoltHook) Levels() []logrus.Level {
	levels := []logrus.Level{}
	for _, level := range logrus.AllLevels {
		if level <= b.Level {
			levels = append(levels, level)
		}
	}
	return levels
}
