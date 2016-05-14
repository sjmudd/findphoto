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
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/sjmudd/findphoto/log"
)

var (
	// put the files into a map for easier reference
	locations       = make(map[string]([]string))
	start           time.Time
	last            time.Time
	count           int
	directories     int
	notRegular      int
	filenameMatches int
	fullMatches     int
)

// search scans for files which are matching and records their location
func search(path string, filenames []string) {
	log.Printf("Populating locations...\n")
	for i := range filenames {
		locations[filenames[i]] = nil // not found yet
	}

	// walk the tree at Y looking for files in X
	start = time.Now()
	last = start
	log.Printf("Searching...\n")
	if err := filepath.Walk(path, walkPath); err != nil {
		log.Fatal("Problem walking path %q: %v", path, err)
	}
}

func walkPath(path string, info os.FileInfo, err error) error {
	count++

	if info.Mode().IsDir() {
		directories++
		//	log.Printf("Searching: %s\n", path)
		return nil
	}

	if !info.Mode().IsRegular() {
		notRegular++
		// log.Printf("Ignoring non-file %q\n", path)
		return nil // ignore non files
	}
	if time.Now().Sub(last) > time.Second*10 {
		last = time.Now()
		log.Printf("Scanned %d files, %d directories, %d non-regular files. Matches: filename: %d, full: %d\n",
			count,
			directories,
			notRegular,
			filenameMatches,
			fullMatches)
	}

	components := strings.Split(path, "/")
	filePart := components[len(components)-1]

	existing, found := locations[filePart]
	if !found {
		return nil // filename does not match
	}

	filenameMatches++
	log.MsgVerbose("Filename match: %s: %s\n", filePart, path)

	if !checkCameraModel(path) {
		return nil // not matched the camera model
	}

	// we have a match
	log.MsgInfo("Match: %s: %s\n", filePart, path)

	// finally add the filename details
	fullMatches++
	// update the known locations
	existing = append(existing, path)
	locations[filePart] = existing

	return nil
}
