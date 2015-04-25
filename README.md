# Flags [![GoDoc](http://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)](http://godoc.org/github.com/fatih/flags) [![Build Status](http://img.shields.io/travis/fatih/flags.svg?style=flat-square)](https://travis-ci.org/fatih/flags)


Flags is a low level package for parsing or managing single flag arguments and
their associated values from a list of arguments. It's useful for CLI
applications or creating logic for parsing arguments(custom or os.Args)
manually. Checkout the usage below for examples:

## Install

```bash
go get github.com/fatih/flags
```

## Usage and examples

Let us define three flags

```go
args := []string{"--key", "123", "--name=example", "--debug"}
```

Check if a flag exists in the argument list

```go
flags.HasFlag("key", args)    // true
flags.HasFlag("--key", args)  // true
flags.HasFlag("secret", args) // false
```

Get the value for from a flag name

```go
val, _ := flags.ValueFrom("--key", args) // val -> "123"
val, _ := flags.ValueFrom("name", args)  // val -> "example"
val, _ := flags.ValueFrom("debug", args) // val -> "" (means boolean)
```

Exclude a flag and it's value from the argument list

```go
rArgs := flags.ExcludeFlag("key", args)  // rArgs -> ["--name=example", "--debug"]
rArgs := flags.ExcludeFlag("name", args) // rArgs -> ["--key", "123", "--debug"]
rArgs := flags.ExcludeFlag("foo", args)  // rArgs -> ["--key", "123", "--name=example "--debug"]
```

Is a flag in its valid representation?

```go
flags.IsFlag("foo")           // false
flags.IsFlag("--foo")         // true
flags.IsFlag("-key=val")      // true
flags.IsFlag("-name=example") // true
```

Parse a flag and return the name

```go
name, _ := flags.ParseFlag("foo")        // returns error, because foo is invalid
name, _ := flags.ParseFlag("--foo")      // name -> "foo
name, _ := flags.ParseFlag("-foo")       // name -> "foo
name, _ := flags.ParseFlag("-foo=value") // name -> "foo
name, _ := flags.ParseFlag("-foo=")      // name -> "foo
```

