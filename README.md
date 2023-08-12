# todo (in golang)   ![go workflow](https://github.com/elusive/go-todo/actions/workflows/go.yml/badge.svg) 
Go command line utility todo, for saving tasks and completed state to/from a json file.

## Overview
This is a small application built while reading the book "Powerful Command Line Application" by Ricardo Gerardi. 
Which is available from the [Pragmatic Bookshelf](https://pragprog.com/titles/rggo/powerful-command-line-applications-in-god).

## Usage
To use this utility you can build it locally if you have Go installed.  If you do not have Go installed you can follow the 
instructions for your operating system at (https://go.dev/doc/install) and then return here.

1. Clone the repository onto your local drive (or download the zip):
   ```bash
   git clone git@github.com:elusive/go-todo.git
   ```
2. Go to the cloned repository folder and build the application:
   ```bash
   cd go-todo && go build
   ```

Once you build the application you should see the binary in the repository folder.  Either "words.exe" or "words". The 
`go build` command will automatically build the application for your specific OS.  If you wish to build cross-platform
for an OS and Architecture other than your own you can do so using these commands:
```bash
env GOOS=windows GOARCH=amd64 go build
env GOOS=darwin GOARCH=arm64 go build
env GOOS=linux GOARCH=386 go build
```

### CLI Tests
The cmd folder contains the cli interface on top of the todo api. The tests for this cli utilize the os package
and will build the todo binary and use it in the tests via the `os.exec.Command` method in order to make sure 
that the cli commands are functional.

### Execute Tests
Remember that to execute the tests for the todo api you need to have built the todo binary using `go build` and
the cli tests will rebuild the binary on their own as part of the tests code. You can execute these tests also, 
using this command:
```bash
go test -v
```

The current unit tests should result in the following output to the terminal:
```bash
=== RUN   TestCountWords
--- PASS: TestCountWords (0.00s)
=== RUN   TestCountLines
--- PASS: TestCountLines (0.00s)
=== RUN   TestCountBytes
--- PASS: TestCountBytes (0.00s)
=== RUN   TestLines
--- PASS: TestLines (0.00s)
PASS
ok      github.com/elusive/words        0.056s
```

## Further Questions
For any further questions please email:  `me at johng dot info`
