# Go-announcer

Go-announcer lets you output information, warning or errors with optional key-value contexts.

## Install

```
go get github.com/nicholastcs/go-announcer
```

## Feature

* Auto word-wrapping if the announcement string is long

* Automatically emit to stderr if use `ann.Error(...)`

* Emphasis bar prepend before messages, color and bar symbol changes when use `ann.Warn(...)` or `ann.Error(...)`

## Examples

Basic 

```go
ann := announcer.New()

// simple announcement
ann.Tell("Discarded value")

// ... with context
ann.Tell("Authentication success", announcer.Args().
        AddContext("Session-ID", sId).
        AddContext("Role", userRole))

// warning announcement
ann.Warn("Limited access", announcer.Args().
        AddContext("Session-ID", sId))

// error announcement, third argument in .AddContext(...)
// will colorize Cause to red, if true.
ann.Error("Unauthorized access", announcer.Args().
        AddContext("Cause", err.Error(), true))
```

