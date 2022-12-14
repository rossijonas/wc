<h1 align="center">wc (Word Counter)</h1>

*<p align="center">A Simple Word Counter for the CLI.</p>*

## About

`wc` reads STDIN and counts words, lines, or bytes.

### Status:

[![Actions Status](https://github.com/rossijonas/wc/workflows/Test/badge.svg)](https://github.com/rossijonas/wc/actions)
[![Actions Status](https://github.com/rossijonas/wc/workflows/Build/badge.svg)](https://github.com/rossijonas/wc/actions)

### Features:

- Cross platform:  Linux / Macos / Windows.

- Counts words, lines or bytes from an input in STDIN.

## Installation

### Requirements:

- [Go](https://go.dev/) version 1.18.6 (or above)

### How to install:

- Run: 

  ```
  $ go install github.com/rossijonas/wc@latest
  ```

## Usage

### Options:

```
$ wc -h
Usage of wc:
  -b    Count bytes
  -l    Count lines
```

### Examples:

Pipe text to `wc` command, it returns the number of words by default.

#### Count words:

```
$ echo "two words" | ./wc
2 # result
```

#### Count lines using `-l` flag:

```
$ cat main.go | ./wc -l
32 # result
```

#### Count bytes using `-b` flag:

```
$ echo "lala" | ./wc -b
5 # result
```

## Backlog

- Add cover image to README file.

- Support accepting multiple files.

## Credits

_This is an exercise from the book "Powerful Command-Line Applications in Go", but it may differ from the original exercise._
