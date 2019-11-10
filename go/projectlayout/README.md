# Overview

A very common question Go beginners have is "how do I organize my code?". Some of the things like:

* How does my repository structure reflect the way users import my code?
* How do I distribute commands (command-line programs that users can install) in addition to code?
* How do modules change the way I organize my code?
* How do multiple packages coexist in a single module?

Unfortunately, there is some easy-to-find advice online that's outdated and over-complicated

The concepts demonstrated here:

* Splitting a module into multiple packages, each importable by users and some of these packages importing others (from within the same module).
* Internal packages, only importable from other packages in their module, not by outside users.
* Commands/programs that users can install with go get.

developers use the modules either by import-ing it in their code, or by go get-ing a program.

## Getting started

The sample project is on GitHub: https://github.com/phaneendra/go-eggs/project-layout

In this case, the project path is the module name. The go.mod file for the project contains this line:

```
module github.com/phaneendra/go-eggs/project-layout
```
It is very common for Go projects to be named by their GitHub path. Go also supports custom names. Throughout this doc, you can substitute `github.com/phaneendra/go-eggs/project-layout` with `github.com/<your-handle>/<your-project>` or `your-project-domain.io`, whatever works for you.

The module name is extremely important, because it serves as the basis of imported names in user code:


## Project layout
Here is the directory and file layout of the `go-eggs/project-layout` repository:

```
├── LICENSE
├── README.md
├── config.go
├── go.mod
├── go.sum
├── clientlib
│   ├── lib.go
│   └── lib_test.go
├── cmd
│   ├── client
│   │   └── main.go
│   └── server
│       └── main.go
├── internal
│   └── auth
│       ├── auth.go
│       └── auth_test.go
└── serverlib
    └── lib.go
```


* `LICENSE` and `README.md` are fairly obvious and I won't spend time on them here.

* go.mod is the module definition file. It contains the module name shown above and that's it - in this case the project has no dependencies.

* go.sum contains all the dependency checksums, and is managed by the go tools. You don't have to worry about it, but keep it checked into source control alongside go.mod.

* config.go this is the first code file we're examining; it contains a single trivial function [1]:

```
package project-layout

func Config() string {
  return "project-layout config"
}
```

The most important part here is the `package project-layout`. Since this file is at the top level of the module, its package name is considered to be the module name. This is what you get when you just `import github.com/phaneendra/go-eggs/project-layout`. The user code can look like this:

```
package main

import "fmt"
import "github.com/phaneendra/go-eggs/project-layout"

func main() {
  fmt.Println(project-layout.Config())
}
```

So the rule is simple: if your module provides a single package, or you want to export code from the top-level package of the module, place all the code for this at the top-level directory of the module, and name the package as the last part of the module's path (unless you're using vanity imports, in which case it's more flexible).

## Additional packages

- `clientlib/lib.go` is a file in the clientlib package of our module. It doesn't matter what the file is called, and many packages consist of multiple files. What's important is that the package declaration at the top of the file says clientlib:

```
package clientlib

func Hello() string {
  return "clientlib hello"
}
```

User code will import this package with `github.com/phaneendra/go-eggs/project-layout/clientlib`, as follows:

```
package main

import "fmt"
import "github.com/phaneendra/go-eggs/project-layout"
import "github.com/phaneendra/go-eggs/project-layout/clientlib"

func main() {
  fmt.Println(go-eggs/project-layout.Config())
  fmt.Println(clientlib.Hello())
}
```

- The serverlib directory contains another package users can import. There's nothing new there - just showing how multiple packages live alongside each other.

A quick word on nesting of packages: it can go as deep as you need. The package name visible to users is determined by the relative path from the module root. For example, if we have a subdirectory called `clientlib/tokens` with some code in the tokens package, the user will import that with import `github.com/phaneendra/clientlib/tokens`.

It's also important to highlight that for some modules a single top-level package is sufficient. In the case of `go-eggs/project-layout` this would mean no subdirectories with user-importable packages, but all code being in the top directory in a single or multiple Go files all in package `go-eggs/project-layout`.

## Commands / programs
Some Go projects distribute programs, or commands, instead of (or in addition to) importable packages. If this isn't relevant to your project, don't add a `cmd` directory.

The `cmd` directory is the conventional location of all the command-line programs made available by the project. The naming scheme for programs is typically:

`github.com/username/projectname/modulename/cmd/cmd-name`

Such commands can be installed by the user using the go tool as follows:

```
$ go get github.com/phaneendra/go-eggs/project-layout/cmd/cmd-name

# Go downloads, builds and installs cmd-name into the default location.
# The bin/ directory in the default location is often in $PATH, so we can
#just invoke cmd-name now

$ cmd-name ...
```
In `go-eggs/project-layout`, there are two different command-line programs provided, as an example: `client` and `server`. In each of them, the code is in `package main;` the filename is also called `main.go`, but this isn't a requirement. It doesn't matter what the file names are called, as long as they're in `package main`.

In fact, since `go-eggs/project-layout` is a real repository, you can install and run these tools on your machine:

```
$ go get github.com/phaneendra/go-eggs/project-layout/cmd/go-eggs/project-layout-client
$ go-eggs/project-layout-client
Running client
Config: go-eggs/project-layout config
clientlib hello

$ go get github.com/phaneendra/go-eggs/project-layout/cmd/go-eggs/project-layout-server
$ go-eggs/project-layout-server
Running server
Config: go-eggs/project-layout config
Auth: thou art authorized
serverlib hello

# Clean up...
$ rm -f `which go-eggs/project-layout-server` `which go-eggs/project-layout-client`
```

It's instructional to take a look at the code of `go-eggs/project-layout-server`:

```
package main

import (
  "fmt"

  "github.com/phaneendra/go-eggs/project-layout"
  "github.com/phaneendra/go-eggs/project-layout/internal/auth"
  "github.com/phaneendra/go-eggs/project-layout/serverlib"
)

func main() {
  fmt.Println("Running server")
  fmt.Println("Config:", go-eggs/project-layout.Config())
  fmt.Println("Auth:", auth.GetAuth())
  fmt.Println(serverlib.Hello())
}
```

The important thing to highlight here is how it imports other code from `go-eggs/project-layout`. In Go, absolute imports are the way to go - note how this is used here. This applies to packages as well, not just commands. If code in package clientlib needs to import the main `go-eggs/project-layout` package, it will do so by `import github.com/phaneendra/go-eggs/project-layout`.

# Internal packages
Another important concept is internal (or private) packages - packages that are used internally by a project, but which we don't want to export to users. This is especially important in Go with modules, due to semantic versioning. Everything exported by your project in v1 becomes a public API, and has to abide by semantic versioning compatibility guarantees. Therefore, it's imperative to export only the minimal API surface that's essential for users of your project. All the other code which your package needs for its implementation should live in internal.

The Go tooling recognizes internal as a special path. Packages in the same module can import it as usual (see the previous code snippet, for example). Users (that is, code outside the module) cannot import it, though. If we try to do this, we get an error:

```
use of internal package github.com/phaneendra/go-eggs/project-layout/internal/auth not allowed
```

In the `go-eggs/project-layout` project, there's a single package in internal. In real projects, there is often a whole tree of packages there.

If you're wondering whether some package belongs in internal, it's prudent to begin by answering "yes". It's easy to take an internal API and export it to users - just a quick renaming/refactoring commit. It's very painful to take an external API and un-export it (user code may depend on it); at stable module versions (v1 and beyond), this requires a major version bump to break compatibility.

Put as much as possible in internal, not only private Go packages needed by my module. For example, if the repository contains the source code of the website of the project, I'd place that in internal/website. The same goes for internal tools/scripts needed to work on the project. The idea is that the root directory of a project should be minimal and clear to users. In a way, it's self-documentation. A user looking at my project's GitHub page should get an immediate sense of where the things they need are located. Since users don't typically really need the stuff I use to develop the project, hiding it in internal makes sense.