gomr: go [m]utli-[r]epository manager
-----

gomr is a golang command-line tool used to execute arbitrary commands on a set of tagged directories.

## Features

* Configuration is stored in the registered directory (allows for version-controlled configuration)
* Any arbitrary command can be run, no ties to version-control specific commands.


## Installation

I'm not distributing binaries yet (TODO) for now checkout and `go install` to build the binary.

## Command Help

For help run:
* `gomr --help`
* `gomr register --help`
* `gomr run --help`
