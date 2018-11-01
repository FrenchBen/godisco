GoDisco: Discourse REST API Client
===============================
[![Build Status](https://img.shields.io/travis/FrenchBen/godisco/master.svg?style=flat-square)](https://travis-ci.org/FrenchBen/godisco)
[![Codecov branch](https://img.shields.io/codecov/c/github/FrenchBen/godisco/master.svg?style=flat-square)](https://codecov.io/gh/FrenchBen/godisco)
[![GoDoc](https://img.shields.io/badge/godoc-reference-5272B4.svg?style=flat-square)](https://godoc.org/github.com/FrenchBen/godisco)
[![Go Report Card](https://goreportcard.com/badge/github.com/FrenchBen/godisco?style=flat-square)](https://goreportcard.com/report/github.com/FrenchBen/godisco)

<p align="center">
  <a href="http://golang.org" target="_blank"><img alt="Go package" src="https://golang.org/doc/gopher/pencil/gopherhat.jpg" width="20%" /></a>
  <a href="https://www.discourse.org/" target="_blank"><img src="https://raw.githubusercontent.com/discourse/discourse/master/images/discourse.png" alt="Discourse Logo"/></a>
</p>

About
----------------
Unofficial Golang client for the Discourse.org REST API: https://meta.discourse.org/t/discourse-api-documentation/22706.

Requires Go `1.5.3`

Installation
----------------
The recommended way of installing the client is via `go get`. Simply run the following command to add the package.

    go get github.com/FrenchBen/godisco/

Usage
----------------
Below is an example of how to use this library

```
package main

import (
	"github.com/FrenchBen/godisco"
	"github.com/Sirupsen/logrus"
)


func main() {
  discourseClient, err := godisco.NewClient("http://discourse.example.com", "api_token", "api_username")
	if err != nil {
		logrus.Fatal(err)
	}
  discourseUser, err := godisco.GetUser(discourseClient, "SomeDiscourseUserName")
  if err != nil {
    logrus.Error(err)
  }
  logrus.Infof("User Info: %v", discourseUser)
}
```

To view more the token and fields sent with the request, set your log level to debug:
`logrus.SetLevel(logrus.DebugLevel)`


License
----------------
This source is licensed under an MIT License, see the LICENSE file for full details. If you use this code, it would be great to hear from you.
