# aoc-2025

This is my *attempt* at solving AOC 2025.
I am writing this in Golang.
I have been using Golang for the better part of this past year and have really grown to love the language.

## Running the Project

### Install Go

Instructions can be found [here](https://go.dev/doc/install).

### Install Dependencies

```
go mod tidy
```

### Run the CLI App

The below command will print out help docs:

```
go run main.go --help
```

Below is an example of running a particular day's solution:

```
go run main.go day01 --input <path-to-your-input-fule> --part 1
```

### Run Tests

```
go test ./...
```
