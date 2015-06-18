rumr: [Ru]by [m]utli-[r]epository Tool
-----

Rumr is a ruby command-line tool used to execute arbitrary commands on a set of tagged directories.

## Features

* Configuration is stored in the registered directory (allows for version-controlled configuration)
* Any arbitrary command can be run, no ties to version-control specific commands.


## Installation

From rubygems: `gem install rumr`

From source:
* `gem build rumr.gemspec`
* `gem install rumr-[version].gem`

## Command Help

For help run:
* `rumr help`
* `rumr help register`
* `rumr help exec`
