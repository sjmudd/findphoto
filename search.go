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
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/sjmudd/findphoto/log"
)

var (
	// put the files into a map for easier reference
	locations = make(map[string]([]string))
)

// search scans for files which are matching and records their location
func search(path string, filenames []string) {
	// log.Printf("Populating locations...\n")
	for i := range filenames {
		locations[filenames[i]] = nil // not found yet
	}

	// walk the tree at Y looking for files in X
	log.Printf("Searching...\n")
	if err := filepath.Walk(path, walkPath); err != nil {
		log.Fatal("Problem walking path %q: %v", path, err)
	}
}

// create a symlink from symlinkDir/filePart -> path
// - if there's a symlink there already then just return
func symlinkMatch(filePart, path string) {
	link := symlinkDir + "/" + filePart

	info, err := os.Stat(link)
	if err == nil {
		// Stat() succeeds
		if (info.Mode() & os.ModeSymlink) != 0 {
			log.MsgInfo("Symlink %s already exists: ignoring file: %s\n", link, path)
			return
		}
	} else {
		// ignore errors as probably the file is missing.
		// perhaps I should only ignore the error "not found" or whatever but worry about that later.
		log.MsgDebug("Error with Stat(%q): %v", link, err)
	}

	err = os.Symlink(path, link)
	if err != nil {
		log.MsgError("Failed to create symlink: %s -> %s: %v\n", link, path, err)
		os.Exit(1)
	}
	log.MsgInfo("Created symlink: %s -> %s\n", link, path)
}

// walkPath scans through the directory tree given looking for files
// which match the name in locations.  Every progressInterval seconds
// an update is logged showing the current status as scanning a large
// filesystem such as one containing backups can take a long time.
// We check the exif data of the file to see if it matches the
// camera-model, and if we have a match we either show the match
// or make a symlink to the symlink-dir if defined.
func walkPath(path string, info os.FileInfo, err error) error {
	log.MsgDebug("walkPath(%q,...)\n", path)
	counters.Count++

	if info.Mode().IsDir() {
		counters.Directories++
		//	log.Printf("Searching: %s\n", path)
		return nil
	}
	log.MsgDebug("%q is not a directory\n", path)

	if !info.Mode().IsRegular() {
		counters.NotRegular++
		// log.Printf("Ignoring non-file %q\n", path)
		return nil // ignore non files
	}
	log.MsgDebug("%q is a regular file\n", path)
	log.MsgDebug("showCameraModel: %v\n", showCameraModel)

	if showCameraModel {
		log.MsgDebug("looking for camera model\n")
		cameraModel, err := getCameraModel(path)
		if err == nil {
			log.MsgDebug("Found camera model: %v for %v\n", cameraModel, path)
			fmt.Printf("%s: %s\n", cameraModel, path)
		} else {
			log.MsgDebug("getCameraModel(%q) returns error: %v\n", path, err)
		}
		return nil
	}
	log.MsgDebug("not showing the camera model\n", path)

	log.MsgInfo("done some camera model stuff already maybe\n")

	// report counters if needed
	counters.Report()

	components := strings.Split(path, "/")
	filePart := components[len(components)-1]

	existing, found := locations[filePart]
	if !found {
		return nil // filename does not match
	}

	counters.FilenameMatches++
	// log.MsgVerbose("Filename match: %s: %s\n", filePart, path)

	if !checkCameraModel(path) {
		return nil // not matched the camera model
	}

	counters.FullMatches++ // finally add the filename details
	// update the known locations
	existing = append(existing, path)
	locations[filePart] = existing

	if symlinkDir != "" {
		symlinkMatch(filePart, path)
	} else {
		// Just show the match
		log.MsgInfo("Match: %s: %s\n", filePart, path)
	}

	return nil
}
