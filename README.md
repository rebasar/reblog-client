# REBLog Client

This is a simple client for adding entries to [reblog](https://github.com/rebasar/reblog) by directly writing the entries to mongodb. It is mainly written for learning a bit of [Go](https://golang.org) though it lacks most "fancy" features of the language (for example: it does not use any channels)

## Install

Assuming that your go environment is [setup correctly](https://golang.org/doc/code.html), just run:

```
go get github.com/ogier/pflag
go get github.com/spf13/viper
go get gopkg.in/mgo.v2
go get github.com/rebasar/reblog-client
```

TODO: Add a Makefile for automating this process.

For more information on how to use, write `reblog-client -h` in the command line.

Have fun!
