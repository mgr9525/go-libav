# go-libav

[![License](http://img.shields.io/badge/license-MIT-red.svg?style=flat)](https://github.com/mgr9525/go-libav/blob/master/LICENSE.txt)
[![GoDoc](https://godoc.org/github.com/mgr9525/go-libav?status.svg)](https://godoc.org/github.com/mgr9525/go-libav)
[![Build
Status](http://img.shields.io/travis/imkira/go-libav.svg?style=flat)](https://travis-ci.org/imkira/go-libav)
[![Coverage](http://img.shields.io/codecov/c/github/imkira/go-libav.svg?style=flat)](https://codecov.io/github/imkira/go-libav)

[Go](https://golang.org) language bindings for [ffmpeg](https://ffmpeg.org)
libraries.

This is still a work in progress. This package still lacks a lot of the libav's
functionality. Please expect many additions/changes in the future.

# Why

I am aware of other Go language bindings for ffmpeg.
The reason I decided to build go-libav was because I wanted to have:

- A more Object-Oriented Programming approach.
- A more Go-like approach to error handling.
- Easier garbage collection.

# Installation

First, install ffmpeg 3.x libraries on your system.

If you need ffmpeg2.x support, use
[ffmpeg2](https://github.com/mgr9525/go-libav/tree/ffmpeg2) branch (deprecated).

Then, open the terminal and install the following packages:

```
go get -u github.com/mgr9525/go-libav/avcodec
go get -u github.com/mgr9525/go-libav/avfilter
go get -u github.com/mgr9525/go-libav/avformat
go get -u github.com/mgr9525/go-libav/avutil
```

# Documentation

For advanced usage, make sure to check the following documentation:

- [avcodec](http://godoc.org/github.com/mgr9525/go-libav/avcodec)
- [avfilter](http://godoc.org/github.com/mgr9525/go-libav/avfilter)
- [avformat](http://godoc.org/github.com/mgr9525/go-libav/avformat)
- [avutil](http://godoc.org/github.com/mgr9525/go-libav/avutil)

# Examples

Please check [here for examples](https://github.com/mgr9525/go-libav/tree/master/examples).

# FFmpeg versions

This library supports multiple versions of FFmpeg 3.x, to build, use

```
go build -tags ffmpeg33
go test -tags ffmpeg33
go run -tags ffmpeg33 examples/mediainfo/mediainfo.go
```

Use `ffmpeg30` for FFmpeg 3.0 API, `ffmpeg33` for FFmpeg 3.3 API.

# Contribute

Found a bug? Want to contribute and add a new feature?

Please fork this project and send me a pull request!

# License

go-libav is licensed under the MIT license:

www.opensource.org/licenses/MIT

# Copyright

Copyright (c) 2015 Mario Freitas. See
[LICENSE](http://github.com/mgr9525/go-libav/blob/master/LICENSE)
for further details.


# libav_test
need env:
```
CGO_CFLAGS=-I{ffmpegPath}/include;
CGO_LDFLAGS=-L{ffmpegPath}/bin -lavcodec-57 -lavdevice-57 -lavfilter-6 -lavformat-57 -lavutil-55 -lpostproc-54 -lswresample-2 -lswscale-4
```