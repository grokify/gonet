GoNet
=====

[![Build Status][build-status-svg]][build-status-url]
[![Go Report Card][goreport-svg]][goreport-url]
[![Docs][docs-godoc-svg]][docs-godoc-url]
[![License][license-svg]][license-url]

## Overview

`grokify/gonet` is part of a family of Go modules that provides extended
functionality for common tasks in Go. HTTP related items are currently in
[`grokify/gohttp`](https://github.com/grokify/gohttp).

Go is a small language which is useful from a development and maintenance
perspective but it can result in common tasks being more verbose than other 
languages where there are more productivity methods. The family of packages
has the goal to provide wrappers for common tasks in the same spirit of `io/ioutil`
to make programming Go a bit faster.

See the following modules for extended functionality.

* Mogo - [github.com/grokify/mogo](https://github.com/grokify/mogo) (limited dependencies)
* GoAuth - [github.com/grokify/goauth](https://github.com/grokify/goauth)
* GoCharts - [github.com/grokify/gocharts](https://github.com/grokify/gocharts)
* GoHTTP - [github.com/grokify/gohttp](https://github.com/grokify/gohttp)
* GoNet - [github.com/grokify/gonet](https://github.com/grokify/gonet)
* GoPhoneNumbers - [github.com/grokify/gophonenumbers](https://github.com/grokify/gophonenumbers)

Note: `grokify/gonet` was initially implemented in `grokify/mogo` and extracted
to reduce the number of dependences in `grokify/mogo`.

## Documentation

Documentation is provided using godoc and available on [GoDoc.org](https://godoc.org/github.com/grokify/gonet).

- [imaputil](https://pkg.go.dev/github.com/grokify/gonet/imaputil)
- [mailutil](https://pkg.go.dev/github.com/grokify/gonet/mailutil)
- [sftputil](https://pkg.go.dev/github.com/grokify/gonet/sftputil)
- [urlutil](https://pkg.go.dev/github.com/grokify/gonet/urlutil)

## Installation

```bash
$ go get github.com/grokify/gonet/...
```

## Contributing

Features, Issues, and Pull Requests are always welcome.

To contribute:

1. Fork it ( http://github.com/grokify/gonet/fork )
2. Create your feature branch (`git checkout -b my-new-feature`)
3. Commit your changes (`git commit -am 'Add some feature'`)
4. Push to the branch (`git push origin my-new-feature`)
5. Create new Pull Request

Please report issues and feature requests on [Github](https://github.com/grokify/mogo).

 [used-by-svg]: https://sourcegraph.com/github.com/grokify/gonet/-/badge.svg
 [used-by-url]: https://sourcegraph.com/github.com/grokify/gonet?badge
 [build-status-svg]: https://github.com/grokify/gonet/workflows/go%20build/badge.svg?branch=master
 [build-status-url]: https://github.com/grokify/gonet/actions
 [goreport-svg]: https://goreportcard.com/badge/github.com/grokify/gonet
 [goreport-url]: https://goreportcard.com/report/github.com/grokify/gonet
 [codeclimate-status-svg]: https://codeclimate.com/github/grokify/gonet/badges/gpa.svg
 [codeclimate-status-url]: https://codeclimate.com/github/grokify/gonet
 [docs-godoc-svg]: https://pkg.go.dev/badge/github.com/grokify/gonet
 [docs-godoc-url]: https://pkg.go.dev/github.com/grokify/gonet
 [license-svg]: https://img.shields.io/badge/license-MIT-gonet.svg
 [license-url]: https://github.com/grokify/gonet/blob/master/LICENSE
