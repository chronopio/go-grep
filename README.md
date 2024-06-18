<h1 align="center">
  Go-Grep<br/>
</h1>

<br/>

<center>

[![forthebadge](https://forthebadge.com/images/badges/built-with-love.svg)](https://forthebadge.com) &nbsp;

</center>

## TL;DR

We all know grep, short for “global regular expression print”, is a command used for searching and matching text patterns in files. This is an improvement that implements go concurrency patterns to make it faster using multithreading.

## Built With
This project was built using these technologies.

- Go

## Flags

**SearchTerm:**
The pattern that is going to be used in the search.
- string - positional / required

**SearchDir:**
- string - positional / defaults to current directory.

**Currently amount of workers is not modifiable and it's set to 15.**

## Getting Started

This software uses go 1.22.3 and it's required to use the tool.

## 🛠 Installation and Setup Instructions

1. Go to the project root folder.
2. Run go mod tidy.
3. Run `go run ./mgrep (SearchTerm) (SearchDir)`

### Useful Links

- Grep reference: https://www.freecodecamp.org/news/grep-command-in-linux-usage-options-and-syntax-examples/
