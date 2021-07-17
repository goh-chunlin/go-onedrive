# go-onedrive

<div align="center">
    <img src="https://gclstorage.blob.core.windows.net/images/go-onedrive-banner.png" />
</div>

![Go Build](https://github.com/goh-chunlin/go-onedrive/workflows/Go%20Build/badge.svg?branch=main)
![CodeQL Scan](https://github.com/goh-chunlin/go-onedrive/workflows/CodeQL%20Scan/badge.svg?branch=main)
[![Go Report Card](https://goreportcard.com/badge/github.com/goh-chunlin/go-onedrive)](https://goreportcard.com/report/github.com/goh-chunlin/go-onedrive)
[![Go Reference](https://pkg.go.dev/badge/github.com/goh-chunlin/go-onedrive.svg)](https://pkg.go.dev/github.com/goh-chunlin/go-onedrive)
[![License: GPL v3](https://img.shields.io/badge/License-GPLv3-blue.svg)](https://www.gnu.org/licenses/gpl-3.0)
[![Donate](https://img.shields.io/badge/$-donate-ff69b4.svg)](https://www.buymeacoffee.com/chunlin)

go-onedrive is a Golang client library for accessing the [Microsoft OneDrive REST API](https://docs.microsoft.com/en-us/onedrive/developer/rest-api/?view=odsp-graph-online).

This project is inspired by a few open-source projects, especially the [go-github project from Google](https://github.com/google/go-github).

Currently, **go-onedrive requires Golang version 1.15 or greater**.  go-onedrive tracks [Golang version support policy](https://golang.org/doc/devel/release.html#policy). I'll do my best not to break older versions of Golang if I don't have to, but due to tooling constraints, I don't always test older versions.

## Getting Started ##

Module support was introduced in Go 1.15. Starting from Go 1.16, module-aware mode is enabled by default. Hence, I'll assume the module-aware mode is enabled when using this library.

In the go.mod file, please make sure the correct package with the correct version is used.

```
...

require (
	github.com/goh-chunlin/go-onedrive v1.1.1
	...
)
```

The current latest version should be **v1.1.1** (updated on **17th July 2021**, as shown on the [Releases page](https://github.com/goh-chunlin/go-onedrive/releases)).

In other go source files, you can then import the go-onedrive library as follows.
```go
import "github.com/goh-chunlin/go-onedrive/onedrive"
```

Construct a new OneDrive client, then use the various services on the client to access different parts of the OneDrive API. For example:

```go
ctx := context.Background()
ts := oauth2.StaticTokenSource(
	&oauth2.Token{AccessToken: "..."},
)
tc := oauth2.NewClient(ctx, ts)

client := onedrive.NewClient(tc)

// list all OneDrive drives for the current logged in user
drives, err := onedrive.Drives.List(ctx)
```

NOTE: Using the [context](https://godoc.org/context) package, one can easily pass cancelation signals and deadlines to various services of the client for handling a request. In case there is no context available, then `context.Background()` can be used as a starting point.

## Authentication ##

The go-onedrive library does not directly handle authentication. Instead, when creating a new client, pass an `http.Client` that can handle authentication for you. The easiest and recommended way to do this is using the [oauth2](https://github.com/golang/oauth2)
library.

Note that when using an authenticated Client, all calls made by the client will
include the specified OAuth token. Therefore, authenticated clients should
almost never be shared between different users.

See the [oauth2 docs](https://godoc.org/golang.org/x/oauth2) for complete instructions on using that library.

## Contributing ##

This library is being initially developed as a library for my personal project as listed below.
- [Lunar.Music.Web](https://github.com/goh-chunlin/Lunar.Music.Web).

Hence, API methods will likely be implemented in the order that they are needed by my personal project. However, I still welcome you to contribute to this project to support the following features.

- [x] General
	- [x] Async job to track progress
    - [x] Search
- [x] Drives
	- [x] Get default drive
	- [x] Get individual drive
	- [x] List all available drives
- [x] Folders
    - [x] Create
	- [x] Copy
	- [x] Delete
	- [x] List children (items)
    - [x] Move
    - [x] Rename
	- [x] Create share link to a folder and its content
	- [x] List share links of a folder
- [x] Items
	- [x] Get individual item	
	- [x] Copy
	- [x] Delete
    - [x] Move
    - [x] Rename	
	- [x] Create share link to an item
	- [x] Delete share link (or permission) of an item
	- [x] List share links of an item
    - [x] Upload simple item size < 4MB
    - [x] Upload and then replace with item size < 4MB

## Sensei Projects ##

Special thanks go to the following projects for providing useful references which help me in the development of this library.
- [google/go-github](https://github.com/google/go-github);
- [ggordan/go-onedrive](https://github.com/ggordan/go-onedrive);
- [googleapis/google-api-go-client](https://github.com/googleapis/google-api-go-client).

## License ##

This library is distributed under the GPL-3.0 License found in the [LICENSE](./LICENSE) file.
