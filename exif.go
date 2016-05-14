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

	"github.com/rwcarlsen/goexif/exif"

	"github.com/sjmudd/findphoto/log"
)

// scan the EXIF data looking for the camera model and check if it's what we are looking for
func checkCameraModel(path string) bool {
	// log.MsgVerbose("checkCameraModel(%q)\n", path)
	if cameraModel == "" {
		return true // we don't care about the camera
	}

	f, err := os.Open(path)
	if err != nil {
		log.Fatalf("Unable to open file %q: %v", path, err)
	}
	defer f.Close()

	// Optionally register camera makenote data parsing - currently Nikon and
	// Canon are supported.
	// exif.RegisterParsers(mknote.All...)

	x, err := exif.Decode(f)
	if err != nil {
		log.Fatal(err)
	}

	camModel, err := x.Get(exif.Model)
	if err != nil {
		log.Fatal("Could not get camera model from file %s: %v", path, err)
	}
	foundModel, err := camModel.StringVal()
	if err != nil {
		log.Fatal("Could not get camera model from file (StringVal failed) %s: %v", path, err)
	}

	if foundModel != cameraModel {
		//		log.Printf("%s taken with camera: %s, expecting: %s\n",
		//			path,
		//			foundModel,
		//			cameraModel)
		return false
	}

	return true // we have matched
}
