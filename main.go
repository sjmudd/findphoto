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

// Package main searches for images given a path matching the specific filename
package main

import (
	"bufio"
	"flag"
	"os"

	"github.com/sjmudd/findphoto/log"
)

var (
	cameraModel string // e.g. Camera Model Name : Canon PowerShot S100
	searchFile  string // file containing photo names
)

// given a filename to collect names from return a list of names
func getFiles(filename string) []string {
	var filenames []string

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		entry := scanner.Text()
		log.MsgVerbose("Entry: %s\n", entry)
		filenames = append(filenames, entry)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return filenames
}

// usage returns a usage message and exits with the requested exit code
func usage(exitCode int) {
	log.Printf("Usage: %s <options> <directory_to_search>\n\n", os.Args[0])
	flag.PrintDefaults()

	os.Exit(exitCode)
}

func main() {
	// get options
	flag.BoolVar(&log.Verbose, "verbose", false, "Enable verbose logging")
	flag.StringVar(&searchFile, "search-file", "", "Required: File to use containing a line of the base filesnames to search for")
	flag.StringVar(&cameraModel, "camera-model", "", "camera model (in exif data e.g. 'Canon PowerShot S100'")
	flag.Parse()

	if cameraModel != "" {
		log.Printf("camera-model: %s\n", cameraModel)
	}
	if searchFile == "" {
		log.Printf("missing option --search-file=XXXX\n")
		usage(1)
	}

	// check we have all needed parameters
	if len(flag.Args()) != 1 {
		log.Printf("Wrong number of parameters. Got %d, expected: %d\n", len(flag.Args()), 1)
		usage(1)
	}

	// [optionally] log what we are going to do
	log.MsgVerbose("Checking for files in : %q\n", searchFile)
	filenames := getFiles(searchFile)
	log.MsgVerbose("Found %d files in %q\n", len(filenames), searchFile)

	searchPath := flag.Args()[0]
	log.MsgVerbose("Search path: %q\n", searchPath)

	search(searchPath, filenames)
}
