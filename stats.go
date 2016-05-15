package main

import (
	"time"

	"github.com/sjmudd/findphoto/log"
)

type Stats struct {
	last            time.Time
	interval        time.Duration
	Count           int
	Directories     int
	NotRegular      int
	FilenameMatches int
	FullMatches     int
}

// return a Stats struct with reporting interval configured
func NewStats(interval time.Duration) *Stats {
	s := &Stats{
		last:     time.Now(),
		interval: interval,
	}
	return s
}

// report on current status when appropriate
func (s *Stats) Report() {
	if time.Now().Sub(s.last) > time.Second*time.Duration(s.interval) {
		log.Printf("Scanned %d files, %d directories, %d non-regular files. Matches: filename: %d, found: %d\n",
			s.Directories,
			s.NotRegular,
			s.FilenameMatches,
			s.FullMatches)
		s.last = time.Now()
	}
}
