<h1 align="center">virt-lab-creator</h1>

*<p align="center">Basic CLI word counter App built with Go</p>*

## About

This is an exercise from the book _Powerful Command-Line Applications in Go_.

This is a simple command-line tool that reads STDIN and count words, lines or bytes.

## Usage

Pipe text to `wc` command, it returns the number of words by default.

Using flags to modify the counter:

- `-l` - count lines
- `-b` - count bytes

```
# Count words:
$ echo "two words" | ./wc
2 # result
```

```
# Count lines using -l flag:
$ cat main.go | ./wc -l
32 # result
```

```
# Count bytes using -b flag:
$ echo "lala" | ./wc -b
5 # result
```
