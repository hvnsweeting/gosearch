# gosearch: The awesome package explorer for Golang
[![Build Status](https://travis-ci.org/hvnsweeting/gosearch.svg?branch=master)](https://travis-ci.org/hvnsweeting/gosearch)

`gosearch` turns the [awesome-go](https://github.com/avelino/awesome-go) into a
CLI command that you can search for packages offline easily.

## Why this?
I've back to coding Golang after a long time off.
I want to install glide, logrus... and all the awesome packages, but I
don't remember its full package import path, how to not Google?

Go ships with the good `go get` to install package, but not a way to search for packages.

Inspired by [Python pip](https://pip.pypa.io/en/stable/).

## Install

```
$ go get -u github.com/hvnsweeting/gosearch
```

## Usage

```sh
$ gosearch
Usage: gosearch packagename
       gosearch [OPTIONS] [OPTIONS arguments]

Options:
  -c category
    	Show packages in category. Use `all` for list of all categories.
  -r	Show the raw data of Awesome-go
```

### Search package by name

```sh
$ gosearch logrus
Package: github.com/Sirupsen/logrus
Category: logging
Description-en: Structured logger for Go.
$ gosearch echo
Package: github.com/labstack/echo
Category: web frameworks
Description-en: High performance, minimalist Go web framework.
```

### Listing packages by category

```sh
$ gosearch -c logging
github.com/kpango/glg - glg is simple and fast leveled logging library for Go.
github.com/golang/glog - Leveled execution logs for Go.
github.com/utahta/go-cronowriter - Simple writer that rotate log files automatically based on current date and time, like cronolog.
github.com/siddontang/go-log - Log lib supports level and multi handlers.
github.com/ian-kent/go-log - Log4j implementation in Go.
github.com/apsdehal/go-logger - Simple logger of Go Programs, with level handlers.
...
```

### List all categories

```
 $ gosearch -c all
audio and music: 15 packages
authentication and oauth: 21 packages
benchmarks: 14 packages
code analysis: 22 packages
command line: 40 packages
conferences: 11 packages
configuration: 19 packages
continuous integration: 4 packages
css preprocessors: 3 packages
data structures: 29 packages
...
```

### Get the raw awesome-go file

so you can fallback if `gosearch` ever goes wrong.

```
$ gosearch -r
# Awesome Go [![Build Status](https://travis-ci.org/avelino/awesome-go.svg?branch=master)](https://travis-ci.org/avelino/awesome-go) [![Awesome](https://cdn.rawgit.com/sindresorhus/awesome/d7305f38d29fed78fa85652e3a63e154dd8e8829/media/badge.svg)](https://github.com/sindresorhus/awesome) [![Join the chat at https://gitter.im/avelino/awesome-go](https://badges.gitter.im/avelino/awesome-go.svg)](https://gitter.im/avelino/awesome-go?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge&utm_content=badge)

A curated list of awesome Go frameworks, libraries and software. Inspired by [awesome-python](https://github.com/vinta/awesome-python).

### Contributing
...
```

## LICENSE

This package is made available under an MIT-style license. See LICENSE.txt.
