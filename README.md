GoDisco: Discourse REST API Client
===============================
<p align="center">
  <a href="http://golang.org" target="_blank"><img alt="Go package" src="https://golang.org/doc/gopher/pencil/gopherhat.jpg" width="20%" /></a>
  <a href="https://www.discourse.org/" target="_blank"><img src="https://pbs.twimg.com/profile_images/3264780953/6c9a2cd7bb2efcb4c53d32900e52c8ac_400x400.png" alt="Discourse Logo"/></a>
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
