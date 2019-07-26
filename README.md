# Go Event Source

## Problem Statement
Simple project of how to impelemet event source in Golang. If you don't what event source is, you can [here](https://martinfowler.com/eaaDev/EventSourcing.html)

## Dependencies
In this time, I use mac as an os to running the system.

#### 1. Golang
First of all export some paths, and save them in your `.zshrc` or `.bashrc` files for easy use. Use sudo if you get error.

```
# Go development
```

```
export GOPATH="${HOME}/.go"
```

```
export GOROOT="$(brew --prefix golang)/libexec"
```

```
export PATH="$PATH:${GOPATH}/bin:${GOROOT}/bin"
```

```
test -d "${GOPATH}" || mkdir "${GOPATH}"
```

And then, you can install the Go
```
brew install go

```

## How to use
You can run the program by running this command at terminal :
```
go run app/cmd/main.go
```
