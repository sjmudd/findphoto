# findphoto

findphoto is a simple script to find photos given certain criteria.

## Installation

Install as follows:
`go get -u github.com/sjmudd/ps-top/findphoto`

The sources will be downloaded together with the dependencies and
the binary will be built and installed into `$GOPATH/bin/`. If
this path is in your `PATH` setting then the program can be run
directly without having to specify any specific path.

## Dependencies

These are included using standard Go Vendoring: `GO15VENDOREXPERIMENT=1`

## Configuration

No configuration needed.

## Usage

Currently this is very simple. It allows to search for files by
filename (e.g. in your backup) and with a specific camera model
(taken from the EXIF data). This is because some file names may be
duplicated and from different cameras.

If we have a file with missing photos such as `missing.ALL`
```
IMG_0001.JPG
IMG_0002.JPG
IMG_0003.JPG
```
then we can do the following:

* Find missing files in `missing.ALL` in path `/search/path`
```
$ findphoto \
	--verbose \
	--camera-model="Canon EOS 20D" \
	--search-file=missing.ALL \
	/search/path
```

* Find missing files in `missing.ALL` in path `/search/path` and make
symbolic links to them in directory `results`.
```
$ findphoto \
	--symlink-dir=results \
	--camera-model="Canon EOS 20D" \
	--search-file=missing.ALL \
	/search/path
```

* Show the camera model of some files:

```
$ ./findphoto \
	--verbose \
	--show-camera-model \
	/path/to/some/photo/files
```

## To do

* add more search criteria or filters
* copy the files found to a specific location for further analysis

## Contributing

Suggestions and improvements are most welcome. If you have such
feedback please create an issue on github.

## Licensing

BSD 2-Clause License

## Feedback

Feedback and patches welcome. Feedback other than patches or
bug reports can be sent to my email address below.

Simon J Mudd
<sjmudd@pobox.com>

## Code Documenton
[godoc.org/github.com/sjmudd/findphoto](http://godoc.org/github.com/sjmudd/findphoto)
