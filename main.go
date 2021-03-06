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
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/sjmudd/findphoto/log"
)

const (
	copyright = "(C) 2016 Simon J Mudd <sjmudd@pobox.com>"
	myVersion = "0.0.6"
)

var (
	help             bool   // do we want help?
	myName           string // program name taken from os.Args[0]
	cameraModel      string // e.g. Camera Model Name : Canon PowerShot S100
	searchFile       string // file containing photo names
	progressInterval int    // interval at which to give progress on the search
	version          bool   // show the program version
	symlinkDir       string // directory where to make symlinks
	showCameraModel  bool   // show the camera model of the given file
	showExifData     bool   // show the EXIF data of a file (part of debugging)
	verbose          bool   // change to verbose mode
	debug            bool   // change to debug mode
	counters         *Stats
)

func init() {
	myName = os.Args[0]
	// Split at /
	stuff := strings.Split(myName, "/")
	if len(stuff) > 1 {
		myName = stuff[len(stuff)-1]
	}
}

// given a filename to collect names from return a list of names
func getFilenames(filename string) []string {
	if filename == "" {
		return nil // no input -> no output
	}
	var filenames []string

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		entry := scanner.Text()
		// log.MsgVerbose("Entry: %s\n", entry)
		filenames = append(filenames, entry)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return filenames
}

// showVersion shows the program version and exits
func showVersion() {
	fmt.Printf("%s version %s\n", myName, myVersion)
	os.Exit(0)
}

// usage returns a usage message and exits with the requested exit code
func usage(exitCode int) {
	fmt.Printf("%s version %s %s\n\n", myName, myVersion, copyright)
	fmt.Printf("Find photo files given various search criteria\n\n")
	fmt.Printf("Usage: %s <options> <directory_to_search>\n", myName)
	flag.PrintDefaults()

	os.Exit(exitCode)
}

func checkSymlinkDir(name string) {
	info, err := os.Stat(name)
	if err != nil {
		log.Fatal("Failed to stat symlink-dir %s: %v", name, err)
	}
	if !info.Mode().IsDir() {
		log.Fatal("symlinkdir %s is not a directory", name)
	}

	log.Printf("symlink dir: %s\n", symlinkDir)
}

func main() {
	// get options
	flag.BoolVar(&help, "help", false, "shows this help message")
	flag.BoolVar(&verbose, "verbose", false, "Enable verbose logging")
	flag.BoolVar(&debug, "debug", false, "Enable debug logging")
	flag.BoolVar(&showCameraModel, "show-camera-model", false, "Show the camera model for the file(s) specified")
	flag.BoolVar(&showExifData, "show-exif-data", false, "Show the EXIF data of a file. (intended for debugging)")
	flag.BoolVar(&version, "version", false, "shows the program version and exits")
	flag.IntVar(&progressInterval, "progress-interval", 60, "time in verbose mode to give an indication of progress")
	flag.StringVar(&cameraModel, "camera-model", "", "provide camera model (in exif data) e.g. 'Canon PowerShot S100'")
	flag.StringVar(&searchFile, "search-file", "", "Required: File to use containing a line of the base filesnames to search for")
	flag.StringVar(&symlinkDir, "symlink-dir", "", "directory to symlink found files against")
	flag.Parse()

	if help {
		usage(0)
	}
	if version {
		showVersion()
	}
	if debug {
		log.LogLevel = log.LogLevelDebug
	} else if verbose {
		log.LogLevel = log.LogLevelVerbose
	}
	log.MsgVerbose("LogLevel: %s", log.LogLevelName())

	// show the version when running in verbose mode
	log.Printf("%s version %s\n", myName, myVersion)

	if cameraModel != "" {
		log.Printf("camera-model: %s\n", cameraModel)
	}
	if symlinkDir != "" {
		checkSymlinkDir(symlinkDir)
	}
	if searchFile == "" && !showCameraModel {
		log.Printf("missing option --search-file=XXXX\n")
		usage(1)
	}
	if progressInterval <= 0 {
		log.Printf("--progress-interval should be a positive number of seconds\n")
		usage(1)
	} else {
		counters = NewStats(time.Duration(progressInterval))
	}

	// check we have all needed parameters
	if len(flag.Args()) != 1 {
		log.Printf("Wrong number of parameters. Got %d, expected: %d\n", len(flag.Args()), 1)
		usage(1)
	}

	log.MsgVerbose("progress interval: %d\n", progressInterval)

	// [optionally] log what we are going to do
	var filenames []string
	if searchFile != "" {
		filenames = getFilenames(searchFile)
		log.MsgVerbose("Found %d filename(s) in %q\n", len(filenames), searchFile)
	}

	searchPath := flag.Args()[0]
	log.MsgVerbose("Search path: %q\n", searchPath)

	search(searchPath, filenames)
}
