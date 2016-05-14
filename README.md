## findphoto

findphoto - simple script to find photos given certain criteria


ps-top is a program which collects information from MySQL 5.6+'s
performance_schema database and uses this information to display
server load in real-time. Data is shown by table or filename and
the metrics also show how this is split between select, insert,
update or delete activity.  User activity is now shown showing the
number of different hosts that connect with the same username and
the activity of those users.  There are also statistics on mutex
and sql stage timings.

ps-stats is a similar utility which provides output in stdout mode.

### Installation

Install as follows:
`go get -u github.com/sjmudd/ps-top/findphoto`

The sources will be downloaded together with the dependencies and
the binary will be built and installed into `$GOPATH/bin/`. If
this path is in your `PATH` setting then the program can be run
directly without having to specify any specific path.

### Dependencies

These are included using GO15VENDOREXPERIMENT=1

### Configuration

### Usage

Currently this is very simple. It allows to search for files by
filename (e.g. in your backup) and with a specific camera model
(taken from the EXIF data). This is because some file names may be
duplicated and from different cameras.

`$ findphoto --verbose --camera-model="Canon EOS 20D" --search-file=missing.ALL /path/where/to/look/for/files`

### To do

* add more search criteria or filters
* copy the files found to a specific location for further analysis

### Contributing

Suggestions and improvements are most welcome. If you have such
feedback please create an issue on github.

### Licensing

BSD 2-Clause License

### Feedback

Feedback and patches welcome. Feedback other than patches or
bug reports can be sent to my email address below.

Simon J Mudd
<sjmudd@pobox.com>

### Code Documenton
[godoc.org/github.com/sjmudd/findphoto](http://godoc.org/github.com/sjmudd/findphoto)
