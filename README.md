# Go Event Source

## Problem Statement
Simple project of how to impelemet event source in Golang. If you don't know what event source is, I will give you a glimpse about that. Event Sourcing ensures that all changes to application state are stored as a sequence of events. Not just can we query these events, we can also use the event log to reconstruct past states, and as a foundation to automatically adjust the state to cope with retroactive changes. For futher information you can check [here](https://martinfowler.com/eaaDev/EventSourcing.html)

## How it works
The fundamental idea of Event Sourcing is that of ensuring every change to the state of an application is captured in an event object, and that these event objects are themselves stored in the sequence they were applied for the same lifetime as the application state itself.

Let's consider a simple example to do with shipping notifications. In this example we have many ships on the high seas, and we need to know where they are. A simple way to do this is to have a tracking application with methods to allow us to tell when a ship arrives or leaves at a port.

![Github Logo](https://martinfowler.com/eaaDev/eventSourcing/simpleService.gif)

In this case when the service is called, it finds the relevant ship and updates its location. The ship objects record the current known state of the ships.

Introducing Event Sourcing adds a step to this process. Now the service creates an event object to record the change and processes it to update the ship.

![Github Logo](https://martinfowler.com/eaaDev/eventSourcing/simpleEventCd.gif)

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
