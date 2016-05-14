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

// Package log is for doing logging
// - it can use the same interface as "log" but also has other options
package log

import (
	realLog "log"
)

const (
	backslashN = "\n"
)

var (
	Verbose         bool
	backslashNValue = backslashN[0]
)

// -------------------------------------------------------------------------
// wrapper around the existing names
// -------------------------------------------------------------------------

// Print behaves the same as log.Print
func Print(v ...interface{}) {
	realLog.Print(v...)
}

// Printf behaves the same as log.Printf
func Printf(format string, v ...interface{}) {
	MsgVerbose(format, v...)
}

// Fatal behaves the same as log.Fatal
func Fatal(v ...interface{}) {
	realLog.Fatal(v...)
}

// Fatalf behaves the same as log.Fatalf
func Fatalf(format string, v ...interface{}) {
	realLog.Fatalf(format, v...)
}

// -------------------------------------------------------------------------
// wrapper around the existing names
// -------------------------------------------------------------------------

// MsgInfo is like log.Printf
func MsgInfo(format string, v ...interface{}) {
	realLog.Printf(format, v...)
}

// MsgVerbose is like log.Printf IFF Verbose is set
func MsgVerbose(format string, v ...interface{}) {
	if Verbose {
		MsgInfo(format, v...)
	}
}

// MsgError is like MsgInfo but prefixes the message with ERORR:
func MsgError(format string, v ...interface{}) {
	MsgInfo("ERROR: "+format, v...)
}
