/*
Copyright (c) 2016, Simon J Mudd <sjmudd@pobox.com>
All rights reserved.

Redistribution and use in source and binary forms, with or without
modification, are permitted provided that the following conditions are met:

1. Redistributions of source code must retain the above copyright notice, this
   list of conditions and the following disclaimer.
2. Redistributions in binary form must reproduce the above copyright notice,
   this list of conditions and the following disclaimer in the documentation
   and/or other materials provided with the distribution.

THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND
ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED
WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT OWNER OR CONTRIBUTORS BE LIABLE FOR
ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES
(INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES;
LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND
ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
(INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS
SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
*/

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
